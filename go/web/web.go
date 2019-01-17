package web

import (
	"os"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	cli "github.com/micro/go-plugins/client/grpc"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	srv "github.com/micro/go-plugins/server/grpc"
	"github.com/micro/go-web"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	"github.com/micro/go-plugins/selector/static"
)

// NewService returns a web service for kubernetes
func NewService(opts ...web.Option) web.Service {
	// set grpc transport
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()

	// set the static selector
	os.Setenv("MICRO_SELECTOR", "static")

	client.DefaultClient.Init(
		client.Selector(static.NewSelector()),
	)

	// set kubernetes registry
	os.Setenv("MICRO_REGISTRY", "kubernetes")

	return web.NewService(opts...)
}
