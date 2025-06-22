# 📦 Kafka Order Processing System

A lightweight **event-driven order processing system** built using **Go** and **Apache Kafka**. This project demonstrates the core concepts of Kafka — such as producers, consumers, and topics — through a simple order placement use case.


## ✅ Features

- Users can place orders via a web form (e.g., `Jalebi`, `Laptop`, `Lion`)
- Orders are sent to a Kafka topic named `orders`
- A Kafka consumer listens for new orders and logs them
- Easily extendable: Add more consumers for shipping, inventory, etc.


## 🚀 Tech Stack

| Tech         | Purpose                           |
|--------------|------------------------------------|
| Go (Golang)  | Backend, Kafka Producer/Consumer   |
| Kafka        | Messaging broker for event flow    |
| Zookeeper    | Kafka coordination                 |
| Docker       | Containerized Kafka + Zookeeper    |
| HTML/CSS     | Simple web UI for placing orders   |

---


### 🐳 1. Start Kafka and Zookeeper

Make sure Docker is installed, then run:

```bash
docker-compose up -d
```

### 🌐 2. Start the HTTP Server (Producer API)

This serves the HTML form and sends orders to Kafka.

```bash
cd cmd/api
go run main.go

```
Visit in your browser:

http://localhost:8080

---

### ✅ Expected Output from Consumer

🧾 Kafka Consumer started...
🛒 Order Received: Product=Pen | Quantity=5



## ⚙️ Project Structure (Overview)

```
📁 order-app/
├── cmd/
│   ├── api/              # HTTP Server (handles HTML form and POST)
│ 
├── internal/
│   ├── order/     # Domain model
│   │   ├── createTopic.go  
│   │   ├── handler.go    # Business logic
│   │   ├── model.go      # Domain model
│   │   ├── producer.go   # Kafka producer
│   │   └── consumer.go   # Kafka consumer
├── web/
│   └── index.html        # Order form
├── go.mod
└── docker-compose.yml    # Kafka and Zookeeper
```

## Extendible Project Structure 
```
kafka-microservices/
├── order-service/
│   ├── main.go                 # Starts order service
│   ├── handler/
│   │   └── order_handler.go    # HTTP handler for placing orders
│   ├── kafka/
│   │   ├── producer.go         # Kafka producer for orders
│   │   └── consumer.go         # (optional) consumer, if order service wants to consume
│   └── models/
│       └── order.go            # Order struct definition
│
├── shipping-service/
│   ├── main.go                 # Starts shipping consumer
│   ├── kafka/
│   │   └── consumer.go         # Shipping-specific consumer
│   └── models/
│       └── order.go            # Reused Order struct (copy or shared)
│
├── web/
│   └── index.html              # Frontend form for placing order
│
├── docker-compose.yml          # Runs Zookeeper + Kafka
└── README.md                   # Project instructions
```

---

Feel free to fork this project, customize it, and make it your own!

🙌 Happy coding!
