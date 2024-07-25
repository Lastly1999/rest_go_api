package models

import (
	"resetgoapi.com/go-admin-api/models"
	"time"
)

type SysLoginLog struct {
	models.BaseModel
	// '用户名'
	UserName string `json:"userName" form:"userName"`
	// '状态 0成功 1失败'
	Status int64 `json:"status" form:"status"`
	// '登录IP地址'
	Ipaddr string `json:"ipaddr" form:"ipaddr"`
	// '登录地点'
	LoginLocation string `json:"loginLocation" form:"loginLocation"`
	// '浏览器类型'
	Browser string `json:"browser" form:"browser"`
	// '操作系统'
	Os string `json:"os" form:"os"`
	// '登录系统'
	Platform string `json:"platform" form:"platform"`
	// '登录时间'
	LoginTime time.Time `json:"loginTime" form:"loginTime"`
}
