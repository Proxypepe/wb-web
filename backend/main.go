package main

import (
	"fmt"
	"github.com/Proxypepe/wb-web/backend/cache"
	pg "github.com/Proxypepe/wb-web/backend/db"
	"github.com/Proxypepe/wb-web/backend/http"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	addr := fmt.Sprintf("postgres://%s:%s@localhost:17200/%s?sslmode=disable", "alex", "postgres", "wb")
	repo, err := pg.NewPostgresRepository("postgres", addr)
	if err != nil {
		log.Printf(err.Error())
		log.Printf("Error creating postgres repository")
		return
	}

	pg.SetRepository(repo)

	red, err := cache.NewRedisStore(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err != nil {
		log.Printf(err.Error())
	}

	cache.SetCacheService(red)

	server := http.NewServer()
	server.Run()

	defer func() {
		err := pg.CloseConn()
		if err != nil {
			return
		}
	}()

}
