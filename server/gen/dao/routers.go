package dao

import (
	"context"

	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type routerDao struct {
	*gorm.DB
}

func newRouterDao(db *gorm.DB) *routerDao {
	return &routerDao{db}
}

func (d *routerDao) WithContext(ctx context.Context) *routerDao {
	d.DB = d.DB.WithContext(ctx)
	return d
}

func (d *routerDao) WhereIDEq(ID int) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "id", Value: ID})
	return d
}
func (d *routerDao) WhereIDNeq(ID int) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "id", Value: ID})
	return d
}
func (d *routerDao) WhereIDIn(ID []int) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "id", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(ID)})
	return d
}
func (d *routerDao) WhereIDNotIn(ID []int) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "id", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(ID)})))
	return d
}
func (d *routerDao) WhereIDLt(ID int) *routerDao {
	d.DB = d.DB.Where(clause.Lt{Column: "id", Value: ID})
	return d
}
func (d *routerDao) WhereIDLte(ID int) *routerDao {
	d.DB = d.DB.Where(clause.Lte{Column: "id", Value: ID})
	return d
}
func (d *routerDao) WhereIDGt(ID int) *routerDao {
	d.DB = d.DB.Where(clause.Gt{Column: "id", Value: ID})
	return d
}
func (d *routerDao) WhereIDGte(ID int) *routerDao {
	d.DB = d.DB.Where(clause.Gte{Column: "id", Value: ID})
	return d
}
func (d *routerDao) WhereIDBetween(left int, right int) *routerDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *routerDao) WhereCreatedAtEq(createdAt int) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "created_at", Value: createdAt})
	return d
}
func (d *routerDao) WhereCreatedAtNeq(createdAt int) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "created_at", Value: createdAt})
	return d
}
func (d *routerDao) WhereCreatedAtIn(createdAt []int) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "created_at", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})
	return d
}
func (d *routerDao) WhereCreatedAtNotIn(createdAt []int) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "created_at", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})))
	return d
}
func (d *routerDao) WhereCreatedAtLt(createdAt int) *routerDao {
	d.DB = d.DB.Where(clause.Lt{Column: "created_at", Value: createdAt})
	return d
}
func (d *routerDao) WhereCreatedAtLte(createdAt int) *routerDao {
	d.DB = d.DB.Where(clause.Lte{Column: "created_at", Value: createdAt})
	return d
}
func (d *routerDao) WhereCreatedAtGt(createdAt int) *routerDao {
	d.DB = d.DB.Where(clause.Gt{Column: "created_at", Value: createdAt})
	return d
}
func (d *routerDao) WhereCreatedAtGte(createdAt int) *routerDao {
	d.DB = d.DB.Where(clause.Gte{Column: "created_at", Value: createdAt})
	return d
}
func (d *routerDao) WhereCreatedAtBetween(left int, right int) *routerDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "created_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *routerDao) WhereUpdatedAtEq(updatedAt int) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *routerDao) WhereUpdatedAtNeq(updatedAt int) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *routerDao) WhereUpdatedAtIn(updatedAt []int) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "updated_at", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})
	return d
}
func (d *routerDao) WhereUpdatedAtNotIn(updatedAt []int) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "updated_at", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})))
	return d
}
func (d *routerDao) WhereUpdatedAtLt(updatedAt int) *routerDao {
	d.DB = d.DB.Where(clause.Lt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *routerDao) WhereUpdatedAtLte(updatedAt int) *routerDao {
	d.DB = d.DB.Where(clause.Lte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *routerDao) WhereUpdatedAtGt(updatedAt int) *routerDao {
	d.DB = d.DB.Where(clause.Gt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *routerDao) WhereUpdatedAtGte(updatedAt int) *routerDao {
	d.DB = d.DB.Where(clause.Gte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *routerDao) WhereUpdatedAtBetween(left int, right int) *routerDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "updated_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *routerDao) WhereDeletedAtEq(deletedAt uint) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "deleted_at", Value: deletedAt})
	return d
}
func (d *routerDao) WhereDeletedAtNeq(deletedAt uint) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "deleted_at", Value: deletedAt})
	return d
}
func (d *routerDao) WhereDeletedAtIn(deletedAt []uint) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "deleted_at", Values: func(v []uint) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(deletedAt)})
	return d
}
func (d *routerDao) WhereDeletedAtNotIn(deletedAt []uint) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "deleted_at", Values: func(v []uint) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(deletedAt)})))
	return d
}
func (d *routerDao) WhereDeletedAtLt(deletedAt uint) *routerDao {
	d.DB = d.DB.Where(clause.Lt{Column: "deleted_at", Value: deletedAt})
	return d
}
func (d *routerDao) WhereDeletedAtLte(deletedAt uint) *routerDao {
	d.DB = d.DB.Where(clause.Lte{Column: "deleted_at", Value: deletedAt})
	return d
}
func (d *routerDao) WhereDeletedAtGt(deletedAt uint) *routerDao {
	d.DB = d.DB.Where(clause.Gt{Column: "deleted_at", Value: deletedAt})
	return d
}
func (d *routerDao) WhereDeletedAtGte(deletedAt uint) *routerDao {
	d.DB = d.DB.Where(clause.Gte{Column: "deleted_at", Value: deletedAt})
	return d
}
func (d *routerDao) WhereDeletedAtBetween(left uint, right uint) *routerDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "deleted_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *routerDao) WhereNameEq(name string) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "name", Value: name})
	return d
}
func (d *routerDao) WhereNameNeq(name string) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "name", Value: name})
	return d
}
func (d *routerDao) WhereNameLike(name string) *routerDao {
	d.DB = d.DB.Where(clause.Like{Column: "name", Value: "%" + name + "%"})
	return d
}
func (d *routerDao) WhereNamePrefixLike(name string) *routerDao {
	d.DB = d.DB.Where(clause.Like{Column: "name", Value: name + "%"})
	return d
}
func (d *routerDao) WhereNameNotLike(name string) *routerDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "name", Value: "%" + name + "%"}))
	return d
}
func (d *routerDao) WhereNameIn(name []string) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "name", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(name)})
	return d
}
func (d *routerDao) WhereNameNotIn(name []string) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "name", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(name)})))
	return d
}
func (d *routerDao) WherePathEq(path string) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "path", Value: path})
	return d
}
func (d *routerDao) WherePathNeq(path string) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "path", Value: path})
	return d
}
func (d *routerDao) WherePathLike(path string) *routerDao {
	d.DB = d.DB.Where(clause.Like{Column: "path", Value: "%" + path + "%"})
	return d
}
func (d *routerDao) WherePathPrefixLike(path string) *routerDao {
	d.DB = d.DB.Where(clause.Like{Column: "path", Value: path + "%"})
	return d
}
func (d *routerDao) WherePathNotLike(path string) *routerDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "path", Value: "%" + path + "%"}))
	return d
}
func (d *routerDao) WherePathIn(path []string) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "path", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(path)})
	return d
}
func (d *routerDao) WherePathNotIn(path []string) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "path", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(path)})))
	return d
}
func (d *routerDao) WhereTitleEq(title string) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "title", Value: title})
	return d
}
func (d *routerDao) WhereTitleNeq(title string) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "title", Value: title})
	return d
}
func (d *routerDao) WhereTitleLike(title string) *routerDao {
	d.DB = d.DB.Where(clause.Like{Column: "title", Value: "%" + title + "%"})
	return d
}
func (d *routerDao) WhereTitlePrefixLike(title string) *routerDao {
	d.DB = d.DB.Where(clause.Like{Column: "title", Value: title + "%"})
	return d
}
func (d *routerDao) WhereTitleNotLike(title string) *routerDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "title", Value: "%" + title + "%"}))
	return d
}
func (d *routerDao) WhereTitleIn(title []string) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "title", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(title)})
	return d
}
func (d *routerDao) WhereTitleNotIn(title []string) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "title", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(title)})))
	return d
}
func (d *routerDao) WherePidEq(pid uint) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "pid", Value: pid})
	return d
}
func (d *routerDao) WherePidNeq(pid uint) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "pid", Value: pid})
	return d
}
func (d *routerDao) WherePidIn(pid []uint) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "pid", Values: func(v []uint) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(pid)})
	return d
}
func (d *routerDao) WherePidNotIn(pid []uint) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "pid", Values: func(v []uint) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(pid)})))
	return d
}
func (d *routerDao) WherePidLt(pid uint) *routerDao {
	d.DB = d.DB.Where(clause.Lt{Column: "pid", Value: pid})
	return d
}
func (d *routerDao) WherePidLte(pid uint) *routerDao {
	d.DB = d.DB.Where(clause.Lte{Column: "pid", Value: pid})
	return d
}
func (d *routerDao) WherePidGt(pid uint) *routerDao {
	d.DB = d.DB.Where(clause.Gt{Column: "pid", Value: pid})
	return d
}
func (d *routerDao) WherePidGte(pid uint) *routerDao {
	d.DB = d.DB.Where(clause.Gte{Column: "pid", Value: pid})
	return d
}
func (d *routerDao) WherePidBetween(left uint, right uint) *routerDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "pid Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *routerDao) WhereSortEq(sort int) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "sort", Value: sort})
	return d
}
func (d *routerDao) WhereSortNeq(sort int) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "sort", Value: sort})
	return d
}
func (d *routerDao) WhereSortIn(sort []int) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "sort", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(sort)})
	return d
}
func (d *routerDao) WhereSortNotIn(sort []int) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "sort", Values: func(v []int) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(sort)})))
	return d
}
func (d *routerDao) WhereSortLt(sort int) *routerDao {
	d.DB = d.DB.Where(clause.Lt{Column: "sort", Value: sort})
	return d
}
func (d *routerDao) WhereSortLte(sort int) *routerDao {
	d.DB = d.DB.Where(clause.Lte{Column: "sort", Value: sort})
	return d
}
func (d *routerDao) WhereSortGt(sort int) *routerDao {
	d.DB = d.DB.Where(clause.Gt{Column: "sort", Value: sort})
	return d
}
func (d *routerDao) WhereSortGte(sort int) *routerDao {
	d.DB = d.DB.Where(clause.Gte{Column: "sort", Value: sort})
	return d
}
func (d *routerDao) WhereSortBetween(left int, right int) *routerDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "sort Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *routerDao) WhereComponentEq(component string) *routerDao {
	d.DB = d.DB.Where(clause.Eq{Column: "component", Value: component})
	return d
}
func (d *routerDao) WhereComponentNeq(component string) *routerDao {
	d.DB = d.DB.Where(clause.Neq{Column: "component", Value: component})
	return d
}
func (d *routerDao) WhereComponentLike(component string) *routerDao {
	d.DB = d.DB.Where(clause.Like{Column: "component", Value: "%" + component + "%"})
	return d
}
func (d *routerDao) WhereComponentPrefixLike(component string) *routerDao {
	d.DB = d.DB.Where(clause.Like{Column: "component", Value: component + "%"})
	return d
}
func (d *routerDao) WhereComponentNotLike(component string) *routerDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "component", Value: "%" + component + "%"}))
	return d
}
func (d *routerDao) WhereComponentIn(component []string) *routerDao {
	d.DB = d.DB.Where(clause.IN{Column: "component", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(component)})
	return d
}
func (d *routerDao) WhereComponentNotIn(component []string) *routerDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "component", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(component)})))
	return d
}

func (d *routerDao) OrderBySort() *routerDao {
	d.DB = d.DB.Order("sort ASC")
	return d
}

func (d *routerDao) Limit(limit int) *routerDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *routerDao) Offset(offset int) *routerDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *routerDao) GetByID(id int64) (*model.Router, error) {
	return d.First(id)
}

func (d *routerDao) Find() ([]*model.Router, int64, error) {
	var list []*model.Router
	var total int64
	var query = d.Model(model.Router{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *routerDao) First(conds ...interface{}) (*model.Router, error) {
	var record = &model.Router{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (d *routerDao) Create(record *model.Router) error {
	return d.DB.Create(record).Error
}

func (d *routerDao) Save(record *model.Router) error {
	return d.DB.Save(record).Error
}

func (d *routerDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.Router{}, conds...).Error
}
