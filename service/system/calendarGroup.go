package system

import (
	"calendar_service/global"
	"calendar_service/model/system/request"
	"calendar_service/system"
)

type CalendarGroupService struct {
}

func (calendarService *CalendarGroupService) CreateGroup(group system.CalendarGroup) (err error) {
	return global.CalendarDB.Create(&group).Error

}
func (calendarService *CalendarGroupService) GetGroup(params request.SearchCalendarGroupParams) (list []system.ApiCalendarGroup, err error) {
	err = global.CalendarDB.Model(&system.CalendarGroup{}).Find(&list).Error
	return
}
func (calendarService *CalendarGroupService) UpdateGroup(group system.ApiCalendarGroup) (err error) {
	err = global.CalendarDB.Model(&system.CalendarGroup{}).Where("group_id = ?", group.GroupID).Updates(map[string]interface{}{
		"user_id":     group.UserID,
		"group_type":  group.GroupType,
		"group_color": group.GroupColor,
		"group_name":  group.GroupName,
	}).Error
	return
}
func (calendarService *CalendarGroupService) DeleteGroup(group system.ApiCalendarGroup) (err error) {
	err = global.CalendarDB.Delete(&system.CalendarGroup{}, group.GroupID).Error
	return
}
