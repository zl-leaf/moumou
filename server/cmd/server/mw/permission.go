package mw

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

func APIPermission(svc *service.Service) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if err := validAPIPermission(ctx, svc); err != nil {
				return nil, err
			}
			return handler(ctx, req)
		}
	}
}

func validAPIPermission(ctx context.Context, svc *service.Service) error {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil
	}
	operation := tr.Operation()
	needPermissionCodes, needCheck := getApiPermissionConfig()[operation]
	if !needCheck {
		return nil
	}

	userInfo, err := svc.UserService.Self(ctx)
	if err != nil {
		return err
	}

	permissions, err := svc.RoleService.GetPermissionCodesByUid(ctx, userInfo.Id)
	if err != nil {
		return err
	}
	permissionCodeMap := make(map[string]bool, len(permissions))
	for _, permission := range permissions {
		permissionCodeMap[permission.Code] = true
	}

	for _, needPermissionCode := range needPermissionCodes {
		if !permissionCodeMap[needPermissionCode] {
			return errors.Forbidden("Forbidden", "无权限")
		}
	}

	return nil
}

func getApiPermissionConfig() map[string][]string {
	return map[string][]string{
		api.OperationUserHandlerGetUserList: {"ManageUserRead"},
		api.OperationUserHandlerGetUserInfo: {"ManageUserRead"},
		api.OperationUserHandlerCreateUser:  {"ManageUserWrite"},
		api.OperationUserHandlerUpdateUser:  {"ManageUserWrite"},

		api.OperationRoleHandlerGetRoleList: {"ManageRoleRead"},
		api.OperationRoleHandlerGetRoleInfo: {"ManageRoleRead"},
		api.OperationRoleHandlerCreateRole:  {"ManageRoleWrite"},
		api.OperationRoleHandlerUpdateRole:  {"ManageRoleWrite"},

		api.OperationPermissionHandlerGetPermissionTree: {"ManagePermissionRead"},
		api.OperationPermissionHandlerGetPermissionInfo: {"ManagePermissionRead"},
		api.OperationPermissionHandlerCreatePermission:  {"ManagePermissionWrite"},
		api.OperationPermissionHandlerUpdatePermission:  {"ManagePermissionWrite"},
	}
}
