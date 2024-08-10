package handler

import (
	"github.com/moumou/server/biz/model"
	api "github.com/moumou/server/gen/proto"
)

func ConvPermissionList2VO(permissionModelList []*model.Permission) []*api.Permission {
	list := make([]*api.Permission, len(permissionModelList))
	for i, permissionModel := range permissionModelList {
		list[i] = ConvPermission2VO(permissionModel)
	}
	return list
}

func ConvPermission2VO(permissionModel *model.Permission) *api.Permission {
	return &api.Permission{
		Id:   permissionModel.ID,
		Name: permissionModel.Name,
		Path: permissionModel.Path,
		Sort: int32(permissionModel.Sort),
	}
}
