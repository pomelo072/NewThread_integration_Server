package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"integration_server/config"
	"integration_server/models"
	"integration_server/utils"
	"strings"
	"time"
)

var Db *gorm.DB

func init() {
	// 合成数据库连接
	dsn := strings.Join([]string{config.Sysconfig.DBUserName, ":", config.Sysconfig.DBPassword, "@(", config.Sysconfig.DBIp, ":", config.Sysconfig.DBPort, ")/", config.Sysconfig.DBName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	var err error
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 数据库设置
	sqlDB, err := Db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(2000)
	sqlDB.SetConnMaxLifetime(1 * time.Second)
	// 生成数据表
	Createtable()
	// 创建管理账号
	admin := new(models.UserInfo)
	Db.FirstOrCreate(&admin, models.UserInfo{
		UID:              config.Sysconfig.Admin,
		UserName:         "New_Thread",
		UserGroup:        "New_Thread",
		UserVerification: utils.GetSHAEncode(utils.GetSHAEncode(config.Sysconfig.AdminVerify)),
		AUTH:             0,
	})
}
