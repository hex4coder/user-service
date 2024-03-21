package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hex4coder/user-service/config"
	"github.com/hex4coder/user-service/database"
	"github.com/hex4coder/user-service/pkg/router"
)

var app *config.AppConfig

func init() {
	// set gin in release mode
	gin.SetMode(gin.ReleaseMode)

	// get app config
	app = config.LoadEnvToConfig()

	// connect to database
	database.ConnectToDB(app)

	// migrate user model
	database.MigrateModels()
}

func main() {

	// run
	router.Run(app)
}
