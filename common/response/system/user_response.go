package response

type GetUserInfoResponse struct {
	Roles    []string `json:"roles"`
	UserId   int64    `json:"userId"`
	Username string   `json:"username"`
	RealName string   `json:"realName"`
	Avatar   string   `json:"avatar"`
	Desc     string   `json:"desc"`
}
