// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include <dirent.h>
#include <stdint.h>
#include <stdlib.h>
#include <sys/mount.h>
#include <sys/syscall.h>

#include <string>
#include <vector>

#include "gmock/gmock.h"
#include "gtest/gtest.h"
#include "test/util/capability_util.h"
#include "test/util/fs_util.h"
#include "test/util/temp_path.h"
#include "test/util/test_util.h"
#include "test/util/verity_util.h"

namespace gvisor {
namespace testing {

namespace {

#define BUF_SIZE 1024

struct linux_dirent {
  uint64_t d_ino;           // Inode number
  int64_t d_off;            // Offset to next linux_dirent64
  unsigned short d_reclen;  // NOLINT, Length of this linux_dirent64
  unsigned char d_type;     // NOLINT, File type
  char d_name[0];           // Filename (null-terminated)
};

class GetDentsTest : public ::testing::Test {
 protected:
  void SetUp() override {
    // Verity is implemented in VFS2.
    SKIP_IF(IsRunningWithVFS1());

    SKIP_IF(!ASSERT_NO_ERRNO_AND_VALUE(HaveCapability(CAP_SYS_ADMIN)));
    // Mount a tmpfs file system, to be wrapped by a verity fs.
    tmpfs_dir_ = ASSERT_NO_ERRNO_AND_VALUE(TempPath::CreateDir());
    ASSERT_THAT(mount("", tmpfs_dir_.path().c_str(), "tmpfs", 0, ""),
                SyscallSucceeds());

    // Create a new file in the tmpfs mount.
    file_ = ASSERT_NO_ERRNO_AND_VALUE(
        TempPath::CreateFileWith(tmpfs_dir_.path(), kContents, 0777));
    filename_ = Basename(file_.path());
  }

  TempPath tmpfs_dir_;
  TempPath file_;
  std::string filename_;
};

TEST_F(GetDentsTest, GetDents) {
  std::string verity_dir =
      ASSERT_NO_ERRNO_AND_VALUE(MountVerity(tmpfs_dir_.path(), filename_));

  auto const verity_fd =
      ASSERT_NO_ERRNO_AND_VALUE(Open(verity_dir, O_RDONLY, 0777));

  char buf[BUF_SIZE];
  int bytes_read = syscall(SYS_getdents64, verity_fd.get(), buf, BUF_SIZE);
  EXPECT_GT(bytes_read, 0);

  std::vector<std::string> children;
  for (int bpos = 0; bpos < bytes_read;) {
    struct linux_dirent *d =
        reinterpret_cast<struct linux_dirent *>(buf + bpos);
    children.push_back(d->d_name);
    bpos += d->d_reclen;
  }

  // Verify the returned contents. It should contain ".", "..", and |filename_|.
  ASSERT_EQ(children.size(), 3);
  EXPECT_THAT(children[0], ::testing::StrEq("."));
  EXPECT_THAT(children[1], ::testing::StrEq(".."));
  EXPECT_THAT(children[2], ::testing::StrEq(filename_));
}

TEST_F(GetDentsTest, Deleted) {
  std::string verity_dir =
      ASSERT_NO_ERRNO_AND_VALUE(MountVerity(tmpfs_dir_.path(), filename_));

  EXPECT_THAT(unlink(JoinPath(tmpfs_dir_.path(), filename_).c_str()),
              SyscallSucceeds());

  auto const verity_fd =
      ASSERT_NO_ERRNO_AND_VALUE(Open(verity_dir, O_RDONLY, 0777));

  char buf[BUF_SIZE];
  EXPECT_THAT(syscall(SYS_getdents64, verity_fd.get(), buf, BUF_SIZE),
              SyscallFailsWithErrno(EIO));
}

TEST_F(GetDentsTest, Renamed) {
  std::string verity_dir =
      ASSERT_NO_ERRNO_AND_VALUE(MountVerity(tmpfs_dir_.path(), filename_));

  std::string new_file_name = "renamed-" + filename_;
  EXPECT_THAT(rename(JoinPath(tmpfs_dir_.path(), filename_).c_str(),
                     JoinPath(tmpfs_dir_.path(), new_file_name).c_str()),
              SyscallSucceeds());

  auto const verity_fd =
      ASSERT_NO_ERRNO_AND_VALUE(Open(verity_dir, O_RDONLY, 0777));

  char buf[BUF_SIZE];
  EXPECT_THAT(syscall(SYS_getdents64, verity_fd.get(), buf, BUF_SIZE),
              SyscallFailsWithErrno(EIO));
}

}  // namespace

}  // namespace testing
}  // namespace gvisor
