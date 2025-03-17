package dao

import (
	"context"

	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type rolePermissionDao struct {
	*gorm.DB
}

func newRolePermissionDao(db *gorm.DB) *rolePermissionDao {
	return &rolePermissionDao{db}
}

func (d *rolePermissionDao) WithContext(ctx context.Context) *rolePermissionDao {
	return &rolePermissionDao{d.DB.WithContext(ctx)}
}

func (d *rolePermissionDao) WhereIdEq(id int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "id", Value: id})
	return d
}
func (d *rolePermissionDao) WhereIdNeq(id int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "id", Value: id})
	return d
}
func (d *rolePermissionDao) WhereIdIn(id []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})
	return d
}
func (d *rolePermissionDao) WhereIdNotIn(id []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(id)})))
	return d
}
func (d *rolePermissionDao) WhereIdLt(id int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "id", Value: id})
	return d
}
func (d *rolePermissionDao) WhereIdLte(id int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "id", Value: id})
	return d
}
func (d *rolePermissionDao) WhereIdGt(id int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "id", Value: id})
	return d
}
func (d *rolePermissionDao) WhereIdGte(id int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "id", Value: id})
	return d
}
func (d *rolePermissionDao) WhereIdBetween(left int64, right int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *rolePermissionDao) WhereCreatedAtEq(createdAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "created_at", Value: createdAt})
	return d
}
func (d *rolePermissionDao) WhereCreatedAtNeq(createdAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "created_at", Value: createdAt})
	return d
}
func (d *rolePermissionDao) WhereCreatedAtIn(createdAt []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})
	return d
}
func (d *rolePermissionDao) WhereCreatedAtNotIn(createdAt []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "created_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(createdAt)})))
	return d
}
func (d *rolePermissionDao) WhereCreatedAtLt(createdAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "created_at", Value: createdAt})
	return d
}
func (d *rolePermissionDao) WhereCreatedAtLte(createdAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "created_at", Value: createdAt})
	return d
}
func (d *rolePermissionDao) WhereCreatedAtGt(createdAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "created_at", Value: createdAt})
	return d
}
func (d *rolePermissionDao) WhereCreatedAtGte(createdAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "created_at", Value: createdAt})
	return d
}
func (d *rolePermissionDao) WhereCreatedAtBetween(left int64, right int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "created_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtEq(updatedAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtNeq(updatedAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtIn(updatedAt []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtNotIn(updatedAt []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "updated_at", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(updatedAt)})))
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtLt(updatedAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtLte(updatedAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtGt(updatedAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtGte(updatedAt int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "updated_at", Value: updatedAt})
	return d
}
func (d *rolePermissionDao) WhereUpdatedAtBetween(left int64, right int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "updated_at Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *rolePermissionDao) WhereRoleIDEq(roleID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "role_id", Value: roleID})
	return d
}
func (d *rolePermissionDao) WhereRoleIDNeq(roleID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "role_id", Value: roleID})
	return d
}
func (d *rolePermissionDao) WhereRoleIDIn(roleID []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "role_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(roleID)})
	return d
}
func (d *rolePermissionDao) WhereRoleIDNotIn(roleID []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "role_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(roleID)})))
	return d
}
func (d *rolePermissionDao) WhereRoleIDLt(roleID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "role_id", Value: roleID})
	return d
}
func (d *rolePermissionDao) WhereRoleIDLte(roleID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "role_id", Value: roleID})
	return d
}
func (d *rolePermissionDao) WhereRoleIDGt(roleID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "role_id", Value: roleID})
	return d
}
func (d *rolePermissionDao) WhereRoleIDGte(roleID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "role_id", Value: roleID})
	return d
}
func (d *rolePermissionDao) WhereRoleIDBetween(left int64, right int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "role_id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}
func (d *rolePermissionDao) WherePermissionIDEq(permissionID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Eq{Column: "permission_id", Value: permissionID})
	return d
}
func (d *rolePermissionDao) WherePermissionIDNeq(permissionID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Neq{Column: "permission_id", Value: permissionID})
	return d
}
func (d *rolePermissionDao) WherePermissionIDIn(permissionID []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.IN{Column: "permission_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(permissionID)})
	return d
}
func (d *rolePermissionDao) WherePermissionIDNotIn(permissionID []int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Not((clause.IN{Column: "permission_id", Values: func(v []int64) []interface{} {
		ret := make([]interface{}, len(v))
		for i, item := range v {
			ret[i] = item
		}
		return ret
	}(permissionID)})))
	return d
}
func (d *rolePermissionDao) WherePermissionIDLt(permissionID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lt{Column: "permission_id", Value: permissionID})
	return d
}
func (d *rolePermissionDao) WherePermissionIDLte(permissionID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Lte{Column: "permission_id", Value: permissionID})
	return d
}
func (d *rolePermissionDao) WherePermissionIDGt(permissionID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gt{Column: "permission_id", Value: permissionID})
	return d
}
func (d *rolePermissionDao) WherePermissionIDGte(permissionID int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Gte{Column: "permission_id", Value: permissionID})
	return d
}
func (d *rolePermissionDao) WherePermissionIDBetween(left int64, right int64) *rolePermissionDao {
	d.DB = d.DB.Where(clause.Expr{SQL: "permission_id Between ? AND ?", Vars: []interface{}{left, right}})
	return d
}

func (d *rolePermissionDao) Limit(limit int) *rolePermissionDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *rolePermissionDao) Offset(offset int) *rolePermissionDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *rolePermissionDao) Page(page, pageSize int) *rolePermissionDao {
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

func (d *rolePermissionDao) GetByID(id int64) (*model.RolePermission, error) {
	return d.First(id)
}

func (d *rolePermissionDao) Find() ([]*model.RolePermission, int64, error) {
	var list []*model.RolePermission
	var total int64
	var query = d.Model(model.RolePermission{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *rolePermissionDao) Count() (int64, error) {
	var total int64
	var query = d.Model(model.RolePermission{})

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (d *rolePermissionDao) First(conds ...interface{}) (*model.RolePermission, error) {
	var record = &model.RolePermission{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}

func (d *rolePermissionDao) Create(record *model.RolePermission) error {
	return d.DB.Create(record).Error
}

func (d *rolePermissionDao) Save(record *model.RolePermission) error {
	return d.DB.Save(record).Error
}

func (d *rolePermissionDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.RolePermission{}, conds...).Error
}
