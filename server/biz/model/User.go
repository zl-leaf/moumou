package model

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(50);not null;comment:用户名;index:idx_username" json:"username"`
	Password string `gorm:"type:varchar(200);not null;comment:密码" json:"-"`
}
