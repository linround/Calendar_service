package system

import (
	"calendar_service/model/common/response"
	"calendar_service/model/system/request"
	"calendar_service/system"
	"calendar_service/utils"
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

	// 初始化token相关的参数信息

	userID := utils.GetUserID(ctx)
	UserAccount := utils.GetUserAccount(ctx)
	calendarEvent.UserID = userID
	calendarEvent.UserAccount = UserAccount

	err = calendarEventService.CreateEvent(calendarEvent)
	if err != nil {
		response.FailWithMessage("创建失败", ctx)
		return
	}
	response.OkWithMessage("创建成功", ctx)
}
func (event *CalendarEventApi) GetEventList(ctx *gin.Context) {
	var params request.SearchEventListParams
	err := ctx.ShouldBindJSON(&params)

	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 初始化token相关的参数信息

	userID := utils.GetUserID(ctx)
	UserAccount := utils.GetUserAccount(ctx)
	params.UserID = userID
	params.UserAccount = UserAccount

	list, err := calendarEventService.GetEventList(params)
	if err != nil {
		response.FailWithMessage("获取失败", ctx)
		return
	}
	response.OkWithDetailed(response.ListResult{List: list}, "获取成功", ctx)

}
func (event *CalendarEventApi) UpdateEvent(ctx *gin.Context) {
	var params system.ApiCalendarEvent
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = calendarEventService.UpdateEvent(params)
	if err != nil {
		response.FailWithMessage("更新失败", ctx)
		return
	}
	response.OkWithMessage("更新成功", ctx)
}
func (event *CalendarEventApi) DeleteEvent(ctx *gin.Context) {
	var params system.ApiCalendarEvent
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = calendarEventService.DeleteEvent(params)
	if err != nil {
		response.FailWithMessage("删除失败", ctx)
		return
	}
	response.OkWithMessage("删除成功", ctx)
}
