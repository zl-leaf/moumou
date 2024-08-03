package main

import (
	"context"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/pkgs/config"
	"github.com/moumou/server/pkgs/database"
	"gorm.io/gorm"
)

func main() {
	err := config.Init("config")
	if err != nil {
		panic(err)
	}
	db, err := database.NewMysqlGorm()
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
