package internal

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	golangjwt "github.com/golang-jwt/jwt/v5"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/data"
	"github.com/moumou/server/gen/dao"
	"github.com/moumou/server/pkgs/config"
)

type LoginService struct {
	db *dao.Dao
}

func NewLoginService(db *dao.Dao) *LoginService {
	return &LoginService{
		db: db,
	}
}

func (svc *LoginService) FindUserByUserNameAndPassword(ctx context.Context, userName, password string) (*model.User, error) {
	user, err := svc.db.UserDao.WithContext(ctx).WhereUsernameEq(userName).First()
	if err != nil {
		return nil, err
	}
	createdAtStr := strconv.FormatInt(user.CreatedAt, 10)
	h := hmac.New(sha256.New, []byte(createdAtStr))

	h.Write([]byte(password))
	encryptPwd := hex.EncodeToString(h.Sum(nil))

	if user.Password != encryptPwd {
		return nil, errors.New("密码错误")
	}

	return user, nil
}

func (svc *LoginService) FindUserByToken(ctx context.Context) (*model.User, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, errors.New("user nof found")
	}
	customClaims, ok := claims.(*data.CustomClaims)
	if !ok {
		return nil, errors.New("parse token fail")
	}

	var userID = customClaims.UserID
	return svc.db.UserDao.WithContext(ctx).GetByID(userID)
}

func (svc *LoginService) CreateToken(userInfo *model.User) (string, error) {
	claims := &data.CustomClaims{
		UserID: userInfo.ID,
		RegisteredClaims: golangjwt.RegisteredClaims{
			ExpiresAt: golangjwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  golangjwt.NewNumericDate(time.Now()),
			NotBefore: golangjwt.NewNumericDate(time.Now()),
			ID:        strconv.FormatInt(userInfo.ID, 10),
		},
	}

	var cnf data.SecurityConfig
	err := config.GetConfig(cnf.GetConfigName(), &cnf)
	if err != nil {
		return "", err
	}

	token := golangjwt.NewWithClaims(golangjwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cnf.JWTKey))
}

func (svc *LoginService) DeleteToken(token string) error {
	if token == "" {
		// TODO
		return errors.New("token不存在")
	}
	return nil
}
