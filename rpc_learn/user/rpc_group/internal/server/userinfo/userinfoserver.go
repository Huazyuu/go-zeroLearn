// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: user.proto

package server

import (
	"context"

	"gozero_learn/rpc_learn/user/rpc_group/internal/logic/userinfo"
	"gozero_learn/rpc_learn/user/rpc_group/internal/svc"
	"gozero_learn/rpc_learn/user/rpc_group/types/user"
)

type UserInfoServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserInfoServer
}

func NewUserInfoServer(svcCtx *svc.ServiceContext) *UserInfoServer {
	return &UserInfoServer{
		svcCtx: svcCtx,
	}
}

func (s *UserInfoServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := userinfologic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}
