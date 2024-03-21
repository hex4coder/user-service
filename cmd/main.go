package main

import (
	"github.com/hex4coder/user-service/config"
)

func main() {
	// get app config
	app := config.LoadEnvToConfig()

	// run
	app.Run()
}
