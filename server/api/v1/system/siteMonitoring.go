package system

import (
	"calendar_service/model/common/response"
	"calendar_service/system"
	"github.com/gin-gonic/gin"
)

type SiteMonitoringApi struct {
}

func (receiver *SiteMonitoringApi) AddRecord(ctx *gin.Context) {
	var record system.SiteMonitoring
	err := ctx.ShouldBindJSON(&record)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	ip := ctx.Request.Header.Get("X-Forward-For")
	record.SourceIP = ip
	err = siteMonitoringRecordService.AddRecord(record)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("创建记录", ctx)
}

func (receiver *SiteMonitoringApi) GetRecords(ctx *gin.Context) {
	records, err := siteMonitoringRecordService.GetRecords()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed(records, "获取成功", ctx)
}
