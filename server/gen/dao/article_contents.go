package dao

import (
	"context"

	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type articleContentDao struct {
	*gorm.DB
}

func newArticleContentDao(db *gorm.DB) *articleContentDao {
	return &articleContentDao{db}
}

func (d *articleContentDao) WithContext(ctx context.Context) *articleContentDao {
	return &articleContentDao{d.DB.WithContext(ctx)}
}

func (d *articleContentDao) WhereIdEq(id int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Eq{Column: "id", Value: id})
	return d
}
func (d *articleContentDao) WhereIdNeq(id int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Neq{Column: "id", Value: id})
	return d
}
func (d *articleContentDao) WhereIdIn(id []int64) *articleContentDao {
	d.DB = d.DB.Where(clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})
	return d
}
func (d *articleContentDao) WhereIdNotIn(id []int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})))
	return d
}
func (d *articleContentDao) WhereIdLt(id int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Lt{Column: "id", Value: id})
	return d
}
func (d *articleContentDao) WhereIdLte(id int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Lte{Column: "id", Value: id})
	return d
}
func (d *articleContentDao) WhereIdGt(id int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Gt{Column: "id", Value: id})
	return d
}
func (d *articleContentDao) WhereIdGte(id int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Gte{Column: "id", Value: id})
	return d
}
func (d *articleContentDao) WhereIdBetween(left int64, right int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *articleContentDao) WhereCreatedAtEq(createdAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Eq{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleContentDao) WhereCreatedAtNeq(createdAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Neq{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleContentDao) WhereCreatedAtIn(createdAt []int64) *articleContentDao {
	d.DB = d.DB.Where(clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})
	return d
}
func (d *articleContentDao) WhereCreatedAtNotIn(createdAt []int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})))
	return d
}
func (d *articleContentDao) WhereCreatedAtLt(createdAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Lt{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleContentDao) WhereCreatedAtLte(createdAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Lte{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleContentDao) WhereCreatedAtGt(createdAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Gt{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleContentDao) WhereCreatedAtGte(createdAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Gte{Column: "created_at", Value: createdAt})
	return d
}
func (d *articleContentDao) WhereCreatedAtBetween(left int64, right int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "created_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *articleContentDao) WhereUpdatedAtEq(updatedAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Eq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleContentDao) WhereUpdatedAtNeq(updatedAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Neq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleContentDao) WhereUpdatedAtIn(updatedAt []int64) *articleContentDao {
	d.DB = d.DB.Where(clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})
	return d
}
func (d *articleContentDao) WhereUpdatedAtNotIn(updatedAt []int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})))
	return d
}
func (d *articleContentDao) WhereUpdatedAtLt(updatedAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Lt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleContentDao) WhereUpdatedAtLte(updatedAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Lte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleContentDao) WhereUpdatedAtGt(updatedAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Gt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleContentDao) WhereUpdatedAtGte(updatedAt int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Gte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *articleContentDao) WhereUpdatedAtBetween(left int64, right int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "updated_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *articleContentDao) WhereArticleIDEq(articleID int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Eq{Column: "article_id", Value: articleID})
	return d
}
func (d *articleContentDao) WhereArticleIDNeq(articleID int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Neq{Column: "article_id", Value: articleID})
	return d
}
func (d *articleContentDao) WhereArticleIDIn(articleID []int64) *articleContentDao {
	d.DB = d.DB.Where(clause.IN{Column: "article_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(articleID)})
	return d
}
func (d *articleContentDao) WhereArticleIDNotIn(articleID []int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "article_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(articleID)})))
	return d
}
func (d *articleContentDao) WhereArticleIDLt(articleID int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Lt{Column: "article_id", Value: articleID})
	return d
}
func (d *articleContentDao) WhereArticleIDLte(articleID int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Lte{Column: "article_id", Value: articleID})
	return d
}
func (d *articleContentDao) WhereArticleIDGt(articleID int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Gt{Column: "article_id", Value: articleID})
	return d
}
func (d *articleContentDao) WhereArticleIDGte(articleID int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Gte{Column: "article_id", Value: articleID})
	return d
}
func (d *articleContentDao) WhereArticleIDBetween(left int64, right int64) *articleContentDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "article_id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *articleContentDao) WhereContentEq(content string) *articleContentDao {
	d.DB = d.DB.Where(clause.Eq{Column: "content", Value: content})
	return d
}
func (d *articleContentDao) WhereContentNeq(content string) *articleContentDao {
	d.DB = d.DB.Where(clause.Neq{Column: "content", Value: content})
	return d
}
func (d *articleContentDao) WhereContentLike(content string) *articleContentDao {
	d.DB = d.DB.Where(clause.Like{Column: "content", Value: "%" + content + "%"})
	return d
}
func (d *articleContentDao) WhereContentPrefixLike(content string) *articleContentDao {
	d.DB = d.DB.Where(clause.Like{Column: "content", Value: content + "%"})
	return d
}
func (d *articleContentDao) WhereContentNotLike(content string) *articleContentDao {
	d.DB = d.DB.Where(clause.Not(clause.Like{Column: "content", Value: "%" + content + "%"}))
	return d
}
func (d *articleContentDao) WhereContentIn(content []string) *articleContentDao {
	d.DB = d.DB.Where(clause.IN{Column: "content", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(content)})
	return d
}
func (d *articleContentDao) WhereContentNotIn(content []string) *articleContentDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "content", Values: func(v []string) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(content)})))
	return d
}

func (d *articleContentDao) Limit(limit int) *articleContentDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *articleContentDao) Offset(offset int) *articleContentDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *articleContentDao) Page(page, pageSize int) *articleContentDao {
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

func (d *articleContentDao) GetByID(id int64) (*model.ArticleContent, error) {
	return d.First(id)
}

func (d *articleContentDao) Find() ([]*model.ArticleContent, int64, error) {
	var list []*model.ArticleContent
	var total int64
	var query = d.Model(model.ArticleContent{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *articleContentDao) Count() (int64, error) {
	var total int64
	var query = d.Model(model.ArticleContent{})

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *articleContentDao) First(conds ...interface{}) (*model.ArticleContent, error) {
	var record = &model.ArticleContent{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (d *articleContentDao) Create(record *model.ArticleContent) error {
	return d.DB.Create(record).Error
}

func (d *articleContentDao) Save(record *model.ArticleContent) error {
	return d.DB.Save(record).Error
}

func (d *articleContentDao) SaveFullAssociations(record *model.ArticleContent) error {
	return d.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(record).Error
}

func (d *articleContentDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.ArticleContent{}, conds...).Error
}
