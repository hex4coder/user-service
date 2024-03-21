package config

import "fmt"

type DatabaseConfiguration struct {
	Port     uint
	Host     string
	User     string
	Password string
	DBName   string
}

func NewDatabaseConfig() *DatabaseConfiguration {
	return &DatabaseConfiguration{}
}

func (dbc *DatabaseConfiguration) GetDBUrl() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbc.Host, dbc.User, dbc.Password, dbc.DBName, dbc.Port)

	return dsn
}
