package system

import (
	"calendar_service/model/common/response"
	"calendar_service/system"
	"github.com/gin-gonic/gin"
)

type CalendarEventApi struct {
}

func (event *CalendarEventApi) CreateEvent(ctx *gin.Context) {
	var calendarEvent system.CalendarEvent
	// ShouldBindJSON 绑定json参数
	err := ctx.ShouldBindJSON(&calendarEvent)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

}