package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozero_learn/model_learn/user/api/internal/config"
	"gozero_learn/model_learn/user/model"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUserModel(mysqlConn),
	}

}
