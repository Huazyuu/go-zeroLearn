package logic

import (
	"context"
	"encoding/json"

	"gozero_learn/api_learn/user/apiv4_jwt/internal/svc"
	"gozero_learn/api_learn/user/apiv4_jwt/internal/types"

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
	user_id := l.ctx.Value("user_id").(json.Number)
	username := l.ctx.Value("username").(string)
	uid, _ := user_id.Int64()
	return &types.UserInfoResponse{
		Id:       uint(uid),
		UserName: username,
		Addr:     "",
	}, nil
}
