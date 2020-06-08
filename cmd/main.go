package main

import (
	"github.com/douyu/jupiter/pkg/worker"
	"github.com/douyu/jupiter/pkg/xlog"
	"jupiter-demo/internal/grpc/client"
	"jupiter-demo/internal/grpc/server"
	"os"
)

func main() {
	input := os.Args[2]
	var wk worker.Worker
	switch input {
	case "client":
		os.Args[1] = "--config=internal/grpc/client/config.toml"
		wk = client.NewEngine()
	case "server":
		os.Args[1] = "--config=internal/grpc/server/config.toml"
		wk = server.NewEngine()
	default:
		xlog.Error("input err!")
		return
	}
	if err := wk.Run(); err != nil {
		xlog.Error(err.Error())
	}
}
