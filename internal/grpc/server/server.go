package server

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/douyu/jupiter/pkg/xlog"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"jupiter-demo/internal/demo/greeter"
)

type ServerEngine struct {
	jupiter.Application
}

func NewEngine() *ServerEngine {
	eng := &ServerEngine{}
	eng.SetGovernor("127.0.0.1:9092")
	if err := eng.Startup(
		eng.gRPC,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

func (eng *ServerEngine) gRPC() error {
	server := xgrpc.StdConfig("direct").Build()
	helloworld.RegisterGreeterServer(server.Server, new(greeter.Greeter))
	return eng.Serve(server)
}
