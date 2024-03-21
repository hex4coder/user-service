package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Address  string `json:"address"`
}

//type NotFound error

// func CreateUser() error
// func findUser(userID uint) (*User, error)
// func UpdateUser(userID uint, updatedUser *User) error
// func DeleteUser(userID uint) error
