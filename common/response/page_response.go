package response

type PageResponse struct {
	Total int64 `json:"total"`
	List  interface{}
}