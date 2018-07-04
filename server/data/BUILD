load("@io_bazel_rules_go//go:def.bzl", "go_library")

java_library(
    name = "data",
)

go_library(
    name = "go_default_library",
    srcs = [
        "data.go",
        "jsonResultStore.go",
    ],
    importpath = "github.com/DiTo04/galtan/server/data",
    visibility = ["//visibility:public"],
)
