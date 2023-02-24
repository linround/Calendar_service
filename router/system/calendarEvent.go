package system

import (
	v1 "calendar_service/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CalendarEventRouter struct {
}

func (s *CalendarEventRouter) InitCalendarEventRouter(Router *gin.RouterGroup) {
	calendarEventRouter := Router.Group("event")
	CalendarEventApi := v1.ApiGroupApp.CalendarEventApi
	calendarEventRouter.POST("create", CalendarEventApi.CreateEvent)
	calendarEventRouter.POST("list", CalendarEventApi.GetEventList)
	calendarEventRouter.POST("update", CalendarEventApi.UpdateEvent)
}
