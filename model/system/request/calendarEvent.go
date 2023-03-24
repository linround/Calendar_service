package request

type SearchEventListParams struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
	BaseClaims
}
