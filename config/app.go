package config

type AppConfig struct {
	Database      *DatabaseConfiguration
	BackendServer *ServerConfiguration
}

func NewAppConfig(db *DatabaseConfiguration, sb *ServerConfiguration) *AppConfig {
	return &AppConfig{
		Database:      db,
		BackendServer: sb,
	}
}
