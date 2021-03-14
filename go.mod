module github.com/c3sr/registry

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.29.1
)

require (
	github.com/c3sr/config v1.0.1
	github.com/c3sr/libkv v1.0.0
	github.com/c3sr/logger v1.0.1
	github.com/c3sr/serializer v1.0.0
	github.com/c3sr/utils v1.0.0
	github.com/c3sr/vipertags v1.0.0
	github.com/k0kubun/pp/v3 v3.0.7
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc/examples v0.0.0-20210312231957-21976fa3e38a // indirect
)
