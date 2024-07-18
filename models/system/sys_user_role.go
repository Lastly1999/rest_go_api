package models

type SysUserRole struct {
	//
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	//
	UserId int64 `json:"userId" form:"userId" gorm:"primaryKey"`
	//
	RoleId int64 `json:"roleId" form:"roleId" gorm:"primaryKey"`
}
