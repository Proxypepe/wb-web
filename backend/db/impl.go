package db

import (
	"github.com/Proxypepe/wb-web/backend/schemas"
)

type RepositoryImpl interface {
	InsertDelivery(delivery schemas.Delivery) (int, error)
	InsertPayment(payment schemas.Payment) (int, error)
	InsertItem(item schemas.Item, orderUid string) (int, error)
	_insertOrder(order schemas.Order, deliveryId int, paymentId int) error
	GetDeliveryIdByDelivery(delivery schemas.Delivery) (int, error)
	GetPaymentIdByPayment(payment schemas.Payment) (int, error)
	GetItemsIdByOrderUid(uid string) ([]schemas.Item, error)
}
