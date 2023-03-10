package system

import "calendar_service/service"

// 命名自己的局部service

var (
	calendarEventService = service.GroupApp.CalendarEventService
	calendarGroupService = service.GroupApp.CalendarGroupService
	calendarUserService  = service.GroupApp.CalendarUserService
	calendarBaseService  = service.GroupApp.CalendarBaseService
)
