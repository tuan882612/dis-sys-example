package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func main() {
	// Kafka configuration
	brokerAddress := "dory.srvs.cloudkafka.com:9094"
	username := "fcxhebea"                         // Default user
	password := "uBsD9qt8-IXHxr2qu6ecmuxLinNOoCGu" // Password

	// Create SASL/SCRAM mechanism
	mechanism, err := scram.Mechanism(scram.SHA512, username, password)
	if err != nil {
		log.Fatal("failed to create mechanism:", err)
	}

	// Create a dialer with SASL/SCRAM authentication
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	// Create a new Kafka reader using the broker and dialer
	conn, err := dialer.DialContext(context.Background(), "tcp", brokerAddress)
	if err != nil {
		log.Fatal("failed to dial:", err)
	}
	defer conn.Close()

	// List topics
	partitions, err := conn.ReadPartitions()
	if err != nil {
		log.Fatal("failed to read partitions:", err)
	}

	topics := map[string]struct{}{}
	for _, p := range partitions {
		topics[p.Topic] = struct{}{}
	}

	fmt.Println("Available Topics:")
	for topic := range topics {
		fmt.Printf(" - %s\n", topic)
	}
}
