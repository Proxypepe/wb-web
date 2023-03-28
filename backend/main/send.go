package main

import (
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

	// Отправляем сообщения в канал test-channel
	for i := 0; i < 10; i++ {
		msg := []byte("Hello NATS Streaming!")
		if err := sc.Publish("bestellungen", msg); err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}
		log.Printf("Sent message: %s", msg)
		time.Sleep(1 * time.Second)
	}
}
