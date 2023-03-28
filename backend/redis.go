package main

import (
	cache "github.com/Proxypepe/wb-web/backend/cache"
	"github.com/Proxypepe/wb-web/backend/http"
	"github.com/go-redis/redis"
	"log"
)

func main() {

	red, err := cache.NewRedisStore(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err != nil {
		log.Fatal(err)
	}

	cache.SetCacheService(red)

	server := http.NewServer()
	server.Run()

	//addr := fmt.Sprintf("postgres://%s:%s@localhost:17200/%s?sslmode=disable", "alex", "postgres", "wb")
	//repo, err := pg.NewPostgresRepository("postgres", addr)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//	fmt.Printf("Error creating postgres repository")
	//	return
	//}
	//
	//pg.SetRepository(repo)
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()
	//
	//orders, err := pg.GetOrders(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//err = cache.SaveOrders(orders)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//err = cache.SaveOrder(schemas.Order{
	//	OrderUID:    "b563feb7b2b84b6test",
	//	TrackNumber: "WBILMTESTTRACK",
	//	Entry:       "WBIL",
	//	Delivery: schemas.Delivery{
	//		Name:    "Test Testov",
	//		Phone:   "+9720000000",
	//		Zip:     "2639809",
	//		City:    "Kiryat Mozkin",
	//		Address: "Ploshad Mira 15",
	//		Region:  "Kraiot",
	//		Email:   "test@gmail.com",
	//	},
	//	Payment: schemas.Payment{
	//		Transaction:  "b563feb7b2b84b6test",
	//		RequestID:    "",
	//		Currency:     "USD",
	//		Provider:     "wbpay",
	//		Amount:       1817,
	//		PaymentDt:    1637907727,
	//		Bank:         "alpha",
	//		DeliveryCost: 1500,
	//		GoodsTotal:   317,
	//		CustomFee:    0,
	//	},
	//	Items: []schemas.Item{
	//		{
	//			ChrtID:      9934930,
	//			TrackNumber: "WBILMTESTTRACK",
	//			Price:       453,
	//			Rid:         "ab4219087a764ae0btest",
	//			Name:        "Mascaras",
	//			Sale:        30,
	//			Size:        "0",
	//			TotalPrice:  317,
	//			NmID:        2389212,
	//			Brand:       "Vivienne Sabo",
	//			Status:      202,
	//		},
	//		{
	//			ChrtID:      9934931,
	//			TrackNumber: "WBILMTESTTRACK",
	//			Price:       555,
	//			Rid:         "dddd",
	//			Name:        "dsa",
	//			Sale:        20,
	//			Size:        "1",
	//			TotalPrice:  520,
	//			NmID:        2389212,
	//			Brand:       "Vivienne Sabo",
	//			Status:      202,
	//		},
	//	},
	//	Locale:            "en",
	//	InternalSignature: "",
	//	CustomerID:        "test",
	//	DeliveryService:   "meest",
	//	Shardkey:          "9",
	//	SmID:              99,
	//	DateCreated:       time.Now().String(),
	//	OofShard:          "1",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//order, err := cache.GetOrder("b563feb7b2b84b6test1")
	//if err == redis.Nil {
	//	return
	//}
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//orderB, err := json.Marshal(order)
	//fmt.Printf(string(orderB))
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
}
