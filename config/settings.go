package config

type Settings struct {
	Environment string `json:"ENVIRONMENT" env:"ENVIRONMENT"`

	ServiceName string `json:"SERVICE_NAME" env:"SERVICE_NAME"`
	ServiceRegisteAddress string `json:"SERVICE_REGISTER_ADDRESS" env:"SERVICE_REGISTER_ADDRESS"`
	ServiceRegisterPort    string `json:"SERVICE_REGISTER_PORT" env:"SERVICE_REGISTER_PORT"`

	DbUseInMem string `json:"DB_USE_INMEM" env:"DB_USE_INMEM"`
	DbType     string `json:"DB_TYPE" env:"DB_TYPE"`
	DbHost     string `json:"DB_HOST" env:"DB_HOST"`
	DbDataBase string `json:"DB_DATABASE" env:"DB_DATABASE"`
	DbUsername string `json:"DB_USERNAME" env:"DB_USERNAME"`
	DbPassword string `json:"DB_PASSWORD" env:"DB_PASSWORD"`
	DbPort     string `json:"DB_PORT" env:"DB_PORT"`
	DbSchema   string `json:"DB_SCHEMA" env:"DB_SCHEMA"`
	DbSSLMode  string `json:"DB_SSL_MODE" env:"DB_SSL_MODE"`
}