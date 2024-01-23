package system

import (
	v1 "calendar_service/server/api/v1"
	"github.com/gin-gonic/gin"
)

type SiteMonitoringRouter struct {
}

func (receiver *SiteMonitoringRouter) InitSiteMonitoringRouter(Router *gin.RouterGroup) {
	siteMonitoringRouter := Router.Group("monitoring")
	siteMonitoringApi := v1.ApiGroupApp.SiteMonitoringApi
	siteMonitoringRouter.POST("add", siteMonitoringApi.AddRecord)
	siteMonitoringRouter.GET("records", siteMonitoringApi.GetRecords)
}
