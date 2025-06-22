package order

import (
	"fmt"
	"net/http"
	"strconv"
)

func OrderHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		http.ServeFile(w, r, "/Users/meerthikasr/Desktop/kafka-go/web/index.html")
		return
	}

	if r.Method == http.MethodPost{
		err:= r.ParseForm();
		if err!=nil{
			http.Error(w, "Parse Error", http.StatusBadRequest)
			return
		}

		product:= r.FormValue("product")
		quantity, _:= strconv.Atoi(r.FormValue("quantity"))
		order:= Order{
			Product: product,
			Quantity: quantity,
		}
		fmt.Println(order.Product, order.Quantity)

		CreateTopic()
		go Consumer(order)
		go Producer(order)

		w.Write([]byte("Order Placed!"))
	}
}