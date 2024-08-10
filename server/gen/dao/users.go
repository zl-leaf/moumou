package dao

import (
	"context"

	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userDao struct {
	*gorm.DB
}

func newUserDao(db *gorm.DB) *userDao {
	return &userDao{db}
}

func (d *userDao) WithContext(ctx context.Context) *userDao {
	return &userDao{d.DB.WithContext(ctx)}
}

func (d *userDao) WhereIDEq(ID int64) *userDao {
	d.DB = d.DB.Where(clause.Eq{Column: "id", Value: ID})
	return d
}
func (d *userDao) WhereIDNeq(ID int64) *userDao {
	d.DB = d.DB.Where(clause.Neq{Column: "id", Value: ID})
	return d
}
func (d *userDao) WhereIDIn(ID []int64) *userDao {
	d.DB = d.DB.Where(clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(ID)})
	return d
}
func (d *userDao) WhereIDNotIn(ID []int64) *userDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(ID)})))
	return d
}
func (d *userDao) WhereIDLt(ID int64) *userDao {
	d.DB = d.DB.Where(clause.Lt{Column: "id", Value: ID})
	return d
}
func (d *userDao) WhereIDLte(ID int64) *userDao {
	d.DB = d.DB.Where(clause.Lte{Column: "id", Value: ID})
	return d
}
func (d *userDao) WhereIDGt(ID int64) *userDao {
	d.DB = d.DB.Where(clause.Gt{Column: "id", Value: ID})
	return d
}
func (d *userDao) WhereIDGte(ID int64) *userDao {
	d.DB = d.DB.Where(clause.Gte{Column: "id", Value: ID})
	return d
}
func (d *userDao) WhereIDBetween(left int64, right int64) *userDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *userDao) WhereCreatedAtEq(createdAt int64) *userDao {
	d.DB = d.DB.Where(clause.Eq{Column: "created_at", Value: createdAt})
	return d
}
func (d *userDao) WhereCreatedAtNeq(createdAt int64) *userDao {
	d.DB = d.DB.Where(clause.Neq{Column: "created_at", Value: createdAt})
	return d
}
func (d *userDao) WhereCreatedAtIn(createdAt []int64) *userDao {
	d.DB = d.DB.Where(clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})
	return d
}
func (d *userDao) WhereCreatedAtNotIn(createdAt []int64) *userDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})))
	return d
}
func (d *userDao) WhereCreatedAtLt(createdAt int64) *userDao {
	d.DB = d.DB.Where(clause.Lt{Column: "created_at", Value: createdAt})
	return d
}
func (d *userDao) WhereCreatedAtLte(createdAt int64) *userDao {
	d.DB = d.DB.Where(clause.Lte{Column: "created_at", Value: createdAt})
	return d
}
func (d *userDao) WhereCreatedAtGt(createdAt int64) *userDao {
	d.DB = d.DB.Where(clause.Gt{Column: "created_at", Value: createdAt})
	return d
}
func (d *userDao) WhereCreatedAtGte(createdAt int64) *userDao {
	d.DB = d.DB.Where(clause.Gte{Column: "created_at", Value: createdAt})
	return d
}
func (d *userDao) WhereCreatedAtBetween(left int64, right int64) *userDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "created_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *userDao) WhereUpdatedAtEq(updatedAt int64) *userDao {
	d.DB = d.DB.Where(clause.Eq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userDao) WhereUpdatedAtNeq(updatedAt int64) *userDao {
	d.DB = d.DB.Where(clause.Neq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userDao) WhereUpdatedAtIn(updatedAt []int64) *userDao {
	d.DB = d.DB.Where(clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})
	return d
}
func (d *userDao) WhereUpdatedAtNotIn(updatedAt []int64) *userDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})))
	return d
}
func (d *userDao) WhereUpdatedAtLt(updatedAt int64) *userDao {
	d.DB = d.DB.Where(clause.Lt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userDao) WhereUpdatedAtLte(updatedAt int64) *userDao {
	d.DB = d.DB.Where(clause.Lte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userDao) WhereUpdatedAtGt(updatedAt int64) *userDao {
	d.DB = d.DB.Where(clause.Gt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userDao) WhereUpdatedAtGte(updatedAt int64) *userDao {
	d.DB = d.DB.Where(clause.Gte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userDao) WhereUpdatedAtBetween(left int64, right int64) *userDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "updated_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *userDao) WhereUsernameEq(username string) *userDao {
	d.DB = d.DB.Where(clause.Eq{Column: "username", Value: username})
	return d
}
func (d *userDao) WhereUsernameNeq(username string) *userDao {
	d.DB = d.DB.Where(clause.Neq{Column: "username", Value: username})
	return d
}
func (d *userDao) WhereUsernameLike(username string) *userDao {
	d.DB = d.DB.Where(clause.Like{Column: "username", Value: "%" + username + "%"})
	return d
}
func (d *userDao) WhereUsernamePrefixLike(username string) *userDao {
	d.DB = d.DB.Where(clause.Like{Column: "username", Value: username + "%"})
	return d
}
func (d *userDao) WhereUsernameNotLike(username string) *userDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "username", Value: "%" + username + "%"}))
	return d
}
func (d *userDao) WhereUsernameIn(username []string) *userDao {
	d.DB = d.DB.Where(clause.IN{Column: "username", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(username)})
	return d
}
func (d *userDao) WhereUsernameNotIn(username []string) *userDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "username", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(username)})))
	return d
}
func (d *userDao) WherePasswordEq(password string) *userDao {
	d.DB = d.DB.Where(clause.Eq{Column: "password", Value: password})
	return d
}
func (d *userDao) WherePasswordNeq(password string) *userDao {
	d.DB = d.DB.Where(clause.Neq{Column: "password", Value: password})
	return d
}
func (d *userDao) WherePasswordLike(password string) *userDao {
	d.DB = d.DB.Where(clause.Like{Column: "password", Value: "%" + password + "%"})
	return d
}
func (d *userDao) WherePasswordPrefixLike(password string) *userDao {
	d.DB = d.DB.Where(clause.Like{Column: "password", Value: password + "%"})
	return d
}
func (d *userDao) WherePasswordNotLike(password string) *userDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "password", Value: "%" + password + "%"}))
	return d
}
func (d *userDao) WherePasswordIn(password []string) *userDao {
	d.DB = d.DB.Where(clause.IN{Column: "password", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(password)})
	return d
}
func (d *userDao) WherePasswordNotIn(password []string) *userDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "password", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(password)})))
	return d
}

func (d *userDao) Limit(limit int) *userDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *userDao) Offset(offset int) *userDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *userDao) Page(page, pageSize int) *userDao {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	d.DB = d.DB.Offset((page - 1) * pageSize)
	d.DB = d.DB.Limit(pageSize)
	return d
}

func (d *userDao) GetByID(id int64) (*model.User, error) {
	return d.First(id)
}

func (d *userDao) Find() ([]*model.User, int64, error) {
	var list []*model.User
	var total int64
	var query = d.Model(model.User{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *userDao) First(conds ...interface{}) (*model.User, error) {
	var record = &model.User{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (d *userDao) Create(record *model.User) error {
	return d.DB.Create(record).Error
}

func (d *userDao) Save(record *model.User) error {
	return d.DB.Save(record).Error
}

func (d *userDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.User{}, conds...).Error
}
