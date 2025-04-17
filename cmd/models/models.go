package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          uint
	Name        string
	Email       *string
	ProfileCode *int
	Profile     Profile `gorm:"foreignKey:ProfileCode"`
}

type Profile struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
}
