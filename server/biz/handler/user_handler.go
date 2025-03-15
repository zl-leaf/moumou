package handler

import (
	"context"

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
