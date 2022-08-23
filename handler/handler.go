package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halo dunia, lagi belajar golang nih"))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	//perlu handle not found untuk selain "/"
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//untuk membaca file html
	templ, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	//selalu gunakan pengecekan error
	if err != nil {
		//biasakan lakukan logging untuk error tracing developer
		log.Println(err)
		//biasakan gunakan error terstruktur untuk user, detail error cukup untuk developer saja
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	//create dummy data map
	data := map[string]string{
		"title":   "Halo ini title",
		"content": "Halo ini content golang",
	}

	//untuk mengexecute file html
	err = templ.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	//cara mendapatkan query string
	//every query string will be string type
	id := r.URL.Query().Get("id")
	//convert from string to int
	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}
	//menulis string dengan fprintf
	//fmt.Fprintf(w, "Product page: %d", idNum)

	//	load html template
	templ, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		//biasakan lakukan logging untuk error tracing developer
		log.Println(err)
		//biasakan gunakan error terstruktur untuk user, detail error cukup untuk developer saja
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	//  baca data
	data := map[string]int{
		"content": idNum,
	}
	//  tampilkan template
	err = templ.Execute(w, data)
	if err != nil {
		//biasakan lakukan logging untuk error tracing developer
		log.Println(err)
		//biasakan gunakan error terstruktur untuk user, detail error cukup untuk developer saja
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

}
