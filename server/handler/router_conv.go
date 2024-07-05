package handler

import (
	"github.com/moumou/server/biz/model"
	api "github.com/moumou/server/proto_gen"
)

func ConvRouterList2ResponseData(routerList []*model.Router) *api.GetRouterListResponseData {
	return &api.GetRouterListResponseData{
		List: convRouterListByPid(routerList, 0),
	}
}

func convRouterListByPid(routerList []*model.Router, pid int) []*api.Router {
	routerModelList := make([]*api.Router, 0, len(routerList))
	for _, routerInfo := range routerList {
		if int(routerInfo.Pid) != pid {
			continue
		}

		routerModel := ConvRouter2VO(routerInfo)
		routerModel.Children = convRouterListByPid(routerList, int(routerModel.Id))
		routerModelList = append(routerModelList, routerModel)
	}
	if len(routerModelList) == 0 {
		return nil
	}
	return routerModelList
}

func ConvRouter2VO(routerInfo *model.Router) *api.Router {
	return &api.Router{
		Id:        int64(routerInfo.ID),
		Name:      routerInfo.Name,
		Path:      routerInfo.Path,
		Title:     routerInfo.Title,
		IsMenu:    routerInfo.IsMenu,
		Pid:       int64(routerInfo.Pid),
		Sort:      int32(routerInfo.Sort),
		Component: routerInfo.Component,
	}
}
