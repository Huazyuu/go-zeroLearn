package svc

import (
	"gorm.io/gorm"
	"gozero_learn/model_learn/common"
	"gozero_learn/model_learn/user_gorm/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     common.InitGorm(c.Mysql.DataSource),
	}

}
