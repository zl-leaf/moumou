package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/moumou/server/biz/conf"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/role"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
	"github.com/moumou/server/pkgs/database"
	"github.com/moumou/server/pkgs/env"
	"gorm.io/gorm"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./config", "config path, eg: -conf ./config")
}

func main() {
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
	err = db.AutoMigrate(
		&model.User{}, &model.Role{}, &model.Permission{}, &model.UserRelRole{}, &model.RolePermission{},
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	dbDao := dao.NewDao(db)
	initUser(ctx, user.NewUserService(&cnf, dbDao), dbDao)
	initRole(ctx, role.NewService(dbDao), dbDao)
	initPermission(ctx, db)
	initRolePermission(ctx, role.NewService(dbDao), dbDao)
}

func initUser(ctx context.Context, userService *user.Service, dao *dao.Dao) {
	// 创建root和demo账号
	_, err := dao.UserDao(ctx).WhereUsernameEq("root").First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = userService.CreateUser(ctx, &model.User{
			Username: "root",
			Password: "root",
		})
	}

	_, err = dao.UserDao(ctx).WhereUsernameEq("demo").First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = userService.CreateUser(ctx, &model.User{
			Username: "demo",
			Password: "demo",
		})
	}
}

func initPermission(ctx context.Context, db *gorm.DB) {
	_ = NewPipeline(db).SetModel(&model.Permission{Code: "ManageUser"}).Before(func(modelObj interface{}) {
		m := modelObj.(*model.Permission)
		m.Name = "账号管理"
	}).Next(func(modelObj interface{}) []*InitTblPipeline {
		m := modelObj.(*model.Permission)
		pipelines := make([]*InitTblPipeline, 0, 10)
		pipelines = append(pipelines, NewPipeline(db).SetModel(&model.Permission{
			Code: "ManageUserRead",
			Pid:  m.Id,
		}).Before(func(modelObj interface{}) {
			m := modelObj.(*model.Permission)
			m.Name = "可读"
			m.Sort = 1
		}))

		pipelines = append(pipelines, NewPipeline(db).SetModel(&model.Permission{
			Code: "ManageUserWrite",
			Pid:  m.Id,
		}).Before(func(modelObj interface{}) {
			m := modelObj.(*model.Permission)
			m.Name = "可写"
			m.Sort = 2
		}))
		return pipelines
	}).Execute(ctx)

	_ = NewPipeline(db).SetModel(&model.Permission{Code: "ManageRole"}).Before(func(modelObj interface{}) {
		m := modelObj.(*model.Permission)
		m.Name = "角色管理"
	}).Next(func(modelObj interface{}) []*InitTblPipeline {
		m := modelObj.(*model.Permission)
		pipelines := make([]*InitTblPipeline, 0, 10)
		pipelines = append(pipelines, NewPipeline(db).SetModel(&model.Permission{
			Code: "ManageRoleRead",
			Pid:  m.Id,
		}).Before(func(modelObj interface{}) {
			m := modelObj.(*model.Permission)
			m.Name = "可读"
			m.Sort = 1
		}))

		pipelines = append(pipelines, NewPipeline(db).SetModel(&model.Permission{
			Code: "ManageRoleWrite",
			Pid:  m.Id,
		}).Before(func(modelObj interface{}) {
			m := modelObj.(*model.Permission)
			m.Name = "可写"
			m.Sort = 2
		}))
		return pipelines
	}).Execute(ctx)

	_ = NewPipeline(db).SetModel(&model.Permission{Code: "ManagePermission"}).Before(func(modelObj interface{}) {
		m := modelObj.(*model.Permission)
		m.Name = "权限管理"
	}).Next(func(modelObj interface{}) []*InitTblPipeline {
		m := modelObj.(*model.Permission)
		pipelines := make([]*InitTblPipeline, 0, 10)
		pipelines = append(pipelines, NewPipeline(db).SetModel(&model.Permission{
			Code: "ManagePermissionRead",
			Pid:  m.Id,
		}).Before(func(modelObj interface{}) {
			m := modelObj.(*model.Permission)
			m.Name = "可读"
			m.Sort = 1
		}))

		pipelines = append(pipelines, NewPipeline(db).SetModel(&model.Permission{
			Code: "ManagePermissionWrite",
			Pid:  m.Id,
		}).Before(func(modelObj interface{}) {
			m := modelObj.(*model.Permission)
			m.Name = "可写"
			m.Sort = 2
		}))
		return pipelines
	}).Execute(ctx)
}

func initRole(ctx context.Context, roleService *role.Service, dao *dao.Dao) {
	// 创建管理员和普通用户角色
	_, err := dao.RoleDao(ctx).WhereNameEq("管理员").First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = dao.RoleDao(ctx).Create(&model.Role{Name: "管理员"})
	}

	_, err = dao.RoleDao(ctx).WhereNameEq("普通用户").First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = dao.RoleDao(ctx).Create(&model.Role{Name: "普通用户"})
	}
}

func initRolePermission(ctx context.Context, roleService *role.Service, dao *dao.Dao) {
	allPermissions, _, _ := dao.PermissionDao(ctx).Find()

	rootUser, err := dao.UserDao(ctx).WhereUsernameEq("root").First()
	if err != nil {
		panic(err)
	}
	adminRole, err := dao.RoleDao(ctx).WhereNameEq("管理员").First()
	if err != nil {
		panic(err)
	}
	// 如果管理员角色没有授权过权限，授权所有，并且绑定root
	adminRolePermissionTotal, _ := dao.RolePermissionDao(ctx).WhereRoleIDEq(adminRole.Id).Count()
	if adminRolePermissionTotal == 0 {
		allPermissionIds := make([]int64, len(allPermissions))
		for i, permission := range allPermissions {
			allPermissionIds[i] = permission.Id
		}
		_ = roleService.UpdateRolePermission(ctx, adminRole.Id, allPermissionIds)
		_ = roleService.BindUsers(ctx, adminRole.Id, []int64{rootUser.Id})
	}

	demoUser, err := dao.UserDao(ctx).WhereUsernameEq("demo").First()
	if err != nil {
		panic(err)
	}
	normalRole, err := dao.RoleDao(ctx).WhereNameEq("普通用户").First()
	if err != nil {
		panic(err)
	}
	// 如果普通用户角色没有授权过权限，授权所有，并且绑定demo
	normalRolePermissionTotal, _ := dao.RolePermissionDao(ctx).WhereRoleIDEq(normalRole.Id).Count()
	if normalRolePermissionTotal == 0 {
		normalPermissionIds := make([]int64, 0, len(allPermissions))
		for _, permission := range allPermissions {
			if permission.Code == "ManageUserRead" || permission.Code == "ManageRoleRead" || permission.Code == "ManagePermissionRead" {
				normalPermissionIds = append(normalPermissionIds, permission.Id)
			}
		}
		_ = roleService.UpdateRolePermission(ctx, normalRole.Id, normalPermissionIds)
		_ = roleService.BindUsers(ctx, normalRole.Id, []int64{demoUser.Id})
	}
}
