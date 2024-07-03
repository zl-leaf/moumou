package vo

import "github.com/moumou/server/biz/model"

func ConvSysUserList2ResponseData(sysUserList []*model.SysUser, total int64) *GetSysUserListResponseData {
	respData := &GetSysUserListResponseData{
		List: make([]*UserModel, 0, len(sysUserList)),
		BaseResponsePageData: BaseResponsePageData{
			Total: total,
		},
	}
	for _, sysUser := range sysUserList {
		respData.List = append(respData.List, ConvSysUser2VO(sysUser))
	}
	return respData
}

func ConvSysUser2VO(sysUser *model.SysUser) *UserModel {
	return &UserModel{SysUser: *sysUser}
}
