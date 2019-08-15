load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        urls = [
            "https://codeload.github.com/grpc-ecosystem/grpc-gateway/tar.gz/fdf063599d922ec89a70819e2d5b7b4b5c642b92",
        ],
        strip_prefix = "grpc-gateway-fdf063599d922ec89a70819e2d5b7b4b5c642b92",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
        version = "v0.0.0-20160126235308-23def4e6c14b",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        urls = [
            "https://codeload.github.com/grpc/grpc-go/tar.gz/df014850f6dee74ba2fc94874043a9f3f75fbfd8",
        ],
        strip_prefix = "grpc-go-df014850f6dee74ba2fc94874043a9f3f75fbfd8", # v1.17.0, latest as of 2019-01-15
        type = "tar.gz",
        # gazelle args: -go_prefix google.golang.org/grpc -proto disable
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:+t9dhfO+GNOIGJof6kPOAenx7YgrZMTdRPV+EsnPabk=",
        version = "v2.0.0-20170812160011-eb3733d160e7",
    )