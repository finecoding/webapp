package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/finecoding/app/views"
)

var homeView *views.View
var contactView *views.View
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)

	}
}

func contact(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)

	}

}
func main() {

	var err error

	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)

}
