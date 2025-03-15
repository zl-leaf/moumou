package model

import "gorm.io/plugin/soft_delete"

type BaseModel struct {
	Id        int64                 `gorm:"column:id;type:bigint unsigned;not null;primarykey" json:"id"`
	CreatedAt int64                 `gorm:"autoCreateTime:milli;type:bigint unsigned;not null;default:0" json:"created_at"`
	UpdatedAt int64                 `gorm:"autoUpdateTime:milli;type:bigint unsigned;not null;default:0" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli;type:bigint unsigned;not null;default:0" json:"deleted_at"`
}
