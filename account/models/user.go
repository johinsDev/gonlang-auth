package models

import "gorm.io/gorm"

type User struct {
	Name  string
	Email string
	gorm.Model
}
