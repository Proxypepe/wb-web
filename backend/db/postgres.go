package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Proxypepe/wb-web/backend/schemas"
	_ "github.com/lib/pq"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(driver string, url string) (*PostgresRepository, error) {
	db, err := sql.Open(driver, url)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (repo *PostgresRepository) CloseConn() error {
	err := repo.db.Close()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func (repo *PostgresRepository) InsertOrder(ctx context.Context, order schemas.Order) error {

	var (
		deliveryID int
		paymentID  int
		err        error
	)
	paymentID, err = repo.insertPayment(order.Payment)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	deliveryID, err = repo.insertDelivery(order.Delivery)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = repo.insertOrder(order, deliveryID, paymentID)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	for _, item := range order.Items {
		_, err = repo.insertItem(item, order.OrderUID)
		if err != nil {
			log.Print(err.Error())
			return err
		}
	}

	return nil
}

func (repo *PostgresRepository) GetOrders(ctx context.Context) ([]schemas.Order, error) {
	rows, err := repo.db.Query(`SELECT
   public.order.order_uid,
   public.order.track_number,
   public.order.entry,
   public.order.locale,
   public.order.internal_signature,
   public.order.customer_id,
   public.order.delivery_service,
   public.order.shardkey,
   public.order.sm_id,
   public.order.date_created,
   public.order.oof_shard,
   public.payment.transaction,
   public.payment.request_id,
   public.payment.currency,
   public.payment.provider,
   public.payment.amount,
   public.payment.payment_dt,
   public.payment.bank,
   public.payment.delivery_cost,
   public.payment.goods_total,
   public.payment.custom_fee,
   public.delivery.name,
   public.delivery.phone,
   public.delivery.zip,
   public.delivery.city,
   public.delivery.address,
   public.delivery.region,
   public.delivery.email
	FROM public.order
        JOIN public.payment
             ON public."order".payment_id = public.payment.id
        Join public.delivery
             ON public."order".delivery_id = public.delivery.id`)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	var orders []schemas.Order
	for rows.Next() {
		order := schemas.Order{}
		if err := rows.Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
			&order.Payment.Transaction,
			&order.Payment.RequestID,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.PaymentDt,
			&order.Payment.Bank,
			&order.Payment.DeliveryCost,
			&order.Payment.GoodsTotal,
			&order.Payment.CustomFee,
			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,
		); err == nil {
			items, err := repo.getItemsIDByOrderUID(order.OrderUID)
			if err != nil {
				log.Print(err.Error())
				return nil, err
			}
			order.Items = items
			orders = append(orders, order)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return orders, nil
}

func (repo *PostgresRepository) GetOrderByUID(ctx context.Context, uid string) (schemas.Order, error) {
	rows, err := repo.db.Query(fmt.Sprintf(`SELECT
   public.order.order_uid,
   public.order.track_number,
   public.order.entry,
   public.order.locale,
   public.order.internal_signature,
   public.order.customer_id,
   public.order.delivery_service,
   public.order.shardkey,
   public.order.sm_id,
   public.order.date_created,
   public.order.oof_shard,
   public.payment.transaction,
   public.payment.request_id,
   public.payment.currency,
   public.payment.provider,
   public.payment.amount,
   public.payment.payment_dt,
   public.payment.bank,
   public.payment.delivery_cost,
   public.payment.goods_total,
   public.payment.custom_fee,
   public.delivery.name,
   public.delivery.phone,
   public.delivery.zip,
   public.delivery.city,
   public.delivery.address,
   public.delivery.region,
   public.delivery.email
	FROM public.order
        JOIN public.payment
             ON public."order".payment_id = public.payment.id
        Join public.delivery
             ON public."order".delivery_id = public.delivery.id
	WHERE public.order.order_uid = '%s'
	LIMIT 1;`, uid))
	if err != nil {
		log.Print(err.Error())
		return schemas.Order{}, err
	}
	items, err := repo.getItemsIDByOrderUID(uid)
	if err != nil {
		log.Print(err.Error())
		return schemas.Order{}, err
	}
	var order schemas.Order
	order.Items = items
	order.Delivery = schemas.Delivery{}
	order.Payment = schemas.Payment{}
	if rows.Next() {
		err := rows.Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
			&order.Payment.Transaction,
			&order.Payment.RequestID,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.PaymentDt,
			&order.Payment.Bank,
			&order.Payment.DeliveryCost,
			&order.Payment.GoodsTotal,
			&order.Payment.CustomFee,
			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,
		)
		if err != nil {
			return schemas.Order{}, err
		}
		return order, nil
	}
	return schemas.Order{}, errors.New("not found")
}

func (repo *PostgresRepository) TruncateOrders(ctx context.Context) error {
	query := `TRUNCATE TABLE public.item, public.delivery, public.payment, public.order CASCADE`
	_, err := repo.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) insertOrder(order schemas.Order, deliveryID int, paymentID int) error {
	query := `INSERT INTO public.order (
           order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Print(err.Error())
			return
		}
	}(stmt)

	row := stmt.QueryRow(
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		deliveryID,
		paymentID,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)
	if row.Err() != nil {
		log.Print(err.Error())
		return row.Err()
	}

	return nil
}

func (repo *PostgresRepository) insertDelivery(delivery schemas.Delivery) (int, error) {
	query := `INSERT INTO public.delivery (
            name, phone, zip, city, address, region, email) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Print(err.Error())
			return
		}
	}(stmt)
	var id int
	err = stmt.QueryRow(
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	).Scan(&id)
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	return id, nil
}

func (repo *PostgresRepository) insertPayment(payment schemas.Payment) (int, error) {
	query := `INSERT INTO public.payment (
			transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			RETURNING id`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Print(err.Error())
			return
		}
	}(stmt)
	var id int
	err = stmt.QueryRow(
		payment.Transaction,
		payment.RequestID,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	).Scan(&id)
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	return id, nil
}

func (repo *PostgresRepository) insertItem(item schemas.Item, orderUID string) (int, error) {
	query := `INSERT INTO public.item (
            order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
			RETURNING id`
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Print(err.Error())
			return
		}
	}(stmt)
	var id int
	err = stmt.QueryRow(
		orderUID,
		item.ChrtID,
		item.TrackNumber,
		item.Price,
		item.Rid,
		item.Name,
		item.Sale,
		item.Size,
		item.TotalPrice,
		item.NmID,
		item.Brand,
		item.Status,
	).Scan(&id)
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	return id, nil
}

func (repo *PostgresRepository) getDeliveryIDByDelivery(delivery schemas.Delivery) (int, error) {
	rows, err := repo.db.Query(fmt.Sprintf(`SELECT id FROM delivery d 	
          					WHERE 	d.name 	= 	'%s'
                            AND 	d.phone = 	'%s'
                            AND 	d.zip 	= 	'%s'
                            AND 	d.address = '%s'
                            AND 	d.city 	= 	'%s'
                            AND 	d.region = 	'%s'
                            AND 	d.email = 	'%s'
                        LIMIT 1;`,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.Address,
		delivery.City,
		delivery.Region,
		delivery.Email,
	))
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Print(err.Error())
			return 0, err
		}
		return id, err
	}
	return 0, errors.New("nil select")
}

func (repo *PostgresRepository) getPaymentIDByPayment(payment schemas.Payment) (int, error) {
	rows, err := repo.db.Query(fmt.Sprintf(`SELECT id FROM payment p 
          WHERE  p.transaction = '%s'
          AND    p.request_id = '%s'
          AND    p.currency = '%s'
          AND    p.provider = '%s'
          AND    p.amount = %v
          AND    p.payment_dt = %v
          AND    p.bank = '%s'
          AND    p.delivery_cost = %v
          AND    p.goods_total = %v
          AND    p.custom_fee = %v
		LIMIT 1;`,
		payment.Transaction,
		payment.RequestID,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	))
	if err != nil {
		log.Print(err.Error())
		return 0, err
	}
	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Print(err.Error())
			return 0, err
		}
		return id, err
	}
	return 0, errors.New("nil select")
}

func (repo *PostgresRepository) getItemsIDByOrderUID(uid string) ([]schemas.Item, error) {
	rows, err := repo.db.Query(fmt.Sprintf(`SELECT 
    		   chrt_id,
			   track_number,
			   price,
			   rid,
			   name,
			   sale,
			   size,
			   total_price,
			   nm_id,
			   brand,
			   status
			FROM public.item
			WHERE public.item.order_uid = '%s';`, uid))
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	var items []schemas.Item
	for rows.Next() {
		item := schemas.Item{}
		if err = rows.Scan(
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.Brand,
			&item.Status); err == nil {
			items = append(items, item)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
