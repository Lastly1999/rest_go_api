package models

import "time"

type BaseModel struct {
	// 主键
	Id int64 `json:"id" form:"id" gorm:"primaryKey"`
	// 创建时间
	CreatedAt time.Time `json:"createdAt" form:"createdAt" gorm:"autoCreateTime"`
	// 更新时间
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt" gorm:"autoUpdateTime"`
}
