package core

import "calendar_service/initialize"

// 启动相关服务

func RunServer() {
	Router := initialize.Routers()
	println(Router)
}
