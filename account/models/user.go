package models

import "gorm.io/gorm"

type User struct {
	Name     string
	Email    string
	Password string `gorm:"not null"`
	gorm.Model
}
