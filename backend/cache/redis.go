package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Proxypepe/wb-web/backend/schemas"
	"github.com/go-redis/redis"
	"log"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisStore(config *redis.Options) (*RedisCache, error) {
	client := redis.NewClient(config)
	pong, err := client.Ping().Result()
	if pong != "PONG" && err != nil {
		return nil, errors.New("unreachable redis connection")
	}
	return &RedisCache{
		client: client,
	}, nil
}

func (cache *RedisCache) SaveOrder(order schemas.Order) error {
	orderB, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = cache.client.Set(order.OrderUID, orderB, 0).Result()
	return err
}

func (cache *RedisCache) SaveOrders(orders []schemas.Order) error {
	for _, order := range orders {
		orderB, err := json.Marshal(order)
		if err != nil {
			log.Printf(err.Error())
			return err
		}
		_, err = cache.client.Set(order.OrderUID, orderB, 0).Result()
		if err != nil {
			log.Printf(err.Error())
			return err
		}
	}
	return nil
}

func (cache *RedisCache) GetOrder(uid string) (*schemas.Order, error) {
	var order schemas.Order
	val := cache.client.Get(uid)
	err := val.Err()
	if err == redis.Nil {
		fmt.Printf("Error")
		return nil, err
	}
	b, err := val.Bytes()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = json.Unmarshal(b, &order)
	return &order, err
}
