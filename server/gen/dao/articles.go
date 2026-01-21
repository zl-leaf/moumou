package dao

import (
	"context"

	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type articleDao struct {
	*gorm.DB
}

func newArticleDao(db *gorm.DB) *articleDao {
	return &articleDao{db}
}

func (d *articleDao) WithContext(ctx context.Context) *articleDao {
	return &articleDao{d.DB.WithContext(ctx)}
}

func (d *articleDao) WhereIdEq(id int64) *articleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "id", Value: id})
	return d
}
func (d *articleDao) WhereIdNeq(id int64) *articleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "id", Value: id})
	return d
}
func (d *articleDao) WhereIdIn(id []int64) *articleDao {
	d.DB = d.DB.Where(clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})
	return d
}
func (d *articleDao) WhereIdNotIn(id []int64) *articleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})))
	return d
}
func (d *articleDao) WhereIdLt(id int64) *articleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "id", Value: id})
	return d
}
func (d *articleDao) WhereIdLte(id int64) *articleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "id", Value: id})
	return d
}
func (d *articleDao) WhereIdGt(id int64) *articleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "id", Value: id})
	return d
}
func (d *articleDao) WhereIdGte(id int64) *articleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "id", Value: id})
	return d
}
func (d *articleDao) WhereIdBetween(left int64, right int64) *articleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *articleDao) WhereCreatedAtEq(createdAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleDao) WhereCreatedAtNeq(createdAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleDao) WhereCreatedAtIn(createdAt []int64) *articleDao {
	d.DB = d.DB.Where(clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})
	return d
}
func (d *articleDao) WhereCreatedAtNotIn(createdAt []int64) *articleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})))
	return d
}
func (d *articleDao) WhereCreatedAtLt(createdAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleDao) WhereCreatedAtLte(createdAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleDao) WhereCreatedAtGt(createdAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleDao) WhereCreatedAtGte(createdAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleDao) WhereCreatedAtBetween(left int64, right int64) *articleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "created_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *articleDao) WhereUpdatedAtEq(updatedAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleDao) WhereUpdatedAtNeq(updatedAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleDao) WhereUpdatedAtIn(updatedAt []int64) *articleDao {
	d.DB = d.DB.Where(clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})
	return d
}
func (d *articleDao) WhereUpdatedAtNotIn(updatedAt []int64) *articleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})))
	return d
}
func (d *articleDao) WhereUpdatedAtLt(updatedAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Lt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleDao) WhereUpdatedAtLte(updatedAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Lte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleDao) WhereUpdatedAtGt(updatedAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Gt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleDao) WhereUpdatedAtGte(updatedAt int64) *articleDao {
	d.DB = d.DB.Where(clause.Gte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleDao) WhereUpdatedAtBetween(left int64, right int64) *articleDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "updated_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *articleDao) WhereTitleEq(title string) *articleDao {
	d.DB = d.DB.Where(clause.Eq{Column: "title", Value: title})
	return d
}
func (d *articleDao) WhereTitleNeq(title string) *articleDao {
	d.DB = d.DB.Where(clause.Neq{Column: "title", Value: title})
	return d
}
func (d *articleDao) WhereTitleLike(title string) *articleDao {
	d.DB = d.DB.Where(clause.Like{Column: "title", Value: "%" + title + "%"})
	return d
}
func (d *articleDao) WhereTitlePrefixLike(title string) *articleDao {
	d.DB = d.DB.Where(clause.Like{Column: "title", Value: title + "%"})
	return d
}
func (d *articleDao) WhereTitleNotLike(title string) *articleDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "title", Value: "%" + title + "%"}))
	return d
}
func (d *articleDao) WhereTitleIn(title []string) *articleDao {
	d.DB = d.DB.Where(clause.IN{Column: "title", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(title)})
	return d
}
func (d *articleDao) WhereTitleNotIn(title []string) *articleDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "title", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(title)})))
	return d
}

func (d *articleDao) Limit(limit int) *articleDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *articleDao) Offset(offset int) *articleDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *articleDao) Page(page, pageSize int) *articleDao {
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

func (d *articleDao) PreloadArticleContent(args ...interface{}) *articleDao {
	d.DB = d.DB.Preload("ArticleContent", args)
	return d
}

func (d *articleDao) GetByID(id int64) (*model.Article, error) {
	return d.First(id)
}

func (d *articleDao) Find() ([]*model.Article, int64, error) {
	var list []*model.Article
	var total int64
	var query = d.Model(model.Article{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *articleDao) Count() (int64, error) {
	var total int64
	var query = d.Model(model.Article{})

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *articleDao) First(conds ...interface{}) (*model.Article, error) {
	var record = &model.Article{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (d *articleDao) Create(record *model.Article) error {
	return d.DB.Create(record).Error
}

func (d *articleDao) Save(record *model.Article) error {
	return d.DB.Save(record).Error
}

func (d *articleDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.Article{}, conds...).Error
}
