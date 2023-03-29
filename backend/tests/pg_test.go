package tests

import (
	"context"
	"fmt"
	conf "github.com/Proxypepe/wb-web/backend/config"
	pg "github.com/Proxypepe/wb-web/backend/db"
	"github.com/Proxypepe/wb-web/backend/schemas"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func setupSuiteDB(tb testing.TB) func(tb testing.TB) {
	config := conf.NewTestConfig()
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.TestPostgresUser,
		config.TestPostgresPassword,
		config.TestPostgresHost,
		config.TestPostgresPort,
		config.TestPostgresDb,
	)
	repo, err := pg.NewPostgresRepository("postgres", addr)
	if err != nil {
		log.Printf("Error creating postgres repository")
	}
	pg.SetRepository(repo)
	ctx := context.Background()
	err = pg.TruncateOrders(ctx)
	if err != nil {
		log.Printf(err.Error())
	}

	return func(tb testing.TB) {
		log.Println("teardown suite")
		err = pg.TruncateOrders(ctx)
		if err != nil {
			log.Printf(err.Error())
		}
		defer func() {
			err := pg.CloseConn()
			if err != nil {
				return
			}
		}()
	}
}

func TestGetOrderBD(t *testing.T) {
	teardownSuite := setupSuiteDB(t)
	defer teardownSuite(t)
	orderUid := "123"
	order := GetExampleOrder(orderUid)
	ctx := context.Background()
	err := pg.InsertOrder(ctx, order)
	assert.Nil(t, err)

	getOrder, err := pg.GetOrderByUID(ctx, orderUid)
	assert.Nil(t, err)

	assert.Equal(t, getOrder, order)
}

func TestGetOrdersBD(t *testing.T) {
	teardownSuite := setupSuiteDB(t)
	defer teardownSuite(t)
	orders := []schemas.Order{
		GetExampleOrder("1234"),
		GetExampleOrder("12345"),
	}
	ctx := context.Background()
	for _, order := range orders {
		err := pg.InsertOrder(ctx, order)
		assert.Nil(t, err)
	}
	gotOrders, err := pg.GetOrders(ctx)
	assert.Nil(t, err)
	assert.Equal(t, orders, gotOrders)
}

func TestInsertOrderDB(t *testing.T) {
	teardownSuite := setupSuiteDB(t)
	defer teardownSuite(t)
	orderUid := "123"
	order := GetExampleOrder(orderUid)
	ctx := context.Background()
	err := pg.InsertOrder(ctx, order)
	assert.Nil(t, err)
}

func TestGetNotExistsOrderDB(t *testing.T) {
	teardownSuite := setupSuiteDB(t)
	defer teardownSuite(t)
	ctx := context.Background()
	_, err := pg.GetOrderByUID(ctx, "123")
	assert.Equal(t, err.Error(), "not found")
}
