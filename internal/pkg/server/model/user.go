package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NetID string `gorm:"unique"`
	Name  string
	Email string
}
