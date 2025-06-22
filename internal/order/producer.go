package order

import(
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"encoding/json"
	"fmt"
	"time"
)

func Producer(order Order){

	data, err:= json.Marshal(order)

	if err!=nil{
		log.Fatal(err)
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "orders",
	})
	defer writer.Close()

	time.Sleep(1 * time.Second)

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(order.Product),
			Value: data,
		},
	)

	if err != nil {
		log.Fatal("Producer error:", err)
	}
	fmt.Println("✅ Kafkaesque Producer sent message.")
}