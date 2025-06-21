package order

import(
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func InitProducer(broker, topic string){
	kafkaWriter = &kafka.Writer{
		Addr: kafka.TCP(broker),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	}
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

	return kafkaWriter.WriteMessages(context.Background(), msg)
}