package conv

import (
	"github.com/moumou/server/biz/model"
	userdata "github.com/moumou/server/biz/service/user/data"
	api "github.com/moumou/server/gen/proto"
)

// goverter:converter
// goverter:extend IntToInt32
// goverter:extend Int32ToInt
// goverter:output:file ../../gen/conv/generated.go
// goverter:ignoreUnexported
// goverter:ignoreMissing
type IConverter interface {
	ConvertPermissionTreeNodeListToVO(source []*model.Permission) []*api.PermissionTreeNode
	// goverter:autoMap BaseModel
	ConvertPermissionTreeNodeToVO(source *model.Permission) *api.PermissionTreeNode
	ConvertPermissionListToVO(source []*model.Permission) []*api.Permission
	// goverter:autoMap BaseModel
	ConvertPermissionToVO(source *model.Permission) *api.Permission
	// goverter:update target
	ConvertCreatePermissionRequestDataToBO(source *api.CreatePermissionRequestData, target *model.Permission)
	// goverter:update target
	ConvertUpdatePermissionRequestDataToBO(source *api.UpdatePermissionRequestData, target *model.Permission)

	ConvertRoleListToVO(source []*model.Role) []*api.Role
	// goverter:autoMap BaseModel
	ConvertRoleToVO(source *model.Role) *api.Role
	// goverter:update target
	ConvertCreateRoleRequestDataToBO(source *api.CreateRoleRequestData, target *model.Role)
	// goverter:update target
	ConvertUpdateRoleRequestDataToBO(source *api.UpdateRoleRequestData, target *model.Role)

	ConvertUserListToVO(source []*model.User) []*api.User
	// goverter:autoMap BaseModel
	ConvertUserToVO(source *model.User) *api.User
	ConvertGetUserListRequestFilter(source *api.GetUserListRequestFilter) *userdata.ListUserFilter
	// goverter:update target
	ConvertCreateUserRequestDataToBO(source *api.CreateUserRequestData, target *model.User)
	// goverter:update target
	ConvertUpdateUserRequestDataToBO(source *api.UpdateUserRequestData, target *model.User)
}

func IntToInt32(value int) int32 {
	return int32(value)
}

func Int32ToInt(value int32) int {
	return int(value)
}
