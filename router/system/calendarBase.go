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
	calendarUserApi := v1.ApiGroupApp.CalendarUserApi

	calendarBaseRouter.POST("login", calendarBaseApi.Login)
	calendarBaseRouter.POST("register", calendarUserApi.RegisterUser)
}
