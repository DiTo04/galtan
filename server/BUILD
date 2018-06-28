load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@k8s_deploy//:defaults.bzl", "k8s_deploy")

go_library(
    name = "go_default_library",
    srcs = [
        "jsonResultStore.go",
        "main.go",
    ],
    importpath = "github.com/DiTo04/spexflix/server",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_gorilla_mux//:go_default_library",
    ],
)

go_binary(
    name = "server",
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
    name = "deploy",
    images = {
        "gcr.io/spexflix/galtan:dev": ":image",
    },
    template = ":k8s.yaml",
)
