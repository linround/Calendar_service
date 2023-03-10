package system

import (
	v1 "calendar_service/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CalendarBaseRouter struct {
}

func (s *CalendarBaseRouter) InitCalendarBaseRouter(Router *gin.RouterGroup) {
	calendarBaseRouter := Router.Group("base")
	calendarBaseApi := v1.ApiGroupApp.CalendarBaseApi
	calendarBaseRouter.POST("login", calendarBaseApi.Login)
}
