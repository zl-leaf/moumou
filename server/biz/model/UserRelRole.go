package model

type UserRelRole struct {
	BaseModel
	UserID int64 `gorm:"column:user_id;type:bigint unsigned;not null;index:idx_user_role"`
	RoleID int64 `gorm:"column:role_id;type:bigint unsigned;not null;index:idx_user_role"`
}

func (UserRelRole) TableName() string {
	return "user_rel_role"
}