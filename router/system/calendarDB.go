package system

import (
	v1 "calendar_service/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CalendarDBRouter struct {
}

func (d *CalendarDBRouter) InitCalendarDBRouter(Router *gin.RouterGroup) {
	calendarDBRouter := Router.Group("init")
	calendarDBApi := v1.ApiGroupApp.CalendarDBApi
	calendarDBRouter.POST("initDB", calendarDBApi.InitDB)
	calendarDBRouter.POST("checkDB", calendarDBApi.CheckDB)
}
