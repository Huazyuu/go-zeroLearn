package logic

import (
	"context"
	"errors"
	"gozero_learn/model_learn/user_gorm/api/internal/svc"
	"gozero_learn/model_learn/user_gorm/api/internal/types"
	"gozero_learn/model_learn/user_gorm/model"

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
	user := model.UserModel{
		Username: req.UserName,
		Password: req.Password,
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		return "", errors.New("登录失败")
	}
	return user.Username, nil
}
