package user

import (
	"context"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/internal"
	"github.com/moumou/server/biz/service/user/param"
	"gorm.io/gorm"
)

type Service struct {
	loginSvc  *internal.LoginService
	manageSvc *internal.ManageService
}

func NewUserService(db *gorm.DB) *Service {
	return &Service{
		loginSvc:  internal.NewLoginService(db),
		manageSvc: internal.NewManageService(db),
	}
}

func (svc *Service) Login(ctx context.Context, userName, password string) (token string, userInfo *model.User, err error) {
	userInfo, err = svc.loginSvc.FindUserByUserNameAndPassword(userName, password)
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
	return svc.loginSvc.FindUserByToken("")
}

func (svc *Service) GetUserList(ctx context.Context, filter *param.ListUserFilter, currentPage, pageSize int) ([]*model.User, int64, error) {
	return svc.manageSvc.List(ctx, filter, currentPage, pageSize)
}

func (svc *Service) GetUserInfo(ctx context.Context, id int) (*model.User, error) {
	return svc.manageSvc.GetByID(ctx, id)
}
