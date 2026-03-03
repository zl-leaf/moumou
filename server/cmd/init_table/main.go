package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/moumou/server/biz/conf"
	"github.com/moumou/server/biz/service/permission"
	"github.com/moumou/server/biz/service/system"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
	"github.com/moumou/server/pkgs/database"
	"github.com/moumou/server/pkgs/env"
)

var (
	flagconf         string
	flagAdminUser    string
	flagAdminPasswd  string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./config", "config path, eg: -conf ./config")
	flag.StringVar(&flagAdminUser, "admin_username", "", "admin username for root user")
	flag.StringVar(&flagAdminPasswd, "admin_password", "", "admin password for root user")
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("%s/config_%s.yaml", flagconf, env.GetEnv())),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var cnf conf.Data
	if err := c.Scan(&cnf); err != nil {
		panic(err)
	}

	db, err := database.NewMysqlGorm(&cnf.DBConfig)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	dbDao := dao.NewDao(db)
	userSvc := user.NewUserService(&cnf, dbDao)
	permSvc := permission.NewService(dbDao)
	systemSvc := system.NewService(db, dbDao, userSvc, permSvc)

	adminUsername := flagAdminUser
	adminPassword := flagAdminPasswd
	if adminUsername == "" {
		adminUsername = cnf.SystemConfig.Username
	}
	if adminPassword == "" {
		adminPassword = cnf.SystemConfig.Password
	}
	if adminUsername == "" || adminPassword == "" {
		fmt.Fprintln(os.Stderr, "error: admin_username and admin_password are required (use -admin_username and -admin_password, or set in config system.username/system.password)")
		os.Exit(1)
	}
	if err := systemSvc.Initialize(ctx, adminUsername, adminPassword); err != nil {
		panic(err)
	}
}
