// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.27.1
// source: user.proto

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

const OperationUserHandlerGetUserInfo = "/moumou.server.api.UserHandler/GetUserInfo"
const OperationUserHandlerGetUserList = "/moumou.server.api.UserHandler/GetUserList"

type UserHandlerHTTPServer interface {
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
}

func RegisterUserHandlerHTTPServer(s *http.Server, srv UserHandlerHTTPServer) {
	r := s.Route("/")
	r.POST("/user/list", _UserHandler_GetUserList0_HTTP_Handler(srv))
	r.POST("/user/info", _UserHandler_GetUserInfo0_HTTP_Handler(srv))
}

func _UserHandler_GetUserList0_HTTP_Handler(srv UserHandlerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserHandlerGetUserList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserList(ctx, req.(*GetUserListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserListResponse)
		return ctx.Result(200, reply)
	}
}

func _UserHandler_GetUserInfo0_HTTP_Handler(srv UserHandlerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserHandlerGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoResponse)
		return ctx.Result(200, reply)
	}
}

type UserHandlerHTTPClient interface {
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest, opts ...http.CallOption) (rsp *GetUserInfoResponse, err error)
	GetUserList(ctx context.Context, req *GetUserListRequest, opts ...http.CallOption) (rsp *GetUserListResponse, err error)
}

type UserHandlerHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHandlerHTTPClient(client *http.Client) UserHandlerHTTPClient {
	return &UserHandlerHTTPClientImpl{client}
}

func (c *UserHandlerHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...http.CallOption) (*GetUserInfoResponse, error) {
	var out GetUserInfoResponse
	pattern := "/user/info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserHandlerGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserHandlerHTTPClientImpl) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...http.CallOption) (*GetUserListResponse, error) {
	var out GetUserListResponse
	pattern := "/user/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserHandlerGetUserList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
