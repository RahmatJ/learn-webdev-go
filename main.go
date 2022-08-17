package main

import (
	"log"
	"net/http"
)

func main() {
	//mux bisa kita gunakan untuk serve routing
	mux := http.NewServeMux()

	//untuk mengatur endpoint path dan handlernya
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", rootHandler) //root
	//semua route yang tidak didaftarkan akan selalu masuk ke root endpoints

	log.Println("Starting web on port 8080")
	//	listen port 8080 dengan mux variablenya
	err := http.ListenAndServe(":8080", mux)
	//jika ada error akan dilakukan log fatal
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halo dunia, lagi belajar golang nih"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//perlu handle not found untuk selain "/"
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome home"))
}
