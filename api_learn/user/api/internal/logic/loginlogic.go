package logic

import (
	"context"
	"fmt"

	"gozero_learn/api_learn/user/api/internal/svc"
	"gozero_learn/api_learn/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	fmt.Println(req.UserName, req.Password)
	return &types.LoginResponse{
		Code: 0,
		Data: "登录 token xxx.xxx.xxx",
		Msg:  "成功",
	}, nil
}
