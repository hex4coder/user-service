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
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dbc.User, dbc.Password, dbc.Host, dbc.Port, dbc.DBName)
}
