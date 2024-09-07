package user

import (
	"context"

	"github.com/moumou/server/biz/conf"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/data"
	"github.com/moumou/server/biz/service/user/internal"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	loginSvc *internal.LoginService
	db       *dao.Dao
}

func NewUserService(cnf *conf.Data, db *dao.Dao) *Service {
	return &Service{
		loginSvc: internal.NewLoginService(cnf, db),
		db:       db,
	}
}

func (svc *Service) Login(ctx context.Context, userName, password string) (token string, userInfo *model.User, err error) {
	userInfo, err = svc.loginSvc.FindUserByUserNameAndPassword(ctx, userName, password)
	if err != nil {
		return
	}

	token, err = svc.loginSvc.CreateToken(userInfo)
	if err != nil {
		return "", nil, err
	}
	return
}

func (svc *Service) Logout(ctx context.Context, token string) error {
	return svc.loginSvc.DeleteToken(token)
}

func (svc *Service) Self(ctx context.Context) (*model.User, error) {
	return svc.loginSvc.FindUserByToken(ctx)
}

func (svc *Service) GetUserList(ctx context.Context, filter *data.ListUserFilter, currentPage, pageSize int) ([]*model.User, int64, error) {
	query := svc.db.UserDao(ctx)
	if filter.UsernameLike != nil {
		query = query.WhereUsernameLike(*filter.UsernameLike)
	}
	return query.Page(currentPage, pageSize).Find()
}
