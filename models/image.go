package models

type Image struct {
	ID        uint   `gorm:"primaryKey"`
	ImageType string `gorm:"type:varchar(10)"`
	FatherID  uint   `gorm:"type:int"`
	ImageURL  string `gorm:"type:varchar(100)"`
}
