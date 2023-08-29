package main

import (
	"context"
	"fmt"
	"time"
	"flag"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Wait for kafka to go online
	time.Sleep(time.Duration(20) * time.Second)

	var partition = flag.Int("p", 0, "partition")
	var offset = flag.Int("o", 0, "offset")
	var sleep = flag.Int("w", 0, "wait")
	flag.Parse()
	Read(*partition, "data", *offset, *sleep)
	fmt.Println("Shutting down reader")
}

func Read(partition int, topic string, timestampOffsetSeconds int, sleep int) {
	fmt.Printf("Reading from partition %d, with an offset of %ds after %ds\n", partition, timestampOffsetSeconds, sleep)

	time.Sleep(time.Duration(sleep) * time.Second)

	start := time.Now().Add(time.Duration(-timestampOffsetSeconds) * time.Second)
	fmt.Printf("Start reading at %s\n", start.Format(time.UnixDate))

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"broker:9092"},
		Topic:     topic,
		Partition: partition,
		//GroupID:   "data-readers", //if the GroupID is set we can't rewind
	})
	err := r.SetOffsetAt(context.Background(), start)

	if err != nil {
		fmt.Printf("failed to rewind: %s\n", err)
		return
	}

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d, time %s: %s\n", m.Offset, m.Time.Format(time.UnixDate), string(m.Value))
	}
	
	if err := r.Close(); err != nil {
		fmt.Printf("failed to close reader: %s", err)
	}
}
