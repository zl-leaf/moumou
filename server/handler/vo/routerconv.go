package vo

import "github.com/moumou/server/biz/model"

func ConvRouterList2RouterTreeData(routerList []*model.SysRouter) *RouterTreeResponseData {
	data := &RouterTreeResponseData{
		Routers: convRouterListByPid(routerList, 0),
	}
	return data
}

func convRouterListByPid(routerList []*model.SysRouter, pid int) []*RouterRecord {
	routerRecordList := make([]*RouterRecord, 0, len(routerList))
	for _, routerModel := range routerList {
		if int(routerModel.Pid) != pid {
			continue
		}

		routerRecord := convRouter(routerModel)
		routerRecord.Children = convRouterListByPid(routerList, int(routerModel.ID))
		routerRecordList = append(routerRecordList, routerRecord)
	}
	return routerRecordList
}

func convRouter(routerModel *model.SysRouter) *RouterRecord {
	return &RouterRecord{
		Name:   routerModel.Name,
		Title:  routerModel.Title,
		Path:   routerModel.Path,
		IsMenu: routerModel.IsMenu,
	}
}
