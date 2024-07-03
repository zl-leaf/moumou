package vo

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/sys_user/param"
)

type UserModel struct {
	model.SysUser
}

type GetSysUserListRequestFilter struct {
	param.SysUserFilter
}

type GetSysUserListRequest struct {
	BaseRequest
	BaseRequestPage
	Filter *GetSysUserListRequestFilter `json:"filter"`
}

type GetSysUserListResponseData struct {
	BaseResponsePageData
	List []*UserModel `json:"list"`
}

type GetSysUserListResponse struct {
	BaseResponse
	Data *GetSysUserListResponseData `json:"data"`
}

type GetSysUserInfoResponse struct {
	BaseResponse
	Data *UserModel `json:"data"`
}

type GetSysUserInfoRequest struct {
	BaseRequest
	ID int `json:"id"`
}
