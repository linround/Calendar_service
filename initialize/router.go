package initialize

import (
	"calendar_service/router"
	"calendar_service/server/middleware"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// 实现路由组

	Router := gin.New()
	// 以下属性注意大小写
	// 大写可以被外部访问
	// 小写无法被外部访问
	// 配置跨域请求
	Router.Use(middleware.Cors())

	/**********************公有路由开始************************/
	calendarBaseRouter := router.GroupApp.CalendarBaseRouter
	calendarDBRouter := router.GroupApp.CalendarDBRouter

	PublicGroup := Router.Group("")

	calendarBaseRouter.InitCalendarBaseRouter(PublicGroup)
	calendarDBRouter.InitCalendarDBRouter(PublicGroup)
	/**********************公有路由结束************************/

	/**********************私有路由开始************************/
	calendarEventRouter := router.GroupApp.CalendarEventRouter
	calendarGroupRouter := router.GroupApp.CalendarGroupRouter
	calendarUserRouter := router.GroupApp.CalendarUserRouter

	PrivateGroup := Router.Group("")
	// 私有路由添加校验
	PrivateGroup.Use(middleware.JWTAuth())

	calendarEventRouter.InitCalendarEventRouter(PrivateGroup)
	calendarGroupRouter.InitCalendarGroupRouter(PrivateGroup)
	calendarUserRouter.InitCalendarUserRouter(PrivateGroup)
	/**********************私有路由结束************************/
	return Router
}
