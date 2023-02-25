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
	// 记录用户更新之前的旧数据
	err = global.CalendarDB.Model(&system.CalendarEvent{}).Where("id = ?", event.ID).First(&oldEvent).Error
	if err != nil {
		return
	}
	err = global.CalendarDB.Model(&system.CalendarEvent{}).Where("id = ?", event.ID).Updates(map[string]interface{}{
		"name":      event.Name,
		"start":     event.Start,
		"end":       event.End,
		"color":     event.Color,
		"all_day":   event.AllDay,
		"author":    event.Author,
		"location":  event.Location,
		"personnel": event.Personnel,
		"timed":     event.Timed,
	}).Error
	if err != nil {
		return
	}
	err = global.CalendarDB.Save(&event).Error
	return
}

func (calendarService *CalendarEventService) DeleteEvent(event system.ApiCalendarEvent) (err error) {
	// 模型包含了一个 gorm.deletedAt 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力
	// 使用软删除的方式
	err = global.CalendarDB.Delete(&system.CalendarEvent{}, event.ID).Error
	if err != nil {
		return
	}
	return
}
