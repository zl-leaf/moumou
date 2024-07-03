package vo

import "github.com/moumou/server/biz/model"

func ConvSysRouter2VO(sysRouter *model.SysRouter) *RouterModel {
	return &RouterModel{
		SysRouter: *sysRouter,
	}
}

func ConvSysRouterList2ResponseData(routerList []*model.SysRouter) *GetSysRouterListResponseData {
	data := &GetSysRouterListResponseData{
		List: convSysRouterListByPid(routerList, 0),
	}
	return data
}

func convSysRouterListByPid(routerList []*model.SysRouter, pid int) []*RouterModel {
	routerModelList := make([]*RouterModel, 0, len(routerList))
	for _, routerInfo := range routerList {
		if int(routerInfo.Pid) != pid {
			continue
		}

		routerModel := ConvSysRouter2VO(routerInfo)
		routerModel.Children = convSysRouterListByPid(routerList, int(routerModel.ID))
		routerModelList = append(routerModelList, routerModel)
	}
	if len(routerModelList) == 0 {
		return nil
	}
	return routerModelList
}
