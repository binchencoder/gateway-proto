module github.com/binchencoder/gateway-proto

go 1.13

replace (
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	google.golang.org/genproto v0.0.0-20200702021140-07506425bd67
	google.golang.org/grpc v1.30.0
)
