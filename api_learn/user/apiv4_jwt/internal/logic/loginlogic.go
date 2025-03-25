package logic

import (
	"context"
	"gozero_learn/api_learn/user/common/jwt"

	"gozero_learn/api_learn/user/apiv4_jwt/internal/svc"
	"gozero_learn/api_learn/user/apiv4_jwt/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	c := l.svcCtx.Config.Auth
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   1,
		Username: "1",
		Role:     1,
	}, c.AccessSecret, c.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, nil
}
