package initialize

import (
	"calendar_service/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// 实现路由组

	Router := gin.New()
	// 以下属性注意大小写
	// 大写可以被外部访问
	// 小写无法被外部访问
	CalendarEventRouter := router.GroupApp.CalendarEventRouter
	fmt.Println(CalendarEventRouter)
	return Router
}
