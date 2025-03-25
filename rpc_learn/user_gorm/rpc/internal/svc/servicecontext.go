package svc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	models "gozero_learn/rpc_learn/user_gorm/model"
	"gozero_learn/rpc_learn/user_gorm/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initGorm(c.Mysql.DataSource)
	db.AutoMigrate(&models.UserModel{})
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}

// InitGorm gorm初始化
func initGorm(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	return db
}
