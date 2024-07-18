package models

type SysRoleMenu struct {
	//
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	//
	RoleId int64 `json:"roleId" form:"roleId" gorm:"primaryKey"`
	//
	MenuId int64 `json:"menuId" form:"menuId" gorm:"primaryKey"`
}
