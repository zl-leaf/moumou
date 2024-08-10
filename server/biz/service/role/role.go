package role

import (
	"context"

	"github.com/moumou/server/biz/service/role/internal"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	db *dao.Dao
}

func NewService(db *dao.Dao) *Service {
	return &Service{db: db}
}

func (s *Service) UpdateRolePermission(ctx context.Context, roleID int64, permissionIDs []int64) error {
	return internal.NewUpdateRolePermissionService(s.db).Execute(ctx, roleID, permissionIDs)
}
