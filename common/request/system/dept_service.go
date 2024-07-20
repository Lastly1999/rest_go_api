package request

import "resetgoapi.com/rest_go_api/common/request"

type CreateDeptRequest struct {
	// 部门名称
	DeptName string `json:"deptName" form:"deptName"`
	// 父部门ID
	ParentId int64 `json:"parentId" form:"parentId"`
	// 显示排序
	DeptSort int64 `json:"deptSort" form:"deptSort"`
	// 负责人
	Leader string `json:"leader" form:"leader"`
	// 联系电话
	Phone string `json:"phone" form:"phone"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
type UpdateDeptRequest struct {
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	// 部门名称
	DeptName string `json:"deptName" form:"deptName"`
	// 父部门ID
	ParentId int64 `json:"parentId" form:"parentId"`
	// 显示排序
	DeptSort int64 `json:"deptSort" form:"deptSort"`
	// 负责人
	Leader string `json:"leader" form:"leader"`
	// 联系电话
	Phone string `json:"phone" form:"phone"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
type DeptListRequest struct {
	request.PageRequest
}
