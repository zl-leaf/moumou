package dao

import (
	"context"

	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type permissionDao struct {
	*gorm.DB
}

func newPermissionDao(db *gorm.DB) *permissionDao {
	return &permissionDao{db}
}

func (d *permissionDao) WithContext(ctx context.Context) *permissionDao {
	return &permissionDao{d.DB.WithContext(ctx)}
}

func (d *permissionDao) WhereIdEq(id int64) *permissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "id", Value: id})
	return d
}
func (d *permissionDao) WhereIdNeq(id int64) *permissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "id", Value: id})
	return d
}
func (d *permissionDao) WhereIdIn(id []int64) *permissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})
	return d
}
func (d *permissionDao) WhereIdNotIn(id []int64) *permissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})))
	return d
}
func (d *permissionDao) WhereIdLt(id int64) *permissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "id", Value: id})
	return d
}
func (d *permissionDao) WhereIdLte(id int64) *permissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "id", Value: id})
	return d
}
func (d *permissionDao) WhereIdGt(id int64) *permissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "id", Value: id})
	return d
}
func (d *permissionDao) WhereIdGte(id int64) *permissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "id", Value: id})
	return d
}
func (d *permissionDao) WhereIdBetween(left int64, right int64) *permissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *permissionDao) WhereCreatedAtEq(createdAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "created_at", Value: createdAt})
	return d
}
func (d *permissionDao) WhereCreatedAtNeq(createdAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "created_at", Value: createdAt})
	return d
}
func (d *permissionDao) WhereCreatedAtIn(createdAt []int64) *permissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})
	return d
}
func (d *permissionDao) WhereCreatedAtNotIn(createdAt []int64) *permissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})))
	return d
}
func (d *permissionDao) WhereCreatedAtLt(createdAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "created_at", Value: createdAt})
	return d
}
func (d *permissionDao) WhereCreatedAtLte(createdAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "created_at", Value: createdAt})
	return d
}
func (d *permissionDao) WhereCreatedAtGt(createdAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "created_at", Value: createdAt})
	return d
}
func (d *permissionDao) WhereCreatedAtGte(createdAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "created_at", Value: createdAt})
	return d
}
func (d *permissionDao) WhereCreatedAtBetween(left int64, right int64) *permissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "created_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *permissionDao) WhereUpdatedAtEq(updatedAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *permissionDao) WhereUpdatedAtNeq(updatedAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *permissionDao) WhereUpdatedAtIn(updatedAt []int64) *permissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})
	return d
}
func (d *permissionDao) WhereUpdatedAtNotIn(updatedAt []int64) *permissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})))
	return d
}
func (d *permissionDao) WhereUpdatedAtLt(updatedAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *permissionDao) WhereUpdatedAtLte(updatedAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *permissionDao) WhereUpdatedAtGt(updatedAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *permissionDao) WhereUpdatedAtGte(updatedAt int64) *permissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *permissionDao) WhereUpdatedAtBetween(left int64, right int64) *permissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "updated_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *permissionDao) WhereNameEq(name string) *permissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "name", Value: name})
	return d
}
func (d *permissionDao) WhereNameNeq(name string) *permissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "name", Value: name})
	return d
}
func (d *permissionDao) WhereNameLike(name string) *permissionDao {
	d.DB = d.DB.Where(clause.Like{Column: "name", Value: "%" + name + "%"})
	return d
}
func (d *permissionDao) WhereNamePrefixLike(name string) *permissionDao {
	d.DB = d.DB.Where(clause.Like{Column: "name", Value: name + "%"})
	return d
}
func (d *permissionDao) WhereNameNotLike(name string) *permissionDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "name", Value: "%" + name + "%"}))
	return d
}
func (d *permissionDao) WhereNameIn(name []string) *permissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "name", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(name)})
	return d
}
func (d *permissionDao) WhereNameNotIn(name []string) *permissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "name", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(name)})))
	return d
}
func (d *permissionDao) WhereCodeEq(code string) *permissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "code", Value: code})
	return d
}
func (d *permissionDao) WhereCodeNeq(code string) *permissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "code", Value: code})
	return d
}
func (d *permissionDao) WhereCodeLike(code string) *permissionDao {
	d.DB = d.DB.Where(clause.Like{Column: "code", Value: "%" + code + "%"})
	return d
}
func (d *permissionDao) WhereCodePrefixLike(code string) *permissionDao {
	d.DB = d.DB.Where(clause.Like{Column: "code", Value: code + "%"})
	return d
}
func (d *permissionDao) WhereCodeNotLike(code string) *permissionDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "code", Value: "%" + code + "%"}))
	return d
}
func (d *permissionDao) WhereCodeIn(code []string) *permissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "code", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(code)})
	return d
}
func (d *permissionDao) WhereCodeNotIn(code []string) *permissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "code", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(code)})))
	return d
}
func (d *permissionDao) WherePidEq(pid int64) *permissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "pid", Value: pid})
	return d
}
func (d *permissionDao) WherePidNeq(pid int64) *permissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "pid", Value: pid})
	return d
}
func (d *permissionDao) WherePidIn(pid []int64) *permissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "pid", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(pid)})
	return d
}
func (d *permissionDao) WherePidNotIn(pid []int64) *permissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "pid", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(pid)})))
	return d
}
func (d *permissionDao) WherePidLt(pid int64) *permissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "pid", Value: pid})
	return d
}
func (d *permissionDao) WherePidLte(pid int64) *permissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "pid", Value: pid})
	return d
}
func (d *permissionDao) WherePidGt(pid int64) *permissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "pid", Value: pid})
	return d
}
func (d *permissionDao) WherePidGte(pid int64) *permissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "pid", Value: pid})
	return d
}
func (d *permissionDao) WherePidBetween(left int64, right int64) *permissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "pid Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *permissionDao) WhereSortEq(sort int) *permissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "sort", Value: sort})
	return d
}
func (d *permissionDao) WhereSortNeq(sort int) *permissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "sort", Value: sort})
	return d
}
func (d *permissionDao) WhereSortIn(sort []int) *permissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "sort", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(sort)})
	return d
}
func (d *permissionDao) WhereSortNotIn(sort []int) *permissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "sort", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(sort)})))
	return d
}
func (d *permissionDao) WhereSortLt(sort int) *permissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "sort", Value: sort})
	return d
}
func (d *permissionDao) WhereSortLte(sort int) *permissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "sort", Value: sort})
	return d
}
func (d *permissionDao) WhereSortGt(sort int) *permissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "sort", Value: sort})
	return d
}
func (d *permissionDao) WhereSortGte(sort int) *permissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "sort", Value: sort})
	return d
}
func (d *permissionDao) WhereSortBetween(left int, right int) *permissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "sort Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}

func (d *permissionDao) OrderBySort() *permissionDao {
	d.DB = d.DB.Order("sort ASC")
	return d
}

func (d *permissionDao) Limit(limit int) *permissionDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *permissionDao) Offset(offset int) *permissionDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *permissionDao) Page(page, pageSize int) *permissionDao {
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

func (d *permissionDao) GetByID(id int64) (*model.Permission, error) {
	return d.First(id)
}

func (d *permissionDao) Find() ([]*model.Permission, int64, error) {
	var list []*model.Permission
	var total int64
	var query = d.Model(model.Permission{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *permissionDao) Count() (int64, error) {
	var total int64
	var query = d.Model(model.Permission{})

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *permissionDao) First(conds ...interface{}) (*model.Permission, error) {
	var record = &model.Permission{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (d *permissionDao) Create(record *model.Permission) error {
	return d.DB.Create(record).Error
}

func (d *permissionDao) Save(record *model.Permission) error {
	return d.DB.Save(record).Error
}

func (d *permissionDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.Permission{}, conds...).Error
}
