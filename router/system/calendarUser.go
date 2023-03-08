package system

import (
	v1 "calendar_service/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CalendarUserRouter struct {
}

func (u CalendarUserRouter) InitCalendarUserRouter(Router *gin.RouterGroup) {
	calendarUserRouter := Router.Group("user")
	calendarUserApi := v1.ApiGroupApp.CalendarUserApi
	calendarUserRouter.POST("create", calendarUserApi.CreateUser)
	calendarUserRouter.POST("list", calendarUserApi.GetUser)
	calendarUserRouter.POST("update", calendarUserApi.UpdateUser)
	calendarUserRouter.POST("delete", calendarUserApi.DeleteUser)
}
