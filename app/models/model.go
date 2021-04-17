package models

import (
	"goblog/pkg/types"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`

	CreatedAt time.Time `gorm:"column:created_at;index" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"  json:"updated_at"`
}


// GetStringID 获取 ID 的字符串格式
func (a BaseModel)GetStringID()string  {

	return types.Uint64ToString(a.ID)
}
