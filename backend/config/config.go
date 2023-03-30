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
	NatsClientID     string
	NatsHost         string
	NatsPort         string
	NatsSubject      string
	ServerHost       string
	ServerPort       string
}

type TestConfig struct {
	c                    *Config
	TestPostgresUser     string
	TestPostgresPassword string
	TestPostgresDb       string
	TestPostgresHost     string
	TestPostgresPort     string
	TestRedisHost        string
	TestRedisPort        string
	TestRedisDB          int
	TestRedisPassword    string
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
		NatsClusterID:    os.Getenv("NATS_CLUSTER_ID"),
		NatsClientID:     os.Getenv("NATS_CLIENT_ID"),
		NatsHost:         os.Getenv("NATS_HOST"),
		NatsPort:         os.Getenv("NATS_PORT"),
		NatsSubject:      os.Getenv("NATS_SUBJECT"),
		ServerHost:       os.Getenv("SERVER_HOST"),
		ServerPort:       os.Getenv("SERVER_PORT"),
	}
}

func NewTestConfig() *TestConfig {
	return &TestConfig{
		c:                    NewConfig(),
		TestPostgresUser:     "alex",
		TestPostgresPassword: "postgres",
		TestPostgresDb:       "test_wb",
		TestPostgresHost:     "localhost",
		TestPostgresPort:     "17200",
		TestRedisHost:        "localhost",
		TestRedisPort:        "6379",
		TestRedisDB:          0,
		TestRedisPassword:    "",
	}
}
