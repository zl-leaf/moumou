package vo

import "github.com/moumou/server/biz/model"

type RouterModel struct {
	model.SysRouter
	Children []*RouterModel `json:"children"`
}

type GetSysRouterListRequest struct {
	BaseRequest
}

type GetSysRouterListResponseData struct {
	List []*RouterModel `json:"list"`
}

type GetSysRouterListResponse struct {
	BaseResponse
	Data *GetSysRouterListResponseData `json:"data"`
}

type GetSysRouterInfoRequest struct {
	BaseRequest
	ID int `json:"id"`
}

type GetSysRouterInfoResponse struct {
	BaseResponse
	Data *RouterModel `json:"data"`
}
