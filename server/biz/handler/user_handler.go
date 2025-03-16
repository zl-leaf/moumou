package handler

import (
	"context"

	"github.com/moumou/server/biz/model"

	"github.com/moumou/server/biz/conv"
	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type UserHandler struct {
	svc       *service.Service
	converter conv.IConverter
}

func NewUserHandler(svc *service.Service, converter conv.IConverter) api.UserHandlerHTTPServer {
	return &UserHandler{svc: svc, converter: converter}
}

func (h *UserHandler) CreateUser(ctx context.Context, request *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	userInfo := new(model.User)
	h.converter.ConvertCreateUserRequestDataToBO(request.GetUser(), userInfo)

	err := h.svc.UserService.CreateUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &api.CreateUserResponse{
		Data: &api.CreateUserResponseData{Id: userInfo.Id},
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, request *api.UpdateUserRequest) (*api.UpdateUserResponse, error) {
	userInfo, err := h.svc.Dao.UserDao(ctx).GetByID(request.GetUser().GetId())
	if err != nil {
		return nil, err
	}
	h.converter.ConvertUpdateUserRequestDataToBO(request.GetUser(), userInfo)
	err = h.svc.UserService.UpdateUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	return &api.UpdateUserResponse{}, nil
}

func (h *UserHandler) GetUserList(ctx context.Context, request *api.GetUserListRequest) (*api.GetUserListResponse, error) {
	filter := h.converter.ConvertGetUserListRequestFilter(request.GetFilter())
	userList, total, err := h.svc.UserService.GetUserList(ctx, filter, int(request.GetCurrentPage()), int(request.GetPageSize()))

	if err != nil {
		return nil, err
	}
	return &api.GetUserListResponse{
		Data: &api.GetUserListResponseData{
			Total: total,
			List:  h.converter.ConvertUserListToVO(userList),
		},
	}, nil
}

func (h *UserHandler) GetUserInfo(ctx context.Context, request *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	userInfo, err := h.svc.Dao.UserDao(ctx).GetByID(request.GetId())
	if err != nil {
		return nil, err
	}
	return &api.GetUserInfoResponse{
		Data: h.converter.ConvertUserToVO(userInfo),
	}, nil
}
