package handler

import (
	"github.com/moumou/server/biz/model"
	api "github.com/moumou/server/gen/proto"
)

func ConvRouterList2SecurityRouterTreeRespData(routerInfoList []*model.Router) *api.GetSecurityRouterTreeResponseData {
	return &api.GetSecurityRouterTreeResponseData{
		Routers: convRouterListByPid(routerInfoList, 0),
	}
}
