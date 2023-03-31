package cache

import (
	"encoding/json"
	"errors"
	"github.com/Proxypepe/wb-web/backend/schemas"
	"github.com/go-redis/redis"
	"log"
)

// RedisCache Redis based cache implementation
type RedisCache struct {
	client *redis.Client
}

// NewRedisStore returns a structure for working with the cache
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
		log.Print(err)
		return err
	}
	_, err = cache.client.Set(order.OrderUID, orderB, 0).Result()
	return err
}

func (cache *RedisCache) SaveOrders(orders []schemas.Order) error {
	for _, order := range orders {
		orderB, err := json.Marshal(order)
		if err != nil {
			log.Print(err.Error())
			return err
		}
		_, err = cache.client.Set(order.OrderUID, orderB, 0).Result()
		if err != nil {
			log.Print(err.Error())
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
		return nil, err
	}
	b, err := val.Bytes()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	err = json.Unmarshal(b, &order)
	return &order, err
}
