package handler

import (
	"context"

	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type UserHandler struct {
	svc *service.Service
}

func NewUserHandler(svc *service.Service) api.UserHandlerHTTPServer {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) GetUserList(ctx context.Context, request *api.GetUserListRequest) (*api.GetUserListResponse, error) {
	userList, total, err := h.svc.UserService.GetUserList(ctx, ConvVO2UserListFilter(request.GetFilter()), int(request.GetCurrentPage()), int(request.GetPageSize()))

	if err != nil {
		return nil, err
	}
	return &api.GetUserListResponse{
		Data: &api.GetUserListResponseData{
			Total: total,
			List:  ConvUserList2VOList(userList),
		},
	}, nil
}

func (h *UserHandler) GetUserInfo(ctx context.Context, request *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	userInfo, err := h.svc.Dao.UserDao.WithContext(ctx).GetByID(request.GetId())
	if err != nil {
		return nil, err
	}
	return &api.GetUserInfoResponse{
		Data: ConvUser2VO(userInfo),
	}, nil
}
