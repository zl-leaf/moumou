package model

import "gorm.io/plugin/soft_delete"

type BaseModel struct {
	ID        int64                 `gorm:"primarykey" json:"id"`
	CreatedAt int64                 `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64                 `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli" json:"deleted_at"`
}
