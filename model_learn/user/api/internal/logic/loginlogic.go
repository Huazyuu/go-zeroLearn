package logic

import (
	"context"
	"fmt"
	"gozero_learn/model_learn/user/model"

	"gozero_learn/model_learn/user/api/internal/svc"
	"gozero_learn/model_learn/user/api/internal/types"

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
	// 增
	l.svcCtx.UsersModel.Insert(context.Background(), &model.User{
		Username: "测试插入name2",
		Password: "123456",
	})
	fmt.Println("============================")
	// 查
	user, err := l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	fmt.Println(user, err)
	fmt.Println("============================")

	// 查
	user, err = l.svcCtx.UsersModel.FindOneByUsername(context.Background(), "测试插入name")
	fmt.Println(user, err)
	fmt.Println("============================")

	// 改
	l.svcCtx.UsersModel.Update(context.Background(), &model.User{
		Username: "测试插入hhhha",
		Password: "1234567",
		Id:       1,
	})
	user, err = l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	fmt.Println(user, err)
	fmt.Println("============================")

	user, err = l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	fmt.Println(user, err)
	// 删
	l.svcCtx.UsersModel.Delete(context.Background(), 1)
	user, err = l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	l.svcCtx.UsersModel.Delete(context.Background(), 2)
	user, err = l.svcCtx.UsersModel.FindOne(context.Background(), 2)
	fmt.Println(user, err)
	return
}
