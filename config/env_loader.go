package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnvToConfig() *AppConfig {

	// load env to memory
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(".env loaded")

	// get server config
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalln("Can't load server port from .env")
	}
	serverIP := os.Getenv("SERVER_IP")
	server := NewServerConfig(uint(serverPort), serverIP)

	// get database config
	db := NewDatabaseConfig()
	db.DBName = os.Getenv("POSTGRES_DB")
	db.Host = os.Getenv("POSTGRES_HOST")
	db.User = os.Getenv("POSTGRES_USER")
	db.Password = os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	idbPort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatalln("Can't load database port in .env")
	}
	db.Port = uint(idbPort)

	// build new app configuration
	app := NewAppConfig(db, server)

	// return it
	return app
}
