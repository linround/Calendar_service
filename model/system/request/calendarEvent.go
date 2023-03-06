package request

type SearchEventListParams struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

type SearchCalendarGroupParams struct {
	GroupID uint64 `json:"groupId"`
}
