"""Helpers for test consistency.

These targets can be used in BUILD files in order to automatically declare a
group of targets, declare usage of those targets, and assert completeness.

For example, to declare target in foo/BUILD:

  targets(
    name = "targets",
    targets = existing_rules(),
  )

To declare usage of a target in bar/BUILD (*):

  targets(
    name = "foo_usage",
    targets = ["//foo:baz"],
  )

To check for coverage in bar/BUILD of foo/BUILD:

  targets_check(
    name = "targets_check",
    targets = ["//foo:targets"]
    tests = existing_rules(),
  )

Note that these failures will all manifest at the analysis phase (i.e. as build
failures), rather than as runtime test failures. This is make this as simple and
fast to catch these inconsistencies as possible.

(*) Note that this would normally be done as part of another rule implementation,
which can emit a SyscallTargetInfo provider to note this coverage.
"""

SyscallTargetInfo = provider(
    "information about declared system call targets or coverage",
    fields = {
        "targets": "test targets",
    },
)

def _targets_impl(ctx):
    return [SyscallTargetInfo(
        targets = ctx.attr.targets,
    )]

targets = rule(
    implementation = _targets_impl,
    attrs = {
        "targets": attr.label_list(doc = "Targets created or covered.", allow_empty = True),
    },
)

def _targets_check_impl(ctx):
    # Aggregate all coverage.
    covered_labels = []
    for target in ctx.attr.tests:
        if SyscallTargetInfo in target:
            covered_labels += [
                t.label
                for t in target[SyscallTargetInfo].targets
                if not t.label in covered_labels
            ]

    # Compute ignored labels.
    ignored_labels = [target.label for target in ctx.attr.ignore]

    # Check all covered.
    for target in ctx.attr.targets:
        all_targets = target[SyscallTargetInfo].targets
        for target in all_targets:
            # If the target is not the covered labels and not ignored, fail.
            if not target.label in covered_labels and not target.label in ignored_labels:
                fail("Expected coverage of %s, not found." % target.label)

            # On the other hand, if the target *is* covered and *is* ignored, fail.
            if target.label in covered_labels and target.label in ignored_labels:
                fail("Coverage of %s is provided, but is ignored." % target.label)
    return []

targets_check = rule(
    implementation = _targets_check_impl,
    attrs = {
        "targets": attr.label_list(doc = "Targets to cover.", allow_empty = True),
        "tests": attr.label_list(doc = "Tests created.", allow_empty = True),
        "ignore": attr.label_list(doc = "Tests to ignore.", allow_empty = True),
    },
)
