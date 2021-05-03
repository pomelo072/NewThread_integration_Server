package models

import "gorm.io/gorm"

// AwardInfo 奖品表结构
type AwardInfo struct {
	gorm.Model
	AwardType         string `gorm:"type:varchar(10)"`
	AwardName         string `gorm:"type:varchar(50)"`
	AwardIntroduction string `gorm:"type:varchar(100)"`
	NeedIntegration   int    `gorm:"type:int"`
	PeriodNumber	  int	 `gorm:"type:int"`
	InStock           int    `gorm:"type:int"`
	UsedNumber        int    `gorm:"type:int"`
}
