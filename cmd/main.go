package main

import (
	"github.com/hex4coder/user-service/config"
)

func main() {

	app := config.LoadEnvToConfig()

	app.Run()
}
