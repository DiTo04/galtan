load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@k8s_deploy//:defaults.bzl", "k8s_deploy")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    command = "fix",
    extra_args = [
        "-build_file_name",
        "BUILD,BUILD.bazel",
    ],
    prefix = "github.com/DiTo04/spexflix",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/DiTo04/spexflix",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_gorilla_mux//:go_default_library",
    ],
)

go_binary(
    name = "galtan",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "image",
    embed = [":go_default_library"],
    goos = "linux",
    pure = "on",
    visibility = ["//visibility:public"],
)

k8s_deploy(
    name = "dev",
    images = {
        "gcr.io/spexflix/galtan:dev": ":image",
    },
    template = ":k8s.yaml",
)