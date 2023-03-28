package cache

import (
	"github.com/Proxypepe/wb-web/backend/schemas"
)

type Cache interface {
	SaveOrder(order schemas.Order) error
	SaveOrders(orders []schemas.Order) error
	GetOrder(uid string) (*schemas.Order, error)
}

var impl Cache

func SetCacheService(service Cache) {
	impl = service
}
func SaveOrder(order schemas.Order) error {
	return impl.SaveOrder(order)
}

func SaveOrders(orders []schemas.Order) error {
	return impl.SaveOrders(orders)
}

func GetOrder(uid string) (*schemas.Order, error) {
	return impl.GetOrder(uid)
}
