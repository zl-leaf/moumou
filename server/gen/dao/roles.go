package dao

import (
	"context"

	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type roleDao struct {
	*gorm.DB
}

func newRoleDao(db *gorm.DB) *roleDao {
	return &roleDao{db}
}

func (d *roleDao) WithContext(ctx context.Context) *roleDao {
	return &roleDao{d.DB.WithContext(ctx)}
}

func (d *roleDao) WhereIdEq(id int64) *roleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "id", Value: id})
	return d
}
func (d *roleDao) WhereIdNeq(id int64) *roleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "id", Value: id})
	return d
}
func (d *roleDao) WhereIdIn(id []int64) *roleDao {
	d.DB = d.DB.Where(clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})
	return d
}
func (d *roleDao) WhereIdNotIn(id []int64) *roleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})))
	return d
}
func (d *roleDao) WhereIdLt(id int64) *roleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "id", Value: id})
	return d
}
func (d *roleDao) WhereIdLte(id int64) *roleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "id", Value: id})
	return d
}
func (d *roleDao) WhereIdGt(id int64) *roleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "id", Value: id})
	return d
}
func (d *roleDao) WhereIdGte(id int64) *roleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "id", Value: id})
	return d
}
func (d *roleDao) WhereIdBetween(left int64, right int64) *roleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *roleDao) WhereCreatedAtEq(createdAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "created_at", Value: createdAt})
	return d
}
func (d *roleDao) WhereCreatedAtNeq(createdAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "created_at", Value: createdAt})
	return d
}
func (d *roleDao) WhereCreatedAtIn(createdAt []int64) *roleDao {
	d.DB = d.DB.Where(clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})
	return d
}
func (d *roleDao) WhereCreatedAtNotIn(createdAt []int64) *roleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})))
	return d
}
func (d *roleDao) WhereCreatedAtLt(createdAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "created_at", Value: createdAt})
	return d
}
func (d *roleDao) WhereCreatedAtLte(createdAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "created_at", Value: createdAt})
	return d
}
func (d *roleDao) WhereCreatedAtGt(createdAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "created_at", Value: createdAt})
	return d
}
func (d *roleDao) WhereCreatedAtGte(createdAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "created_at", Value: createdAt})
	return d
}
func (d *roleDao) WhereCreatedAtBetween(left int64, right int64) *roleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "created_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *roleDao) WhereUpdatedAtEq(updatedAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *roleDao) WhereUpdatedAtNeq(updatedAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *roleDao) WhereUpdatedAtIn(updatedAt []int64) *roleDao {
	d.DB = d.DB.Where(clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})
	return d
}
func (d *roleDao) WhereUpdatedAtNotIn(updatedAt []int64) *roleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})))
	return d
}
func (d *roleDao) WhereUpdatedAtLt(updatedAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *roleDao) WhereUpdatedAtLte(updatedAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *roleDao) WhereUpdatedAtGt(updatedAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *roleDao) WhereUpdatedAtGte(updatedAt int64) *roleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *roleDao) WhereUpdatedAtBetween(left int64, right int64) *roleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "updated_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *roleDao) WhereNameEq(name string) *roleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "name", Value: name})
	return d
}
func (d *roleDao) WhereNameNeq(name string) *roleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "name", Value: name})
	return d
}
func (d *roleDao) WhereNameLike(name string) *roleDao {
	d.DB = d.DB.Where(clause.Like{Column: "name", Value: "%" + name + "%"})
	return d
}
func (d *roleDao) WhereNamePrefixLike(name string) *roleDao {
	d.DB = d.DB.Where(clause.Like{Column: "name", Value: name + "%"})
	return d
}
func (d *roleDao) WhereNameNotLike(name string) *roleDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "name", Value: "%" + name + "%"}))
	return d
}
func (d *roleDao) WhereNameIn(name []string) *roleDao {
	d.DB = d.DB.Where(clause.IN{Column: "name", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(name)})
	return d
}
func (d *roleDao) WhereNameNotIn(name []string) *roleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "name", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(name)})))
	return d
}

func (d *roleDao) Limit(limit int) *roleDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *roleDao) Offset(offset int) *roleDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *roleDao) Page(page, pageSize int) *roleDao {
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

func (d *roleDao) GetByID(id int64) (*model.Role, error) {
	return d.First(id)
}

func (d *roleDao) Find() ([]*model.Role, int64, error) {
	var list []*model.Role
	var total int64
	var query = d.Model(model.Role{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *roleDao) Count() (int64, error) {
	var total int64
	var query = d.Model(model.Role{})

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *roleDao) First(conds ...interface{}) (*model.Role, error) {
	var record = &model.Role{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (d *roleDao) Create(record *model.Role) error {
	return d.DB.Create(record).Error
}

func (d *roleDao) Save(record *model.Role) error {
	return d.DB.Save(record).Error
}

func (d *roleDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.Role{}, conds...).Error
}
