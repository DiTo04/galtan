load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@k8s_deploy//:defaults.bzl", "k8s_deploy")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")
load("@io_bazel_rules_docker//container:image.bzl", "container_image")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/DiTo04/galtan/server",
    visibility = ["//visibility:private"],
    deps = [
        "//server/data:go_default_library",
        "//server/proccessing:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "go_image",
    embed = [":go_default_library"],
    goos = "linux",
    pure = "on",
    visibility = ["//visibility:public"],
)

pkg_tar(
    name = "static",
    srcs = [
        "//server/static",
    ],
    package_dir = "/static/",
    strip_prefix = "./static",
)

container_image(
    name = "image",
    base = ":go_image",
    tars = [":static"],
)

k8s_deploy(
    name = "deploy",
    images = {
        "gcr.io/spexflix/galtan:dev": ":image",
    },
    template = ":k8s.yaml",
)

go_test(
    name = "go_default_test",
    srcs = ["main_test.go"],
    embed = [":go_default_library"],
    deps = ["//server/data:go_default_library"],
)
