package model

type Article struct {
	BaseModel
	Title string `gorm:"type:varchar(512);not null;comment:标题"`
}
