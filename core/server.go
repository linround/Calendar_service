package core

import (
	"calendar_service/initialize"
	"fmt"
)

type server interface {
	ListenAndServe() error
}

// 启动相关服务

const (
	port int = 8888
)

func RunServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", port)
	s := initServer(address, Router)
	fmt.Println("服务运行在端口：", port)
	s.ListenAndServe().Error()
}
