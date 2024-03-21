package database

import (
	"log"

	"github.com/hex4coder/user-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(app *config.AppConfig) {

	// get dsn/connection string to database
	dsn := app.Database.GetDBUrl()

	var err error

	// open connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// check error
	if err != nil {
		log.Fatalln("Could'nt connect to database", err.Error())
	}
}
