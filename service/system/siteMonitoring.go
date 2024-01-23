package system

import (
	"calendar_service/global"
	"calendar_service/system"
)

type SiteMonitoringRecordService struct {
}

func (recordService *SiteMonitoringRecordService) AddRecord(record system.SiteMonitoring) (err error) {
	return global.CalendarDB.Create(&record).Error
}

func (recordService *SiteMonitoringRecordService) GetRecords() (records system.SiteMonitoringRecords, err error) {
	var total int64
	err = global.CalendarDB.Model(&system.SiteMonitoring{}).Count(&total).Error
	records.Total = total
	return
}
