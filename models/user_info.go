package models

// UserInfo 用户数据表结构
type UserInfo struct {
	UID                string              `gorm:"primaryKey;type:varchar(12);index"`
	UserName           string              `gorm:"type:varchar(10)"`
	UserGroup          string              `gorm:"type:varchar(10)"`
	UserIntegration    int                 `gorm:"type:int"`
	TeamIntegration    int                 `gorm:"type:int"`
	UserVerification   string              `gorm:"type:varchar(70)"`
	AUTH               int                 `gorm:"type:int"`
}