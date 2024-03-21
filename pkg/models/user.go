package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required, email"`
	Password string `json:"-" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Address  string `json:"address"`
}
