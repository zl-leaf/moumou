package model

type RolePermission struct {
	BaseModel
	RoleID       int64 `gorm:"column:role_id;type:bigint unsigned;not null;index:idx_role_permission"`
	PermissionID int64 `gorm:"column:permission_id;type:bigint unsigned;not null;index:idx_role_permission"`
}

func (RolePermission) TableName() string {
	return "role_rel_permission"
}
