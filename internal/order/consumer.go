package order

import(
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"encoding/json"
	"fmt"
)

func main(){
	reader:= kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "order-topic",
		GroupID: "order-service",
	})

	log.Println("Consumer Started...")
	for{
		msg, err:= reader.ReadMessage(context.Background())
		if err!=nil{
			log.Println(err)
			continue
		}

		var o Order
		json.Unmarshal(msg.Value, &o)
		fmt.Printf("Order Received: %s %d\n", o.Product, o.Quantity)

	}
}