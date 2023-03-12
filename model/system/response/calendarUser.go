package response

import "calendar_service/system"

type LoginResponse struct {
	User  system.CalendarUser `json:"user"`
	Token string              `json:"token"`
}
