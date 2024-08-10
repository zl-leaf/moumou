package dao

import (
	"database/sql"

	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB

	UserDao           *userDao
	RouterDao         *routerDao
	RoleDao           *roleDao
	PermissionDao     *permissionDao
	RolePermissionDao *rolePermissionDao
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db:                db,
		UserDao:           newUserDao(db),
		RouterDao:         newRouterDao(db),
		RoleDao:           newRoleDao(db),
		PermissionDao:     newPermissionDao(db),
		RolePermissionDao: newRolePermissionDao(db),
	}
}

func (d *Dao) Transaction(fc func(tx *Dao) error, opts ...*sql.TxOptions) (err error) {
	return d.db.Transaction(func(tx *gorm.DB) error {
		newDao := NewDao(tx)
		return fc(newDao)
	}, opts...)
}
