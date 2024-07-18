package models

import "resetgoapi.com/rest_go_api/models"

type SysMenu struct {
	models.BaseModel
	// 菜单名称
	MenuName string `json:"menuName" form:"menuName"`
	// 父菜单ID
	ParentId int64 `json:"parentId" form:"parentId"`
	// 父菜单名称
	ParentName string `json:"parentName" form:"parentName"`
	// 显示排序
	MenuSort int64 `json:"menuSort" form:"menuSort"`
	// 菜单类型 0目录 1菜单 2按钮
	MenuType bool `json:"menuType" form:"menuType"`
	// 菜单状态 0显示 1隐藏
	Visible bool `json:"visible" form:"visible"`
	// 权限标识
	Perms string `json:"perms" form:"perms"`
	// 菜单图标
	MenuIcon string `json:"menuIcon" form:"menuIcon"`
	// 备注
	Remark string `json:"remark" form:"remark"`
	// 是否缓存 0缓存 1不缓存
	IsCache bool `json:"isCache" form:"isCache"`
	// 是否外链 0是 1否
	IsFrame bool `json:"isFrame" form:"isFrame"`
	// 组件路径
	Component string `json:"component" form:"component"`
	// 路由地址
	Router string `json:"router" form:"router"`
	// 路由参数
	Query string `json:"query" form:"query"`
	// 菜单状态 0正常 1停用
	Status bool `json:"status" form:"status"`
}
