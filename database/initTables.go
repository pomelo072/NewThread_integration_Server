package database

import "integration_NewThread/models"

func Createtable() {
	// 数据库自动迁移
	Db.AutoMigrate(
		&models.UserInfo{},
		&models.AwardInfo{},
		&models.AwardDetail{},
		&models.IntegrationDetail{},
		&models.IntegrationApply{},

	)
}
