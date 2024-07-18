package models

type SysUserPost struct {
	//
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	// '用户ID'
	UserId int64 `json:"userId" form:"userId"`
	// '岗位ID'
	PostId int64 `json:"postId" form:"postId"`
}
