package request

type SearchPictureListParams struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
	BaseClaims
}
