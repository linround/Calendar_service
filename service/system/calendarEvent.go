package system

import (
	"calendar_service/global"
	"calendar_service/system"
)

type CalendarEventService struct {
}

func (calendarService *CalendarEventService) CreateEvent(event system.CalendarEvent) (err error) {
	return global.CalendarDB.Create(&event).Error
}
