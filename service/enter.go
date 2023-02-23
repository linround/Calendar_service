package service

import "calendar_service/service/system"

type Group struct {
	CalendarEventService system.CalendarEventService
}

var GroupApp = new(Group)
