package system

import (
	v1 "calendar_service/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CalendarEventRouter struct {
}

func (s *CalendarEventRouter) InitCalendarEventRouter(Router *gin.RouterGroup) {
	calendarEventRouter := Router.Group("CalendarEvent")
	CalendarEventApi := v1.ApiGroupApp.CalendarEventApi
	calendarEventRouter.POST("createEvent", CalendarEventApi.CreateEvent)
}
