package system

import (
	"calendar_service/model/common/response"
	"calendar_service/model/system/request"
	"calendar_service/system"
	"github.com/gin-gonic/gin"
)

type CalendarGroupApi struct {
}

func (group *CalendarGroupApi) CreateGroup(ctx *gin.Context) {
	var calendarGroup system.CalendarGroup
	err := ctx.ShouldBindJSON(&calendarGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = calendarGroupService.CreateGroup(calendarGroup)
	if err != nil {
		response.FailWithMessage("创建失败", ctx)
		return
	}
	response.OkWithMessage("创建成功", ctx)

}

func (group *CalendarGroupApi) GetGroupList(ctx *gin.Context) {
	var params request.SearchCalendarGroupParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	list, err := calendarGroupService.GetGroup(params)
	if err != nil {
		response.FailWithMessage("获取失败", ctx)
		return
	}
	response.OkWithDetailed(response.ListResult{List: list}, "获取成功", ctx)

}

func (group *CalendarGroupApi) UpdateGroup(ctx *gin.Context) {
	var params system.ApiCalendarGroup
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = calendarGroupService.UpdateGroup(params)
	if err != nil {
		response.FailWithMessage("更新失败", ctx)
		return
	}
	response.OkWithMessage("更新成功", ctx)
}
func (group *CalendarGroupApi) DeleteGroup(ctx *gin.Context) {
	var params system.ApiCalendarGroup
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = calendarGroupService.DeleteGroup(params)
	if err != nil {
		response.FailWithMessage("删除失败", ctx)
		return
	}
	response.OkWithMessage("删除成功", ctx)

}
