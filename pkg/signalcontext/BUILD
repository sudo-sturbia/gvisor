load("//tools:defs.bzl", "go_library")

package(licenses = ["notice"])

go_library(
    name = "signalcontext",
    srcs = [
        "context.go",
    ],
    visibility = ["//:sandbox"],
    deps = [
        "//pkg/abi/linux",
        "//pkg/context",
    ],
)
