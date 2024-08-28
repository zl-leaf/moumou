package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/moumou/server/biz/conf"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/pkgs/database"
	"gorm.io/gorm"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
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
	err = db.AutoMigrate(
		&model.User{}, &model.Role{}, &model.Permission{}, &model.UserRelRole{}, &model.RolePermission{},
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	initPermission(ctx, db)
}

func initPermission(ctx context.Context, db *gorm.DB) {
	_ = NewPipeline(db).SetModel(&model.Permission{Path: "/manage/user"}).Before(func(modelObj interface{}) {
		m := modelObj.(*model.Permission)
		m.Name = "账号管理"
	}).Next(func(modelObj interface{}) []*InitTblPipeline {
		m := modelObj.(*model.Permission)
		pipelines := make([]*InitTblPipeline, 0, 10)
		pipelines = append(pipelines, NewPipeline(db).SetModel(&model.Permission{
			Path: "/manage/user/list",
			Pid:  m.ID,
		}).Before(func(modelObj interface{}) {
			m := modelObj.(*model.Permission)
			m.Name = "列表"
		}))
		return pipelines
	}).Execute(ctx)
}
