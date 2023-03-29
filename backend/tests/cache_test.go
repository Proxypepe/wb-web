package tests

import (
	"fmt"
	"github.com/Proxypepe/wb-web/backend/cache"
	conf "github.com/Proxypepe/wb-web/backend/config"
	"github.com/Proxypepe/wb-web/backend/schemas"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	log.Println("setup suite")
	config := conf.NewTestConfig()
	red, err := cache.NewRedisStore(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.TestRedisHost, config.TestRedisPort),
		Password: config.TestRedisPassword,
		DB:       config.TestRedisDB,
	})
	if err != nil {
		log.Println(err)
	}
	cache.SetCacheService(red)

	return func(tb testing.TB) {
		log.Println("teardown suite")
	}
}

func TestConnectRedis(t *testing.T) {
	config := conf.NewTestConfig()
	red, err := cache.NewRedisStore(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.TestRedisHost, config.TestRedisPort),
		Password: config.TestRedisPassword,
		DB:       config.TestRedisDB,
	})

	if err != nil {
		t.Error(err)
	}
	assert.NotEqual(t, red, nil)
}

func TestFaultConnectRedis(t *testing.T) {
	config := conf.NewTestConfig()
	_, err := cache.NewRedisStore(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "1", config.TestRedisPort),
		Password: config.TestRedisPassword,
		DB:       config.TestRedisDB,
	})

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "unreachable redis connection")
}

func TestSaveOrder(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	order := GetExampleOrder("b563feb7b2b84b6test")
	err := cache.SaveOrder(order)
	if err != nil {
		t.Error(err)
		return
	}
	getOrder, err := cache.GetOrder(order.OrderUID)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, getOrder, &order)
}

func TestSaveOrders(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	orders := []schemas.Order{
		GetExampleOrder("b563feb7b2b84b6test"),
		GetExampleOrder("b563feb7b2b84b6test2"),
	}
	err := cache.SaveOrders(orders)
	if err != nil {
		t.Error(err)
		return
	}
	for _, order := range orders {
		getOrder, err := cache.GetOrder(order.OrderUID)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, getOrder, &order)
	}
}

func TestGetNotExistsOrder(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	_, err := cache.GetOrder("")

	assert.Equal(t, err, redis.Nil)
}
