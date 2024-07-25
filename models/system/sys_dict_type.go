package models

import (
	"resetgoapi.com/go-admin-api/models"
)

type SysDictType struct {
	models.BaseModel
	// 字典名称
	DictName string `json:"dictName" form:"dictName"`
	// 字典类型
	DictType string `json:"dictType" form:"dictType"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
