package initialize

import (
	"calendar_service/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// 实现路由组

	Router := gin.New()
	// 以下属性注意大小写
	// 大写可以被外部访问
	// 小写无法被外部访问
	calendarEventRouter := router.GroupApp.CalendarEventRouter

	PrivateGroup := Router.Group("")
	calendarEventRouter.InitCalendarEventRouter(PrivateGroup)
	return Router
}
