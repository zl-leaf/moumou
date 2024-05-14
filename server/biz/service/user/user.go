package user

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/internal"
	"gorm.io/gorm"
)

type Service struct {
	loginSvc *internal.LoginService
}

func NewUserService(db *gorm.DB) *Service {
	return &Service{
		loginSvc: internal.NewLoginService(db),
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
