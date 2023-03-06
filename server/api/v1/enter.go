package v1

import "calendar_service/server/api/v1/system"

type ApiGroup struct {
	CalendarEventApi system.CalendarEventApi
	CalendarGroupApi system.CalendarGroupApi
}

var ApiGroupApp = new(ApiGroup)
