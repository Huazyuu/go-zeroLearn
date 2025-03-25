package logic

import (
	"context"
	"fmt"

	"gozero_learn/rpc_learn/user/rpc/internal/svc"
	"gozero_learn/rpc_learn/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (*user.UserCreateResponse, error) {

	fmt.Println(in.Username, in.Password)

	return &user.UserCreateResponse{}, nil
}
