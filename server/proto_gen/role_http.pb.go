// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.27.1
// source: role.proto

package api

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRoleHandlerCreateRole = "/server.api.RoleHandler/CreateRole"
const OperationRoleHandlerDeleteRole = "/server.api.RoleHandler/DeleteRole"
const OperationRoleHandlerGetRoleList = "/server.api.RoleHandler/GetRoleList"
const OperationRoleHandlerUpdateRole = "/server.api.RoleHandler/UpdateRole"

type RoleHandlerHTTPServer interface {
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleResponse, error)
	GetRoleList(context.Context, *GetRoleListRequest) (*GetRoleListResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
}

func RegisterRoleHandlerHTTPServer(s *http.Server, srv RoleHandlerHTTPServer) {
	r := s.Route("/")
	r.POST("/role/list", _RoleHandler_GetRoleList0_HTTP_Handler(srv))
	r.POST("/role/create", _RoleHandler_CreateRole0_HTTP_Handler(srv))
	r.POST("/role/update", _RoleHandler_UpdateRole0_HTTP_Handler(srv))
	r.POST("/role/delete", _RoleHandler_DeleteRole0_HTTP_Handler(srv))
}

func _RoleHandler_GetRoleList0_HTTP_Handler(srv RoleHandlerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRoleListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleHandlerGetRoleList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRoleList(ctx, req.(*GetRoleListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRoleListResponse)
		return ctx.Result(200, reply)
	}
}

func _RoleHandler_CreateRole0_HTTP_Handler(srv RoleHandlerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleHandlerCreateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRole(ctx, req.(*CreateRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateRoleResponse)
		return ctx.Result(200, reply)
	}
}

func _RoleHandler_UpdateRole0_HTTP_Handler(srv RoleHandlerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleHandlerUpdateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateRole(ctx, req.(*UpdateRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateRoleResponse)
		return ctx.Result(200, reply)
	}
}

func _RoleHandler_DeleteRole0_HTTP_Handler(srv RoleHandlerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRoleHandlerDeleteRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRole(ctx, req.(*DeleteRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteRoleResponse)
		return ctx.Result(200, reply)
	}
}

type RoleHandlerHTTPClient interface {
	CreateRole(ctx context.Context, req *CreateRoleRequest, opts ...http.CallOption) (rsp *CreateRoleResponse, err error)
	DeleteRole(ctx context.Context, req *DeleteRoleRequest, opts ...http.CallOption) (rsp *DeleteRoleResponse, err error)
	GetRoleList(ctx context.Context, req *GetRoleListRequest, opts ...http.CallOption) (rsp *GetRoleListResponse, err error)
	UpdateRole(ctx context.Context, req *UpdateRoleRequest, opts ...http.CallOption) (rsp *UpdateRoleResponse, err error)
}

type RoleHandlerHTTPClientImpl struct {
	cc *http.Client
}

func NewRoleHandlerHTTPClient(client *http.Client) RoleHandlerHTTPClient {
	return &RoleHandlerHTTPClientImpl{client}
}

func (c *RoleHandlerHTTPClientImpl) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...http.CallOption) (*CreateRoleResponse, error) {
	var out CreateRoleResponse
	pattern := "/role/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleHandlerCreateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHandlerHTTPClientImpl) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...http.CallOption) (*DeleteRoleResponse, error) {
	var out DeleteRoleResponse
	pattern := "/role/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleHandlerDeleteRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHandlerHTTPClientImpl) GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...http.CallOption) (*GetRoleListResponse, error) {
	var out GetRoleListResponse
	pattern := "/role/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleHandlerGetRoleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RoleHandlerHTTPClientImpl) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...http.CallOption) (*UpdateRoleResponse, error) {
	var out UpdateRoleResponse
	pattern := "/role/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRoleHandlerUpdateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}