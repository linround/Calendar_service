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

	err = global.CalendarDB.Model(&system.CalendarEvent{}).Where("event_id = ?", event.EventID).Updates(map[string]interface{}{
		"event_name":      event.EventName,
		"start":           event.Start,
		"end":             event.End,
		"event_color":     event.EventColor,
		"all_day":         event.AllDay,
		"user_name":       event.UserName,
		"event_location":  event.EventLocation,
		"event_personnel": event.EventPersonnel,
		"event_timed":     event.EventTimed,
	}).Error
	if err != nil {
		return
	}
	return
}

func (calendarService *CalendarEventService) DeleteEvent(event system.ApiCalendarEvent) (err error) {
	// 模型包含了一个 gorm.deletedAt 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力
	// 使用软删除的方式
	err = global.CalendarDB.Delete(&system.CalendarEvent{}, event.EventID).Error
	if err != nil {
		return
	}
	return
}
