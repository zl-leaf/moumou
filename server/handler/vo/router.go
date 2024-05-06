package vo

import "github.com/moumou/server/biz/model"

func ConvRouterList2RouterTreeData(routerList []*model.Router) *RouterTreeResponseData {
	data := &RouterTreeResponseData{
		Routers: convRouterListByPid(routerList, 0),
	}
	return data
}

func convRouterListByPid(routerList []*model.Router, pid int) []*RouterRecord {
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

func convRouter(routerModel *model.Router) *RouterRecord {
	return &RouterRecord{
		Name:   routerModel.Name,
		Title:  routerModel.Title,
		Path:   routerModel.Path,
		IsMenu: routerModel.IsMenu,
	}
}
