package system

import "calendar_service/global"

type category string

const (
	Visit category = "visit" // 站点访问统计
)

type ApiSiteMonitoring struct {
	timePoint uint64   `json:"timePoint" gorm:"访问网站的时间点"`
	sourceIP  string   `json:"sourceIP" gorm:"访问请求的IP"`
	category  category `json:"category" gorm:"当前的记录的类型"`
}

type SiteMonitoring struct {
	global.CalendarEventModel `yaml:",inline"`
	ApiSiteMonitoring         `yaml:",inline"`
}

func (s SiteMonitoring) TableName() string {
	return "site_monitoring"
}
