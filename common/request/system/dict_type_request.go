package request

import "resetgoapi.com/go-admin-api/common/request"

type CreateDictTypeRequest struct {
	// 字典名称
	DictName string `json:"dictName" form:"dictName"`
	// 字典类型
	DictType string `json:"dictType" form:"dictType"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
type UpdateDictTypeRequest struct {
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	// 字典名称
	DictName string `json:"dictName" form:"dictName"`
	// 字典类型
	DictType string `json:"dictType" form:"dictType"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
type DictTypeListRequest struct {
	request.PageRequest
}
