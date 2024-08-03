package model

type Permission struct {
	BaseModel
	Name string `gorm:"type:varchar(50);not null;comment:权限名称;"`
	Path string `gorm:"type:varchar(256);not null;comment:路径;"`
	Pid  int64  `gorm:"type:bigint unsigned;not null;default:0;comment:父节点id;index:idx_pid"`
	Sort int    `gorm:"type:int;not null;default:0;comment:排序"`
}
