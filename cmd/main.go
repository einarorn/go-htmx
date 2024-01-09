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
				{Make: "Mercedes", Model: "C300e", Year: 2020, Description: "A powerful plug-in hybrid luxury sedan with low milage and 50 km electric range."},
				{Make: "Volvo", Model: "XC90 T8", Year: 2016, Description: "A great family wagon with room for seven people. A plug-in hyprid with good economy 34 km electric range."},
			},
		}
		tmpl.Execute(w, autos)
	}

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
