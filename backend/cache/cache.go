package cache

import (
	"github.com/Proxypepe/wb-web/backend/schemas"
)

// Cache interface is used to provide an abstraction for the use of caching services
type Cache interface {
	// SaveOrder method is used to save one order
	SaveOrder(order schemas.Order) error
	// SaveOrders method is used to save multiple orders
	SaveOrders(orders []schemas.Order) error
	// GetOrder method is used to get an order by its uid
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
