workspace(name = "binchencoder_gateway_proto")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6776d68ebb897625dead17ae510eac3d5f6342367327875210df44dbe2aeeb19",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.17.1/rules_go-0.17.1.tar.gz"],
)
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

go_repository(
    name = "rules_proto",
    commit = "b0cc14be5da05168b01db282fe93bdf17aa2b9f4",
    importpath = "github.com/bazelbuild/rules_proto",
)
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
rules_proto_dependencies()
rules_proto_toolchains()

go_repository(
    name = "binchencoder_third_party_java",
    commit = "8a4b915c4944f616e2ad34e16cf096d9c30a6e96",
    importpath = "github.com/binchencoder/third-party-java",
)

go_repository(
    name = "grpc_ecosystem_grpc_gateway",
    commit = "52a539ec166cf7380660e9c7ebde5206169d385d",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
)

go_repository(
    name = "com_github_fatih_color",
    commit = "3f9d52f7176a6927daacff70a3e8d1dc2025c53e",
    importpath = "github.com/fatih/color",
)

go_repository(
    name = "com_github_klauspost_compress",
    commit = "ae52aff18558bd92cbe681549bfe9e8cbffd5903",
    importpath = "github.com/klauspost/compress",
)

go_repository(
    name = "com_github_klauspost_cpuid",
    commit = "05a8198c0f5a27739aec358908d7e12c64ce6eb7",
    importpath = "github.com/klauspost/cpuid",
)

go_repository(
    name = "com_github_golang_net",
    commit = "4829fb13d2c62012c17688fa7f629f371014946d",
    importpath = "github.com/golang/net",
)

# Also define in Gopkg.toml
go_repository(
    name = "org_golang_google_genproto",
    commit = "383e8b2c3b9e36c4076b235b32537292176bae20",
    importpath = "google.golang.org/genproto",
)

# Also define in Gopkg.toml
go_repository(
    name = "com_github_rogpeppe_fastuuid",
    commit = "6724a57986aff9bff1a1770e9347036def7c89f6",
    importpath = "github.com/rogpeppe/fastuuid",
)

# Also define in Gopkg.toml
go_repository(
    name = "in_gopkg_resty_v1",
    commit = "fa5875c0caa5c260ab78acec5a244215a730247f",
    importpath = "gopkg.in/resty.v1",
)

# Also define in Gopkg.toml
go_repository(
    name = "com_github_ghodss_yaml",
    commit = "0ca9ea5df5451ffdf184b4428c902747c2c11cd7",
    importpath = "github.com/ghodss/yaml",
)

# Also define in Gopkg.toml
go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "eb3733d160e74a9c7e442f435eb3bea458e1d19f",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "com_github_bazelbuild_buildtools",
    importpath = "github.com/bazelbuild/buildtools",
    commit = "36bd730dfa67bff4998fe897ee4bbb529cc9fbee",
)
load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")
buildifier_dependencies()
