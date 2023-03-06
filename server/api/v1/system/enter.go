package system

import "calendar_service/service"

type ApiGroup struct {
	CalendarEventApi
	CalendarGroupApi
}

var (
	calendarEventService = service.GroupApp.CalendarEventService
	calendarGroupService = service.GroupApp.CalendarGroupService
)
