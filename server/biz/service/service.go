package service

import "github.com/moumou/server/biz/service/user"

type Service struct {
	UserService *user.Service
}

func NewService(
	userService *user.Service,
) *Service {
	return &Service{
		UserService: userService,
	}
}
