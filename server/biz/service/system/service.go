package system

import (
	"context"
	"errors"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/permission"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
	"gorm.io/gorm"
)

// 允许无条件的 Delete，用于清空表
var allowGlobalUpdate = &gorm.Session{AllowGlobalUpdate: true}

type Service struct {
	db                *gorm.DB
	dao               *dao.Dao
	userService       *user.Service
	permissionService *permission.Service
}

func NewService(
	db *gorm.DB,
	dao *dao.Dao,
	userService *user.Service,
	permissionService *permission.Service,
) *Service {
	return &Service{
		db:                db,
		dao:               dao,
		userService:       userService,
		permissionService: permissionService,
	}
}

// Initialize 初始化数据表与初始化数据；adminUsername/adminPassword 用于创建管理员账号（原 root 账号）
func (s *Service) Initialize(ctx context.Context, adminUsername, adminPassword string) error {
	if err := s.db.WithContext(ctx).AutoMigrate(
		&model.User{}, &model.Role{}, &model.Permission{}, &model.UserRelRole{}, &model.RolePermission{},
		&model.Article{}, &model.ArticleContent{},
	); err != nil {
		return err
	}

	if err := s.deleteAllTables(ctx); err != nil {
		return err
	}

	s.initUser(ctx, adminUsername, adminPassword)
	s.initRole(ctx)
	s.initPermission(ctx)
	s.initRolePermission(ctx, adminUsername)
	return nil
}

// deleteAllTables 删除所有数据表的数据（依赖表先删，避免外键约束）
func (s *Service) deleteAllTables(ctx context.Context) error {
	db := s.db.WithContext(ctx).Session(allowGlobalUpdate)
	tables := []interface{}{
		&model.ArticleContent{},
		&model.Article{},
		&model.UserRelRole{},
		&model.RolePermission{},
		&model.User{},
		&model.Role{},
		&model.Permission{},
	}
	for _, m := range tables {
		if err := db.Delete(m).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) initUser(ctx context.Context, adminUsername, adminPassword string) {
	_, err := s.dao.UserDao(ctx).WhereUsernameEq(adminUsername).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = s.userService.CreateUser(ctx, &model.User{
			Username: adminUsername,
			Password: adminPassword,
		})
	}

	_, err = s.dao.UserDao(ctx).WhereUsernameEq("demo").First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = s.userService.CreateUser(ctx, &model.User{
			Username: "demo", // demo用户的密码首次需要通过adminUsername账号进行重置
		})
	}
}

func (s *Service) initRole(ctx context.Context) {
	_, err := s.dao.RoleDao(ctx).WhereNameEq("管理员").First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = s.dao.RoleDao(ctx).Create(&model.Role{Name: "管理员"})
	}

	_, err = s.dao.RoleDao(ctx).WhereNameEq("普通用户").First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_ = s.dao.RoleDao(ctx).Create(&model.Role{Name: "普通用户"})
	}
}

func (s *Service) initPermission(ctx context.Context) {
	db := s.db.WithContext(ctx)

	_ = NewPipeline(db).SetModel(&model.Permission{Code: "ManageUser"}).Before(func(modelObj interface{}) {
		m := modelObj.(*model.Permission)
		m.Name = "账号管理"
	}).Next(func(modelObj interface{}) []*InitTblPipeline {
		m := modelObj.(*model.Permission)
		return []*InitTblPipeline{
			NewPipeline(db).SetModel(&model.Permission{Code: "ManageUserRead", Pid: m.Id}).Before(func(o interface{}) {
				x := o.(*model.Permission)
				x.Name = "可读"
				x.Sort = 1
			}),
			NewPipeline(db).SetModel(&model.Permission{Code: "ManageUserWrite", Pid: m.Id}).Before(func(o interface{}) {
				x := o.(*model.Permission)
				x.Name = "可写"
				x.Sort = 2
			}),
		}
	}).Execute(ctx)

	_ = NewPipeline(db).SetModel(&model.Permission{Code: "ManageRole"}).Before(func(modelObj interface{}) {
		m := modelObj.(*model.Permission)
		m.Name = "角色管理"
	}).Next(func(modelObj interface{}) []*InitTblPipeline {
		m := modelObj.(*model.Permission)
		return []*InitTblPipeline{
			NewPipeline(db).SetModel(&model.Permission{Code: "ManageRoleRead", Pid: m.Id}).Before(func(o interface{}) {
				x := o.(*model.Permission)
				x.Name = "可读"
				x.Sort = 1
			}),
			NewPipeline(db).SetModel(&model.Permission{Code: "ManageRoleWrite", Pid: m.Id}).Before(func(o interface{}) {
				x := o.(*model.Permission)
				x.Name = "可写"
				x.Sort = 2
			}),
		}
	}).Execute(ctx)

	_ = NewPipeline(db).SetModel(&model.Permission{Code: "ManagePermission"}).Before(func(modelObj interface{}) {
		m := modelObj.(*model.Permission)
		m.Name = "权限管理"
	}).Next(func(modelObj interface{}) []*InitTblPipeline {
		m := modelObj.(*model.Permission)
		return []*InitTblPipeline{
			NewPipeline(db).SetModel(&model.Permission{Code: "ManagePermissionRead", Pid: m.Id}).Before(func(o interface{}) {
				x := o.(*model.Permission)
				x.Name = "可读"
				x.Sort = 1
			}),
			NewPipeline(db).SetModel(&model.Permission{Code: "ManagePermissionWrite", Pid: m.Id}).Before(func(o interface{}) {
				x := o.(*model.Permission)
				x.Name = "可写"
				x.Sort = 2
			}),
		}
	}).Execute(ctx)
}

func (s *Service) initRolePermission(ctx context.Context, adminUsername string) {
	allPermissions, _, _ := s.dao.PermissionDao(ctx).Find()

	adminUser, err := s.dao.UserDao(ctx).WhereUsernameEq(adminUsername).First()
	if err != nil {
		panic(err)
	}
	adminRole, err := s.dao.RoleDao(ctx).WhereNameEq("管理员").First()
	if err != nil {
		panic(err)
	}
	adminRolePermissionTotal, _ := s.dao.RolePermissionDao(ctx).WhereRoleIDEq(adminRole.Id).Count()
	if adminRolePermissionTotal == 0 {
		allPermissionIds := make([]int64, len(allPermissions))
		for i, p := range allPermissions {
			allPermissionIds[i] = p.Id
		}
		_ = s.permissionService.UpdateRolePermission(ctx, adminRole.Id, allPermissionIds)
		_ = s.permissionService.BindUsers(ctx, adminRole.Id, []int64{adminUser.Id})
	}

	demoUser, err := s.dao.UserDao(ctx).WhereUsernameEq("demo").First()
	if err != nil {
		panic(err)
	}
	normalRole, err := s.dao.RoleDao(ctx).WhereNameEq("普通用户").First()
	if err != nil {
		panic(err)
	}
	normalRolePermissionTotal, _ := s.dao.RolePermissionDao(ctx).WhereRoleIDEq(normalRole.Id).Count()
	if normalRolePermissionTotal == 0 {
		var normalPermissionIds []int64
		for _, p := range allPermissions {
			if p.Code == "ManageUserRead" || p.Code == "ManageRoleRead" || p.Code == "ManagePermissionRead" {
				normalPermissionIds = append(normalPermissionIds, p.Id)
			}
		}
		_ = s.permissionService.UpdateRolePermission(ctx, normalRole.Id, normalPermissionIds)
		_ = s.permissionService.BindUsers(ctx, normalRole.Id, []int64{demoUser.Id})
	}
}
