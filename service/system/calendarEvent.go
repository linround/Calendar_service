package system

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"calendar_service/system"
)

type CalendarEventService struct {
}

func (calendarService *CalendarEventService) CreateEvent(event system.CalendarEvent) (err error) {
	return global.CalendarDB.Create(&event).Error
}
func (calendarService *CalendarEventService) GetEventList(params request.SearchEventListParams) (eventList []system.CalendarEvent, err error) {
	db := global.CalendarDB.Model(&system.CalendarEvent{})
	db = db.Where("start >= ? AND end < ?", params.StartTime, params.EndTime).Find(&eventList)
	return
}
