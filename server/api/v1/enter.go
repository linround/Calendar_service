package v1

import "calendar_service/server/api/v1/system"

// 将API导出

type ApiGroup struct {
	CalendarEventApi system.CalendarEventApi
	CalendarGroupApi system.CalendarGroupApi
	CalendarUserApi  system.CalendarUserApi
}

var ApiGroupApp = new(ApiGroup)
