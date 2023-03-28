package config

import "os"

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDb       string
	PostgresHost     string
	PostgresPort     string
	RedisHost        string
	RedisPort        string
	RedisDB          int
	RedisPassword    string
	NatsClusterID    string
	NatsSubject      string
	ServerHost       string
	ServerPort       string
}

func NewConfig() *Config {
	return &Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDb:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
		RedisDB:          0,
		RedisPassword:    "",
		NatsClusterID:    "",
		NatsSubject:      "",
		ServerHost:       os.Getenv("SERVER_HOST"),
		ServerPort:       os.Getenv("SERVER_PORT"),
	}
}

func (config *Config) validate() bool {

	return true
}
