package system

import (
	"calendar_service/global"
)

type CalendarEvent struct {
	global.CalendarEventModel        // 包含ID、创建时间、更新时间、删除时间
	Name                      string // 名称
	Start                     uint64 // 开始时间
	End                       uint64 // 结束时间
	Color                     string // 颜色
	AllDay                    bool   // 全天
	Author                    string // 创建者
	Location                  string // 地点
	Personnel                 string // 人员
}
