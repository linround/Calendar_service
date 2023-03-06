package service

import "calendar_service/service/system"

type Group struct {
	CalendarEventService system.CalendarEventService
	CalendarGroupService system.CalendarGroupService
}

var GroupApp = new(Group)
