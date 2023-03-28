package main

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
	GinHost          string
	GinPort          string
}
