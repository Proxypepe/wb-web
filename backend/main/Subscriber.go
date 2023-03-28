package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"sync"
)

func block() {
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}

func main() {
	clusterID := "test-cluster"
	clientID := "test-client"
	sc, _ := stan.Connect(clusterID, clientID)

	sub, err := sc.Subscribe("bestellungen", func(m *stan.Msg) {
		fmt.Printf("Got: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Printf("Error")
	}
	defer func(sub stan.Subscription) {
		err := sub.Unsubscribe()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}(sub)
	defer func(sub stan.Subscription) {
		err := sub.Close()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}(sub)
	block()

}
