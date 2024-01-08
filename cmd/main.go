package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/einarorn/go-htmx/model"
)

func main() {
	indexHandler := func(w http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("view/index.html"))
		autos := map[string][]model.Automobile{
			"Autos": {
				{Make: "Mercedes", Model: "C300e", Year: 2020},
				{Make: "Volvo", Model: "XC90 T8", Year: 2016},
			},
		}
		tmpl.Execute(w, autos)
	}

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
