package system

import (
	v1 "calendar_service/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CalendarGroupRouter struct {
}

func (g *CalendarGroupRouter) InitCalendarGroupRouter(Router *gin.RouterGroup) {
	calendarGroupRouter := Router.Group("group")
	CalendarGroupApi := v1.ApiGroupApp.CalendarGroupApi
	calendarGroupRouter.POST("create", CalendarGroupApi.CreateGroup)
	calendarGroupRouter.POST("list", CalendarGroupApi.GetGroupList)
	calendarGroupRouter.POST("update", CalendarGroupApi.UpdateGroup)
	calendarGroupRouter.POST("delete", CalendarGroupApi.DeleteGroup)
}
