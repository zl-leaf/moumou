package model

type Router struct {
	BaseModel
	Name      string `gorm:"type:varchar(50);not null;comment:标识名称;"`
	Path      string `gorm:"type:varchar(50);not null;comment:路径;"`
	Title     string `gorm:"type:varchar(10);not null;comment:菜单名;"`
	IsMenu    bool   `gorm:"type:tinyint(1);not null;default:0;comment:是否菜单;"`
	IsSystem  bool   `gorm:"type:tinyint(1);not null;default:0;comment:是否菜单;"`
	Pid       uint   `gorm:"type:bigint unsigned;not null;default:0;comment:父节点id;index:idx_pid"`
	Sort      int    `gorm:"type:int;not null;default:0;comment:排序"`
	Component string `gorm:"type:varchar(256);not null;default:'';comment:自定义页面;"`
}
