package models

type Image struct {
	ID            	int    	`gorm:"primaryKey"`
	ImageType		string	`gorm:"type:varchar(10)"`
	FatherID		int		`gorm:"type:int"`
	ImageURL		string	`gorm:"type:varchar(100)"`
}
