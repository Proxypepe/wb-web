package nats

import (
	"encoding/json"
	"fmt"
	_ "github.com/Proxypepe/wb-web/backend/schemas"
	"github.com/Proxypepe/wb-web/backend/tests"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	// Подключаемся к nats-streaming
	sc, err := stan.Connect("test-cluster", "test")
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}

	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			log.Printf("Error: %s\n", err)
		}
	}(sc)

	for i := 3; i < 5; i++ {
		order := tests.GetExampleOrder(fmt.Sprintf("test-%d", i))
		bytes, err := json.Marshal(order)
		//bytes := []byte("Hello NATS Streaming!")
		if err != nil {
			if err != nil {
				log.Printf("Error: %s\n", err)
			}
		}
		if err = sc.Publish("backend", bytes); err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}
		log.Printf("Sent message: %s", bytes)
		time.Sleep(1 * time.Second)
	}
}
