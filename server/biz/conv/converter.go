package conv

import (
	"github.com/moumou/server/biz/model"
	userdata "github.com/moumou/server/biz/service/user/data"
	api "github.com/moumou/server/gen/proto"
)

// goverter:converter
// goverter:extend IntToInt32
// goverter:output:file ../../gen/conv/generated.go
// goverter:ignoreUnexported
// goverter:ignoreMissing
type IConverter interface {
	ConvertPermissionListToVO(source []*model.Permission) []*api.Permission
	// goverter:autoMap BaseModel
	ConvertPermissionToVO(source *model.Permission) *api.Permission

	ConvertRoleListToVO(source []*model.Role) []*api.Role
	// goverter:autoMap BaseModel
	ConvertRoleToVO(source *model.Role) *api.Role
	// goverter:update target
	ConvertCreateRoleRequestDataToBO(source *api.CreateRoleRequestData, target *model.Role)

	ConvertUserListToVO(source []*model.User) []*api.User
	// goverter:autoMap BaseModel
	ConvertUserToVO(source *model.User) *api.User
	ConvertGetUserListRequestFilter(source *api.GetUserListRequestFilter) *userdata.ListUserFilter
}

func IntToInt32(value int) int32 {
	return int32(value)
}
