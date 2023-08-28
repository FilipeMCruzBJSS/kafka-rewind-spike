package main

import (
	"time"
    "math/rand"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Wait for kafka to go online
	time.Sleep(time.Duration(20) * time.Second)
	
	Write("data", 5)
	fmt.Println("Shutting down writer")
}


func Write(topic string, tickSeconds int) {
	w := &kafka.Writer{
		Addr:     kafka.TCP("broker:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	interval := time.Duration(tickSeconds) * time.Second
	// create a new Ticker
	tk := time.NewTicker(interval)
	
	var err error

	i := 0
	for range tk.C {
		i++
		err = w.WriteMessages(context.Background(),
			kafka.Message{
				Key  : []byte(string(rand.Intn(10))),
				Value: []byte(time.Now().Format(time.UnixDate)),
			},
		)
		if err != nil || i > 100 {
			break;
		}
	}
	
	if err != nil {
		fmt.Printf("failed to write messages: %s", err)
	}
	
	if err := w.Close(); err != nil {
		fmt.Printf("failed to close writer: %s", err)
	}
}
