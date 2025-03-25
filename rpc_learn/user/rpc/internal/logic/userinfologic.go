package logic

import (
	"context"

	"gozero_learn/rpc_learn/user/rpc/internal/svc"
	"gozero_learn/rpc_learn/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {

	return &user.UserInfoResponse{
		UserId:   in.UserId,
		Username: "yuyuyuhahah",
	}, nil
}
