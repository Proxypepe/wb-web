package main

import (
	"fmt"
	"github.com/Proxypepe/wb-web/backend/cache"
	conf "github.com/Proxypepe/wb-web/backend/config"
	pg "github.com/Proxypepe/wb-web/backend/db"
	"github.com/Proxypepe/wb-web/backend/http"
	"github.com/go-redis/redis"
	"log"
)

func main() {

	config := conf.NewConfig()

	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDb,
	)
	repo, err := pg.NewPostgresRepository("postgres", addr)
	if err != nil {
		log.Print(addr)
		log.Printf(err.Error())
		log.Printf("Error creating postgres repository")
		return
	}

	pg.SetRepository(repo)

	red, err := cache.NewRedisStore(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	if err != nil {
		log.Printf(err.Error())
	}

	cache.SetCacheService(red)

	serverAddr := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)

	server := http.NewServer()
	server.Run(serverAddr)

	defer func() {
		err := pg.CloseConn()
		if err != nil {
			return
		}
	}()

}
