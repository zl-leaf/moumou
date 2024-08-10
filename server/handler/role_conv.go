package handler

import (
	"github.com/moumou/server/biz/model"
	api "github.com/moumou/server/gen/proto"
)

func ConvRoleList2VO(roleModelList []*model.Role) []*api.Role {
	ret := make([]*api.Role, len(roleModelList))
	for i, roleModel := range roleModelList {
		ret[i] = ConvRole2VO(roleModel)
	}
	return ret
}

func ConvRole2VO(roleModel *model.Role) *api.Role {
	return &api.Role{
		Id:   roleModel.ID,
		Name: roleModel.Name,
	}
}

func ConvCreateRequestData2Role(roleVO *api.CreateRoleRequestData) *model.Role {
	return &model.Role{
		Name: roleVO.Name,
	}
}

func ConvUpdateRequestData2Role(roleVO *api.UpdateRoleRequestData) *model.Role {
	return &model.Role{
		BaseModel: model.BaseModel{ID: roleVO.Id},
		Name:      roleVO.Name,
	}
}
