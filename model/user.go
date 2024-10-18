package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	RealName string `gorm:"type:varchar(100)"`
	Avatar   string `gorm:"type:varchar(255)"`
	Mobile   string `gorm:"type:varchar(20)"`
	Email    string `gorm:"type:varchar(128);unique"`
	Password string `gorm:"type:varchar(255);not null"`
}
