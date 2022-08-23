package main

import (
	"learn-webdev-go/handler"
	"log"
	"net/http"
)

func main() {
	//mux bisa kita gunakan untuk serve routing
	mux := http.NewServeMux()

	//untuk mengatur endpoint path dan handlernya
	mux.HandleFunc("/", handler.RootHandler) //root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	//semua route yang tidak didaftarkan akan selalu masuk ke root endpoints

	log.Println("Starting web on port 8080")
	//	listen port 8080 dengan mux variablenya
	err := http.ListenAndServe(":8080", mux)
	//jika ada error akan dilakukan log fatal
	log.Fatal(err)
}
