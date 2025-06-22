package order

import(
	"log"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func CreateTopic(){
	conn, err:= kafka.Dial("tcp", "localhost:29092")
	if err != nil {
		log.Fatal("❌ Failed to connect to Kafka for topic creation:", err)
	}

	defer conn.Close()


	//Get the controller

	controller, err := conn.Controller()
	if err != nil {
		log.Fatal("❌ Failed to get controller:", err)
	}

	controllerConn, err:= kafka.Dial("tcp",fmt.Sprintf("%s:%d", controller.Host, controller.Port))
	if err != nil {
		log.Fatal("❌ Failed to get controllerConnection:", err)
	}

	defer controllerConn.Close()

	topic:= "orders"

	err = controllerConn.CreateTopics(kafka.TopicConfig{
		Topic: topic,
		NumPartitions: 1,
		ReplicationFactor: 1,
	})

	if err != nil {
		log.Println("⚠️  Topic creation warning (may already exist):", err)
	} else {
		log.Println("✅ Kafka topic created:", topic)
	}
}