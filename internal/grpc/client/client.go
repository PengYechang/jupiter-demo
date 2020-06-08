package client

import (
	"context"
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/client/grpc"
	"github.com/douyu/jupiter/pkg/xlog"
	"google.golang.org/grpc/examples/helloworld/helloworld"
	"time"
)

type ClientEngine struct {
	jupiter.Application
}

func NewEngine() *ClientEngine {
	eng := &ClientEngine{}
	eng.SetGovernor("127.0.0.1:9999")
	if err := eng.Startup(
		eng.gRPC,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

func (cli *ClientEngine) gRPC() error {
	conn := grpc.StdConfig("direct").Build()
	client := helloworld.NewGreeterClient(conn)
	go func() {
		for{
			resp,err := client.SayHello(context.Background(),&helloworld.HelloRequest{
				Name: "pyc",
			})
			if err != nil {
				xlog.Error(err.Error())
			}else {
				xlog.Info("receive", xlog.String("resp",resp.Message))
			}
			time.Sleep(1 * time.Second)
		}
	}()
	return nil
}