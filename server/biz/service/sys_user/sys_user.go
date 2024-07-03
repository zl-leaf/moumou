package sys_user

import (
	"context"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/sys_user/internal"
	"github.com/moumou/server/biz/service/sys_user/param"
	"gorm.io/gorm"
)

type Service struct {
	loginSvc  *internal.LoginService
	manageSvc *internal.ManageService
}

func NewSysUserService(db *gorm.DB) *Service {
	return &Service{
		loginSvc:  internal.NewLoginService(db),
		manageSvc: internal.NewManageService(db),
	}
}

func (svc *Service) Login(userName, password string) (token string, userInfo *model.SysUser, err error) {
	userInfo, err = svc.loginSvc.FindUserByUserNameAndPassword(userName, password)
	if err != nil {
		return
	}

	token = svc.loginSvc.CreateToken(userInfo)
	return
}

func (svc *Service) Logout(token string) error {
	return svc.loginSvc.DeleteToken(token)
}

func (svc *Service) Self(token string) (*model.SysUser, error) {
	return svc.loginSvc.FindUserByToken(token)
}

func (svc *Service) GetSysUserList(ctx context.Context, filter *param.SysUserFilter, currentPage, pageSize int) ([]*model.SysUser, int64, error) {
	return svc.manageSvc.List(ctx, filter, currentPage, pageSize)
}

func (svc *Service) GetSysUserInfo(ctx context.Context, id int) (*model.SysUser, error) {
	return svc.manageSvc.GetByID(ctx, id)
}
