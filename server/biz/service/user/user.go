package user

import (
	"context"
	"time"

	"github.com/moumou/server/biz/util/ctxutil"

	"github.com/moumou/server/biz/conf"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/data"
	"github.com/moumou/server/biz/service/user/internal"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	captchaSvc *internal.CaptchaService
	loginSvc   *internal.LoginService
	db         *dao.Dao
}

func NewUserService(cnf *conf.Data, db *dao.Dao) *Service {
	return &Service{
		captchaSvc: internal.NewCaptchaService(),
		loginSvc:   internal.NewLoginService(cnf, db),
		db:         db,
	}
}

func (svc *Service) RandomCaptcha() (string, []byte, error) {
	return svc.captchaSvc.RandomImage()
}

func (svc *Service) VerifyCaptcha(ctx context.Context, captchaId, captcha string) error {
	return svc.captchaSvc.Verify(captchaId, captcha)
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

func (svc *Service) VerifyToken(ctx context.Context) (int64, error) {
	return svc.loginSvc.VerifyToken(ctx)
}

func (svc *Service) Self(ctx context.Context) (*model.User, error) {
	userId := ctxutil.GetCtxUserID(ctx)
	return svc.db.UserDao(ctx).GetByID(userId)
}

func (svc *Service) GetUserList(ctx context.Context, filter *data.ListUserFilter, currentPage, pageSize int) ([]*model.User, int64, error) {
	query := svc.db.UserDao(ctx)
	if filter.UsernameLike != nil {
		query = query.WhereUsernameLike(*filter.UsernameLike)
	}
	return query.Page(currentPage, pageSize).Find()
}

func (svc *Service) UpdateUser(ctx context.Context, userInfo *model.User) error {
	err := svc.db.UserDao(ctx).Save(userInfo)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) CreateUser(ctx context.Context, userInfo *model.User) error {
	userInfo.CreatedAt = time.Now().Unix()
	userInfo.Password = internal.EncryptPassword(userInfo.Password, userInfo.CreatedAt)
	err := svc.db.UserDao(ctx).Create(userInfo)
	if err != nil {
		return err
	}
	return nil
}
