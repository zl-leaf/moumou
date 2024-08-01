package handler

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/param"
	api "github.com/moumou/server/gen/proto"
)

func ConvVO2UserListFilter(filterVO *api.GetUserListRequestFilter) *param.ListUserFilter {
	return &param.ListUserFilter{}
}

func ConvUserList2RespData(userList []*model.User, total int64) *api.GetUserListResponseData {
	userVOList := make([]*api.User, 0, len(userList))
	for _, userinfo := range userList {
		userVOList = append(userVOList, ConvUser2VO(userinfo))
	}

	return &api.GetUserListResponseData{
		Total: total,
		List:  userVOList,
	}
}

func ConvUser2VO(userInfo *model.User) *api.User {
	return &api.User{
		Id:       userInfo.ID,
		Username: userInfo.Username,
	}
}
