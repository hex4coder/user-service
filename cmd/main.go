package main

import (
	"github.com/hex4coder/user-service/config"
	"github.com/hex4coder/user-service/pkg/router"
)

func main() {
	api := router.SetupUserAPI()

	server := config.NewServerConfig(9000, "127.0.0.1")

	db := config.NewDatabaseConfig()

	db.DBName = "userdb-kunix"
	db.Host = "localhost"
	db.Port = 5555
	db.User = "kunix"
	db.Password = "kunixpwd"

	app := config.NewAppConfig(db, server, api)
	app.Run()
}
