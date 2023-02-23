package system

import "calendar_service/service"

type ApiGroup struct {
	CalendarEventApi
}

var (
	calendarEventService = service.GroupApp.CalendarEventService
)
