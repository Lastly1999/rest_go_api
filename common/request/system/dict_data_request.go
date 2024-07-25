package request

import "resetgoapi.com/go-admin-api/common/request"

type CreateDictDataRequest struct {
	// 字典标签
	DictLabel string `json:"dictLabel" form:"dictLabel"`
	// 字典键值
	DictValue string `json:"dictValue" form:"dictValue"`
	// 字典排序
	DictSort int64 `json:"dictSort" form:"dictSort"`
	// 创建者
	CreateBy string `json:"createBy" form:"createBy"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
type UpdateDictDataRequest struct {
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	// 字典标签
	DictLabel string `json:"dictLabel" form:"dictLabel"`
	// 字典键值
	DictValue string `json:"dictValue" form:"dictValue"`
	// 字典排序
	DictSort int64 `json:"dictSort" form:"dictSort"`
	// 创建者
	CreateBy string `json:"createBy" form:"createBy"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
type DictDataListRequest struct {
	request.PageRequest
}
