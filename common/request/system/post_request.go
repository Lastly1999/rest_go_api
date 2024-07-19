package request

import "resetgoapi.com/rest_go_api/common/request"

type PostListRequest struct {
	request.PageRequest
}

type CreatePostRequest struct {
	// 岗位编码
	PostCode string `json:"postCode" form:"postCode"`
	// 岗位名称
	PostName string `json:"postName" form:"postName"`
	// 岗位排序
	PostSort int64 `json:"postSort" form:"postSort"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}

type UpdatePostRequest struct {
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	// 岗位编码
	PostCode string `json:"postCode" form:"postCode"`
	// 岗位名称
	PostName string `json:"postName" form:"postName"`
	// 岗位排序
	PostSort int64 `json:"postSort" form:"postSort"`
	// 启用状态 0 启用 1 禁用
	Status bool `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}
