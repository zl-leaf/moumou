package database

import (
	"time"

	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	DBName   string `yaml:"db_name" json:"db_name"`
	Addr     string `yaml:"addr" json:"addr"`
}

func NewMysqlGorm(cnf *DbConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSNConfig:                 transformConfig(cnf),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	return db, err
}

func transformConfig(cnf *DbConfig) *mysqldriver.Config {
	return &mysqldriver.Config{
		User:      cnf.User,
		Passwd:    cnf.Password,
		DBName:    cnf.DBName,
		Addr:      cnf.Addr,
		ParseTime: true,
		Net:       "tcp",
		Loc:       time.Local,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
		AllowNativePasswords: true,
	}
}
