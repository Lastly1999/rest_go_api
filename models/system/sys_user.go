package system

type SysUser struct {
	Id        int64  `json:"id" form:"id"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Email     string `json:"email" form:"email"`
	NikeName  string `json:"nike_name" form:"nike_name"`
	Gender    int64  `json:"gender" form:"gender"`
	Phone     string `json:"phone" form:"phone"`
	Avatar    string `json:"avatar" form:"avatar"`
	Status    int64  `json:"status" form:"status"`
	DelFlag   int64  `json:"del_flag" form:"del_flag"`
	Remark    string `json:"remark" form:"remark"`
	DeptId    string `json:"dept_id" form:"dept_id"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}
