package core

import (
	"calendar_service/global"
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
	// 从db中加载jwt数据
	if global.CalendarDB != nil {
		//	这里主要是读取jwt名单，从而设置黑名单缓存
	}

	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", port)
	s := initServer(address, Router)
	fmt.Println("服务运行在端口：", port)
	s.ListenAndServe().Error()
}
