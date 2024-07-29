package dao

import (
	"database/sql"
	"gorm.io/gorm"
)

type Dao struct {
	db        *gorm.DB
	RouterDao *routerDao
	UserDao   *userDao
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db:        db,
		RouterDao: newRouterDao(db),
		UserDao:   newUserDao(db),
	}
}

func (d *Dao) Transaction(fc func(tx *Dao) error, opts ...*sql.TxOptions) (err error) {
	return d.db.Transaction(func(tx *gorm.DB) error {
		newDao := NewDao(tx)
		return fc(newDao)
	})
}
