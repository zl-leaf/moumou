package database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/moumou/server/pkgs/config"
	"time"
)

type dbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	Addr     string `yaml:"addr"`
}

func newConfig() (*mysql.Config, error) {
	var cnf dbConfig
	err := config.GetConfig("database", &cnf)
	if err != nil {
		return nil, err
	}

	return &mysql.Config{
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
	}, nil
}
