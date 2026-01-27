package model

type ArticleContent struct {
	BaseModel
	ArticleID int64  `gorm:"column:article_id;type:bigint unsigned;not null;index:idx_article"`
	Content   string `gorm:"type:text;default null;comment:内容"`
}
