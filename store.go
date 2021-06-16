package registry

import (
	"context"

	"github.com/c3sr/libkv"
	"github.com/c3sr/libkv/store"
	"github.com/c3sr/libkv/store/consul"
	"github.com/c3sr/libkv/store/mock"
)

type Store interface {
	store.Store
}

func New(opts ...Option) (Store, error) {
	options := Options{
		Provider:          Config.Provider,
		Endpoints:         cleanupEndpoints(Config.Endpoints),
		Username:          Config.Username,
		Password:          Config.Password,
		Timeout:           Config.Timeout,
		TLSConfig:         nil,
		PersistConnection: true,
		Context:           context.Background(),
	}
	if Config.HeaderTimeoutPerRequest != 0 {
		HeaderTimeoutPerRequest(Config.HeaderTimeoutPerRequest)(&options)
	}
	if Config.Certificate != "" {
		TLSCertificate(Config.Certificate)(&options)
	}
	AutoSync(Config.AutoSync)(&options)
	for _, o := range opts {
		o(&options)
	}
	storeConfig := &store.Config{
		ClientTLS:         &store.ClientTLSConfig{},
		ConnectionTimeout: options.Timeout,
		Username:          options.Username,
		Password:          options.Password,
		TLS:               options.TLSConfig,
		Bucket:            options.Bucket,
		PersistConnection: options.PersistConnection,
		Context:           options.Context,
	}
	if options.Provider == store.Backend("mock") {
		return mock.New(options.Endpoints, storeConfig)
	}
	return libkv.NewStore(options.Provider, options.Endpoints, storeConfig)
}

func init() {
	consul.Register()
}
