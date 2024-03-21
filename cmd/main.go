package main

import (
	"github.com/hex4coder/user-service/config"
	"github.com/hex4coder/user-service/database"
)

var app *config.AppConfig

func init() {
	// get app config
	app = config.LoadEnvToConfig()

	// connect to database
	database.ConnectToDB(app)

	// migrate user model
	database.MigrateModels()
}

func main() {

	// run
	app.Run()
}
