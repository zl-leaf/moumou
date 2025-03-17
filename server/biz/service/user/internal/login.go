package internal

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	golangjwt "github.com/golang-jwt/jwt/v5"
	"github.com/moumou/server/biz/conf"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/data"
	"github.com/moumou/server/gen/dao"
)

type LoginService struct {
	cnf *conf.Data
	db  *dao.Dao
}

func NewLoginService(cnf *conf.Data, db *dao.Dao) *LoginService {
	return &LoginService{
		cnf: cnf,
		db:  db,
	}
}

func (svc *LoginService) FindUserByUserNameAndPassword(ctx context.Context, userName, password string) (*model.User, error) {
	user, err := svc.db.UserDao(ctx).WhereUsernameEq(userName).First()
	if err != nil {
		return nil, err
	}
	if EncryptPassword(password, user.CreatedAt) != user.Password {
		return nil, errors.New("密码错误")
	}

	return user, nil
}

func (svc *LoginService) VerifyToken(ctx context.Context) (int64, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return 0, errors.New("user nof found")
	}
	customClaims, ok := claims.(*data.CustomClaims)
	if !ok {
		return 0, errors.New("parse token fail")
	}

	var userID = customClaims.UserID
	check, err := svc.db.UserDao(ctx).WhereIdEq(userID).Count()
	if err != nil {
		return 0, err
	}
	if check == 0 {
		return 0, errors.New("user nof found")
	}

	return userID, nil
}

func (svc *LoginService) CreateToken(userInfo *model.User) (string, error) {
	claims := &data.CustomClaims{
		UserID: userInfo.Id,
		RegisteredClaims: golangjwt.RegisteredClaims{
			ExpiresAt: golangjwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  golangjwt.NewNumericDate(time.Now()),
			NotBefore: golangjwt.NewNumericDate(time.Now()),
			ID:        strconv.FormatInt(userInfo.Id, 10),
		},
	}

	token := golangjwt.NewWithClaims(golangjwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(svc.cnf.SecurityConfig.JWTKey))
}

func (svc *LoginService) DeleteToken(token string) error {
	if token == "" {
		// TODO
		return errors.New("token不存在")
	}
	return nil
}
