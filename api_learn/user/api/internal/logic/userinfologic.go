package logic

import (
	"context"

	"gozero_learn/api_learn/user/api/internal/svc"
	"gozero_learn/api_learn/user/api/internal/types"

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
		Code: 0,
		Data: types.UserInfo{
			Id:       1,
			UserName: "yu",
			Addr:     "earth",
		},
		Msg: "成功",
	}, err
}
