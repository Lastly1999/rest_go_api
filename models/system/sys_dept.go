package models

import "resetgoapi.com/rest_go_api/models"

type SysDept struct {
	models.BaseModel
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
	Status int `json:"status" form:"status"`
	// 备注
	Remark string `json:"remark" form:"remark"`
}

type SysDeptTreeNode struct {
	*SysDept
	Children []*SysDeptTreeNode `json:"children"`
}

func BuildDeptTree(depts []*SysDept) []*SysDeptTreeNode {
	var tree []*SysDeptTreeNode
	deptMap := make(map[int64]*SysDeptTreeNode)

	// 第一步：构建映射，便于快速查找
	for _, dept := range depts {
		node := &SysDeptTreeNode{
			SysDept:  dept,
			Children: nil,
		}
		deptMap[dept.Id] = node
		if dept.ParentId == 0 { // 假设0是根部门的ParentId
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

	return tree
}
