package system

import "calendar_service/global"

type category string

const (
	Visit category = "visit" // 站点访问统计
)

type SiteMonitoring struct {
	global.CalendarEventModel `yaml:",inline"`
	RecordID                  uint64   `json:"recordId" gorm:"not null;unique;primary_key;comment:记录ID;size:90"`
	TimePoint                 uint64   `json:"timePoint" gorm:"访问网站的时间点"`
	SourceIP                  string   `json:"sourceIP" gorm:"访问请求的IP"`
	SourceSite                string   `json:"sourceSite" gorm:"当前网站来源"`
	Category                  category `json:"category" gorm:"当前的记录的类型"`
}

type SiteMonitoringRecords struct {
	Total int64 `json:"total" gorm:"统计总条数"`
}

func (s SiteMonitoring) TableName() string {
	return "site_monitoring"
}
