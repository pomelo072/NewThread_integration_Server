package models

import (
	"gorm.io/gorm"
)

type IntegrationApply struct {
	gorm.Model
	IntegrationType int    `gorm:"type:int"`
	ApplyUID        string `gorm:"type:varchar(12)"`
	ApplyText       string `gorm:"type:varchar(50)"`
	ApplyType       string `gorm:"type:varchar(10)"`
	ApplyLevel      string `gorm:"type:varchar(10)"`
	ContestLevel    string `gorm:"type:varchar(10)"`
	ApplyNumber     int    `gorm:"int"`
	ResponseReason  string `gorm:"type:varchar(50)"`
	ApplyStatus     int    `gorm:"type:int"`
}
