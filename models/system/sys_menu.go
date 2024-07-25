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
	MenuType int `json:"menuType" form:"menuType"`
	// 菜单状态 0显示 1隐藏
	Visible int `json:"visible" form:"visible"`
	// 权限标识
	Perms string `json:"perms" form:"perms"`
	// 菜单图标
	MenuIcon string `json:"menuIcon" form:"menuIcon"`
	// 备注
	Remark string `json:"remark" form:"remark"`
	// 是否缓存 0缓存 1不缓存
	IsCache int `json:"isCache" form:"isCache"`
	// 是否外链 0是 1否
	IsFrame int `json:"isFrame" form:"isFrame"`
	// 组件路径
	Component string `json:"component" form:"component"`
	// 路由地址
	Router string `json:"router" form:"router"`
	// 路由参数
	Query string `json:"query" form:"query"`
	// 菜单状态 0正常 1停用
	Status int `json:"status" form:"status"`
}

type SysMenuTreeNode struct {
	*SysMenu
	Children []*SysMenuTreeNode `json:"children"`
}

func BuildMenuTree(menus []*SysMenu) []*SysMenuTreeNode {
	var tree []*SysMenuTreeNode
	deptMap := make(map[int64]*SysMenuTreeNode)

	// 第一步：构建映射，便于快速查找
	for _, menu := range menus {
		node := &SysMenuTreeNode{
			SysMenu:  menu,
			Children: nil,
		}
		deptMap[menu.Id] = node
		if menu.ParentId == 0 { // 假设0是根部门的ParentId
			tree = append(tree, node)
		}
	}

	// 第二步：构建层级关系
	for _, node := range deptMap {
		if node.ParentId != 0 {
			parent, exists := deptMap[node.ParentId]
			if exists {
				parent.Children = append(parent.Children, node)
			}
		}
	}
	if tree == nil {
		tree = []*SysMenuTreeNode{}
	}
	return tree
}
