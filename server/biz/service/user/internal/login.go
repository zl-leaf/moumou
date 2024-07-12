package internal

import (
	"errors"
	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
)

type LoginService struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) *LoginService {
	return &LoginService{
		db: db,
	}
}

func (svc *LoginService) FindUserByUserNameAndPassword(userName, password string) (*model.User, error) {
	var user = &model.User{}
	var result = svc.db.Where("username = ?", userName).Where("password = ?", password).First(&user)
	// TODO record not found
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (svc *LoginService) FindUserByToken(token string) (*model.User, error) {
	// TODO
	if token != "token" {
		return nil, errors.New("token异常")
	}

	var userID = 1
	var user = &model.User{}
	var result = svc.db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (svc *LoginService) CreateToken(userInfo *model.User) string {
	return "token"
}

func (svc *LoginService) DeleteToken(token string) error {
	if token != "token" {
		return errors.New("token不存在")
	}
	return nil
}
