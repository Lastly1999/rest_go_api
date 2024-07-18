package models

type SysRoleDept struct {
	//
	Id int64 `json:"id" form:"id"`
	// '角色ID'
	RoleId int64 `json:"roleId" form:"roleId"`
	// '部门ID'
	DeptId int64 `json:"deptId" form:"deptId"`
}
