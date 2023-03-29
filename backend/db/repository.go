package db

import (
	"context"
	"github.com/Proxypepe/wb-web/backend/schemas"
)

type Repository interface {
	CloseConn() error
	InsertOrder(ctx context.Context, order schemas.Order) error
	GetOrders(ctx context.Context) ([]schemas.Order, error)
	GetOrderByUID(ctx context.Context, uid string) (schemas.Order, error)
	TruncateOrders(ctx context.Context) error
}

var impl Repository

func SetRepository(repo Repository) {
	impl = repo
}

func CloseConn() error {
	return impl.CloseConn()
}

func InsertOrder(ctx context.Context, order schemas.Order) error {
	return impl.InsertOrder(ctx, order)
}

func GetOrders(ctx context.Context) ([]schemas.Order, error) {
	return impl.GetOrders(ctx)
}

func GetOrderByUID(ctx context.Context, uid string) (schemas.Order, error) {
	return impl.GetOrderByUID(ctx, uid)
}

func TruncateOrders(ctx context.Context) error {
	return impl.TruncateOrders(ctx)
}
