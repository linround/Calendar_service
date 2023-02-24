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

func (calendarService *CalendarEventService) GetEventList(params request.SearchEventListParams) (list []system.ApiCalendarEvent, err error) {
	// 通过model指定了对应的表
	// 通过sys.ApiCalendarEvent 指定了要查询的相关字段
	// 最终填入list即可
	err = global.CalendarDB.Model(&system.CalendarEvent{}).Find(&list).Error
	return
}

func (calendarService *CalendarEventService) UpdateEvent(event system.ApiCalendarEvent) (err error) {
	var oldEvent system.ApiCalendarEvent
	err = global.CalendarDB.Model(&system.CalendarEvent{}).Where("id = ?", event.ID).First(&oldEvent).Error
	if err != nil {
		return
	}
	err = global.CalendarDB.Model(&system.CalendarEvent{}).Where("id = ?", event.ID).Updates(map[string]interface{}{
		"name": event.Name,
	}).Error
	if err != nil {
		return
	}
	err = global.CalendarDB.Save(&event).Error
	return
}
