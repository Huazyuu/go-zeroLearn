package logic

import (
	"context"

	"gozero_learn/api_learn/user/apiv1/internal/svc"
	"gozero_learn/api_learn/user/apiv1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {

	return &types.UserInfoResponse{
		Id:       1,
		UserName: "1",
		Addr:     "1",
	}, nil
}
