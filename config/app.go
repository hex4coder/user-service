package config

import (
	"fmt"
	"log"
	"net/http"
)

type AppConfig struct {
	Database      *DatabaseConfiguration
	BackendServer *ServerConfiguration
	Router        http.Handler
}

func NewAppConfig(db *DatabaseConfiguration, sb *ServerConfiguration, router http.Handler) *AppConfig {
	return &AppConfig{
		Database:      db,
		BackendServer: sb,
		Router:        router,
	}
}

func (app *AppConfig) Run() {
	fmt.Printf("[User Service] Running server on endpoint %s\n", app.BackendServer.GetEndpoint())
	log.Fatal(http.ListenAndServe(app.BackendServer.GetEndpoint(), app.Router))
}
