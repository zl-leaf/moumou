package handler

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/data"
	api "github.com/moumou/server/gen/proto"
)

func ConvVO2UserListFilter(filterVO *api.GetUserListRequestFilter) *data.ListUserFilter {
	return &data.ListUserFilter{
		UsernameLike: filterVO.UsernameLike,
	}
}

func ConvUserList2RespData(userList []*model.User, total int64) *api.GetUserListResponseData {

	return &api.GetUserListResponseData{
		Total: total,
		List:  ConvUserList2VOList(userList),
	}
}

func ConvUserList2VOList(userModelList []*model.User) []*api.User {
	userVOList := make([]*api.User, 0, len(userModelList))
	for _, userinfo := range userModelList {
		userVOList = append(userVOList, ConvUser2VO(userinfo))
	}
	return userVOList
}

func ConvUser2VO(userInfo *model.User) *api.User {
	return &api.User{
		Id:       userInfo.ID,
		Username: userInfo.Username,
	}
}
