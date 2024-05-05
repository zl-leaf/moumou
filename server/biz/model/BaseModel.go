package model

import "gorm.io/plugin/soft_delete"

type BaseModel struct {
	ID        uint                  `gorm:"primarykey"`
	CreatedAt int64                 `gorm:"autoCreateTime:milli"`
	UpdatedAt int64                 `gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`
}
