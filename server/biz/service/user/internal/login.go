package internal

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/param"
	"github.com/moumou/server/gen/dao"
	"github.com/moumou/server/pkgs/config"
	"strconv"
	"time"
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

func (svc *LoginService) FindUserByToken(token string) (*model.User, error) {
	// TODO
	if token != "token" {
		return nil, errors.New("token异常")
	}

	var userID = int64(1)
	return svc.db.UserDao.GetByID(userID)
}

func (svc *LoginService) CreateToken(userInfo *model.User) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        strconv.FormatInt(userInfo.ID, 10),
	}

	var cnf param.SecurityConfig
	err := config.GetConfig(cnf.GetConfigName(), &cnf)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cnf.JWTKey))
}

func (svc *LoginService) DeleteToken(token string) error {
	if token == "" {
		// TODO
		return errors.New("token不存在")
	}
	return nil
}
