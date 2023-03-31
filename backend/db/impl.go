package db

import (
	"github.com/Proxypepe/wb-web/backend/schemas"
)

type RepositoryImpl interface {
	insertDelivery(delivery schemas.Delivery) (int, error)
	insertPayment(payment schemas.Payment) (int, error)
	insertItem(item schemas.Item, orderUID string) (int, error)
	insertOrder(order schemas.Order, deliveryID int, paymentID int) error
	getDeliveryIDByDelivery(delivery schemas.Delivery) (int, error)
	getPaymentIDByPayment(payment schemas.Payment) (int, error)
	getItemsIDByOrderUID(uid string) ([]schemas.Item, error)
}
