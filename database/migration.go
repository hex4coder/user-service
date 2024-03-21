package database

import "github.com/hex4coder/user-service/pkg/models"

func MigrateModels() {
	DB.AutoMigrate(&models.User{})
}
