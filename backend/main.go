package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Proxypepe/wb-web/backend/cache"
	conf "github.com/Proxypepe/wb-web/backend/config"
	"github.com/Proxypepe/wb-web/backend/db"
	"github.com/Proxypepe/wb-web/backend/http"
	"github.com/Proxypepe/wb-web/backend/schemas"
	"github.com/go-redis/redis"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	config := conf.NewConfig()
	sc, err := stan.Connect(config.NatsClusterID, config.NatsClientID, stan.NatsURL(
		fmt.Sprintf("nats://%s:%s", config.NatsHost, config.NatsPort)))
	if err != nil {
		log.Print(err.Error())
		return
	}
	sub, err := sc.Subscribe(config.NatsSubject, func(m *stan.Msg) {
		var newOrder schemas.Order
		erro := json.Unmarshal(m.Data, &newOrder)
		if erro != nil {
			log.Printf("Error unmarshalling: %s\n", string(m.Data))
			return
		}
		erro = db.InsertOrder(context.Background(), newOrder)
		if erro != nil {
			log.Printf("Error while inserting order: %s\n", string(m.Data))
			return
		}
	})
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresDb,
	)
	if err != nil {
		log.Print("Error")
	}
	defer func(sub stan.Subscription) {
		err := sub.Unsubscribe()
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}(sub)
	defer func(sub stan.Subscription) {
		err := sub.Close()
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}(sub)
	repo, err := db.NewPostgresRepository("postgres", addr)
	if err != nil {
		log.Print(addr)
		log.Print(err.Error())
		log.Print("Error creating postgres repository")
		return
	}

	db.SetRepository(repo)

	red, err := cache.NewRedisStore(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	if err != nil {
		log.Print(err.Error())
	}

	cache.SetCacheService(red)

	serverAddr := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)

	server := http.NewServer()
	server.Run(serverAddr)

	defer func() {
		err := db.CloseConn()
		if err != nil {
			return
		}
	}()

}
