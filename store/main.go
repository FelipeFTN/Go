package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Computer", Description: "Razer", Price: 3999.99, Amount: 45},
		{"MacBook", "Pro, Big Sur, 16GB", 14999.99, 10},
		{"Glasses", "White Smart Glasses", 2499.99, 5},
		{"Smart Watch", "Apple", 5999.99, 12},
	}

	temp.ExecuteTemplate(w, "index", products)
}
