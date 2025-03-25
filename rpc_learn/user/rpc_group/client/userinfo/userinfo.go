// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: user.proto

package userinfo

import (
	"context"

	"gozero_learn/rpc_learn/user/rpc_group/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserCreateRequest  = user.UserCreateRequest
	UserCreateResponse = user.UserCreateResponse
	UserInfoRequest    = user.UserInfoRequest
	UserInfoResponse   = user.UserInfoResponse

	UserInfo interface {
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	}

	defaultUserInfo struct {
		cli zrpc.Client
	}
)

func NewUserInfo(cli zrpc.Client) UserInfo {
	return &defaultUserInfo{
		cli: cli,
	}
}

func (m *defaultUserInfo) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUserInfoClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
