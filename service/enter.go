package service

import "calendar_service/service/system"

// 导出service

type Group struct {
	CalendarEventService system.CalendarEventService
	CalendarGroupService system.CalendarGroupService
	CalendarUserService  system.CalendarUserService
	CalendarBaseService  system.CalendarBaseService
}

var GroupApp = new(Group)
