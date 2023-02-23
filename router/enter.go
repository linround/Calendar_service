package router

import "calendar_service/router/system"

type Group struct {
	system.RouterGroup
}

var GroupApp = new(Group)
