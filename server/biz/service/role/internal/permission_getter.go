package internal

import (
	"context"

	"github.com/moumou/server/biz/service/role/data"
	"github.com/moumou/server/gen/dao"
)

type PermissionGetterService struct {
	db *dao.Dao
}

func NewPermissionGetterService(db *dao.Dao) *PermissionGetterService {
	return &PermissionGetterService{db: db}
}

func (s *PermissionGetterService) GetPermissionTree(ctx context.Context) (*data.PermissionTree, error) {
	permissions, _, err := s.db.PermissionDao(ctx).OrderBySort().Find()
	if err != nil {
		return nil, err
	}
	return data.BuildPermissionTree(permissions), nil
}
