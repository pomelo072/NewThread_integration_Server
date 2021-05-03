package models

import "gorm.io/gorm"

type IntegrationDetail struct {
	gorm.Model
	IntegrationType         int    		`gorm:"type:int"`
	IntegrationUID  string 		`gorm:"type:varchar(12)"`
	OperationDetail string 		`gorm:"type:varchar(50)"`
	OperationType   string 		`gorm:"type:varchar(10)"`
	OperationNumber int    		`gorm:"type:int"`
	AfterOperation  int    		`gorm:"type:int"`
}