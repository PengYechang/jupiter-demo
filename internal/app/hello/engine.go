package hello

import (
	"github.com/douyu/jupiter"
	"jupiter-demo/internal/demo/greeter"
	"github.com/douyu/jupiter/pkg/server/xecho"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/douyu/jupiter/pkg/util/xgo"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/examples/helloworld/helloworld"
)

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
    	xgo.ParallelWithError(
    			eng.serveHTTP,
    			eng.serveGRPC,
    	),
    ); err != nil {
    	xlog.Panic("startup hello", xlog.Any("err", err))
    }
	return eng
}
func (eng *Engine) serveHTTP() error {
	server := xecho.StdConfig("http").Build()
	server.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(200, "pong")
	})
	eng.Serve(server)
	return nil
}

func (eng *Engine) serveGRPC() error {
	server := xgrpc.StdConfig("grpc").Build()
	helloworld.RegisterGreeterServer(server.Server, new(greeter.Greeter))
	eng.Serve(server)
	return nil
}
