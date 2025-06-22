package main

import(
	"net/http"
	"log"
	"kafka-go/internal/order"
)

func main(){
	http.HandleFunc("/", order.OrderHandler)
	http.HandleFunc("/order", order.OrderHandler)
	log.Println("Server Started at LOCALHOST:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}