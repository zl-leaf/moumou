package dao

import (
	"context"

	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRelRoleDao struct {
	*gorm.DB
}

func newUserRelRoleDao(db *gorm.DB) *userRelRoleDao {
	return &userRelRoleDao{db}
}

func (d *userRelRoleDao) WithContext(ctx context.Context) *userRelRoleDao {
	return &userRelRoleDao{d.DB.WithContext(ctx)}
}

func (d *userRelRoleDao) WhereIdEq(id int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "id", Value: id})
	return d
}
func (d *userRelRoleDao) WhereIdNeq(id int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "id", Value: id})
	return d
}
func (d *userRelRoleDao) WhereIdIn(id []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})
	return d
}
func (d *userRelRoleDao) WhereIdNotIn(id []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})))
	return d
}
func (d *userRelRoleDao) WhereIdLt(id int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "id", Value: id})
	return d
}
func (d *userRelRoleDao) WhereIdLte(id int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "id", Value: id})
	return d
}
func (d *userRelRoleDao) WhereIdGt(id int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "id", Value: id})
	return d
}
func (d *userRelRoleDao) WhereIdGte(id int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "id", Value: id})
	return d
}
func (d *userRelRoleDao) WhereIdBetween(left int64, right int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *userRelRoleDao) WhereCreatedAtEq(createdAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "created_at", Value: createdAt})
	return d
}
func (d *userRelRoleDao) WhereCreatedAtNeq(createdAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "created_at", Value: createdAt})
	return d
}
func (d *userRelRoleDao) WhereCreatedAtIn(createdAt []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})
	return d
}
func (d *userRelRoleDao) WhereCreatedAtNotIn(createdAt []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})))
	return d
}
func (d *userRelRoleDao) WhereCreatedAtLt(createdAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "created_at", Value: createdAt})
	return d
}
func (d *userRelRoleDao) WhereCreatedAtLte(createdAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "created_at", Value: createdAt})
	return d
}
func (d *userRelRoleDao) WhereCreatedAtGt(createdAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "created_at", Value: createdAt})
	return d
}
func (d *userRelRoleDao) WhereCreatedAtGte(createdAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "created_at", Value: createdAt})
	return d
}
func (d *userRelRoleDao) WhereCreatedAtBetween(left int64, right int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "created_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtEq(updatedAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtNeq(updatedAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtIn(updatedAt []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtNotIn(updatedAt []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})))
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtLt(updatedAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtLte(updatedAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtGt(updatedAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtGte(updatedAt int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *userRelRoleDao) WhereUpdatedAtBetween(left int64, right int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "updated_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *userRelRoleDao) WhereUserIDEq(userID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "user_id", Value: userID})
	return d
}
func (d *userRelRoleDao) WhereUserIDNeq(userID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "user_id", Value: userID})
	return d
}
func (d *userRelRoleDao) WhereUserIDIn(userID []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.IN{Column: "user_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(userID)})
	return d
}
func (d *userRelRoleDao) WhereUserIDNotIn(userID []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "user_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(userID)})))
	return d
}
func (d *userRelRoleDao) WhereUserIDLt(userID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "user_id", Value: userID})
	return d
}
func (d *userRelRoleDao) WhereUserIDLte(userID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "user_id", Value: userID})
	return d
}
func (d *userRelRoleDao) WhereUserIDGt(userID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "user_id", Value: userID})
	return d
}
func (d *userRelRoleDao) WhereUserIDGte(userID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "user_id", Value: userID})
	return d
}
func (d *userRelRoleDao) WhereUserIDBetween(left int64, right int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "user_id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *userRelRoleDao) WhereRoleIDEq(roleID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "role_id", Value: roleID})
	return d
}
func (d *userRelRoleDao) WhereRoleIDNeq(roleID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "role_id", Value: roleID})
	return d
}
func (d *userRelRoleDao) WhereRoleIDIn(roleID []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.IN{Column: "role_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(roleID)})
	return d
}
func (d *userRelRoleDao) WhereRoleIDNotIn(roleID []int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "role_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(roleID)})))
	return d
}
func (d *userRelRoleDao) WhereRoleIDLt(roleID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "role_id", Value: roleID})
	return d
}
func (d *userRelRoleDao) WhereRoleIDLte(roleID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "role_id", Value: roleID})
	return d
}
func (d *userRelRoleDao) WhereRoleIDGt(roleID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "role_id", Value: roleID})
	return d
}
func (d *userRelRoleDao) WhereRoleIDGte(roleID int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "role_id", Value: roleID})
	return d
}
func (d *userRelRoleDao) WhereRoleIDBetween(left int64, right int64) *userRelRoleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "role_id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}

func (d *userRelRoleDao) Limit(limit int) *userRelRoleDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *userRelRoleDao) Offset(offset int) *userRelRoleDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *userRelRoleDao) Page(page, pageSize int) *userRelRoleDao {
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

func (d *userRelRoleDao) GetByID(id int64) (*model.UserRelRole, error) {
	return d.First(id)
}

func (d *userRelRoleDao) Find() ([]*model.UserRelRole, int64, error) {
	var list []*model.UserRelRole
	var total int64
	var query = d.Model(model.UserRelRole{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *userRelRoleDao) Count() (int64, error) {
	var total int64
	var query = d.Model(model.UserRelRole{})

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *userRelRoleDao) First(conds ...interface{}) (*model.UserRelRole, error) {
	var record = &model.UserRelRole{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (d *userRelRoleDao) Create(record *model.UserRelRole) error {
	return d.DB.Create(record).Error
}

func (d *userRelRoleDao) Save(record *model.UserRelRole) error {
	return d.DB.Save(record).Error
}

func (d *userRelRoleDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.UserRelRole{}, conds...).Error
}
