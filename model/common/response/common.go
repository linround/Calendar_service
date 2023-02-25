package response

type Array interface {
}

type ListResult struct {
	List Array `json:"list"`
}
