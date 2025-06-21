package order

import(
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

var kafkaNewWriter *kafka.Writer

func InitProducer(broker, topic string){
	kafkaNewWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
}

func ProduceOrder(order Order) error{
	data, err:= json.Marshal(order)
	if err!=nil{
		return err
	}

	msg:= kafka.Message{
		Key: []byte(order.Product),
		Value: data,
	}

	return kafkaNewWriter.WriteMessages(context.Background(), msg)
}