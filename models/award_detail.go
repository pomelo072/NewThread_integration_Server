package models

import (
	"gorm.io/gorm"
)

type AwardDetail struct {
	gorm.Model
	AwardUID        string 		`gorm:"type:varchar(12)"`
	AwardID         int    		`gorm:"type:int"`
	AwardName       string 		`gorm:"type:varchar(50)"`
	NeedIntegration int    		`gorm:"type:int"`
	ApplyPeriod		int			`gorm:"type:int"`
	AwardStatus     int    		`gorm:"type:int"`
}