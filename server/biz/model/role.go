package model

type Role struct {
	BaseModel
	Name string `gorm:"type:varchar(256);not null;comment:名称"`
}
