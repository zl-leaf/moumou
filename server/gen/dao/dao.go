package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db: db,
	}
}

func (d *Dao) UserDao(ctx context.Context) *userDao {
	return newUserDao(d.db).WithContext(ctx)
}
func (d *Dao) RoleDao(ctx context.Context) *roleDao {
	return newRoleDao(d.db).WithContext(ctx)
}
func (d *Dao) PermissionDao(ctx context.Context) *permissionDao {
	return newPermissionDao(d.db).WithContext(ctx)
}
func (d *Dao) RolePermissionDao(ctx context.Context) *rolePermissionDao {
	return newRolePermissionDao(d.db).WithContext(ctx)
}
func (d *Dao) UserRelRoleDao(ctx context.Context) *userRelRoleDao {
	return newUserRelRoleDao(d.db).WithContext(ctx)
}

func (d *Dao) Transaction(ctx context.Context, fc func(tx *Dao) error, opts ...*sql.TxOptions) (err error) {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		newDao := NewDao(tx)
		return fc(newDao)
	}, opts...)
}
