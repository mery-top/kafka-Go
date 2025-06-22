package order

import(
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"fmt"
	"encoding/json"
)

func Consumer(){
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "orders",
		GroupID: "orders-group",
	})
	defer reader.Close()

	fmt.Println("👂 Kafkaesque Consumer listening...")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Consumer error:", err)
		}
		var o Order
		json.Unmarshal(msg.Value, &o)
		fmt.Printf("Order Received: %s %d\n", o.Product, o.Quantity)
	}
}