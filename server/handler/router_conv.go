package handler

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/router/data"
	api "github.com/moumou/server/gen/proto"
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
		Id:        routerInfo.ID,
		Name:      routerInfo.Name,
		Path:      routerInfo.Path,
		Title:     routerInfo.Title,
		IsMenu:    routerInfo.IsMenu,
		IsSystem:  routerInfo.IsSystem,
		Pid:       int64(routerInfo.Pid),
		Sort:      int32(routerInfo.Sort),
		Component: routerInfo.Component,
	}
}

func ConvCreateRequestData2FormData(vo *api.CreateRouterRequestData) *data.RouterFormData {
	return &data.RouterFormData{
		Router: model.Router{
			Name:      vo.Name,
			Path:      vo.Path,
			Title:     vo.Title,
			IsMenu:    vo.IsMenu,
			Pid:       uint(vo.Pid),
			Sort:      int(vo.Sort),
			Component: vo.Component,
		},
	}
}

func ConvUpdateRequestData2FormData(vo *api.UpdateRouterRequestData) *data.RouterFormData {
	return &data.RouterFormData{
		Router: model.Router{
			BaseModel: model.BaseModel{ID: vo.Id},
			Name:      vo.Name,
			Path:      vo.Path,
			Title:     vo.Title,
			IsMenu:    vo.IsMenu,
			Pid:       uint(vo.Pid),
			Sort:      int(vo.Sort),
			Component: vo.Component,
		},
	}
}
