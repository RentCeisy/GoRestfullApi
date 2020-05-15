package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Listening server on port 3001")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/health", HealthHandler)

	http.ListenAndServe(":3001", nil)
}

func IndexHandler(w http.ResponseWriter, r * http.Request) {
	res, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	res.ExecuteTemplate(w, "index", nil)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("templates/health.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	res.ExecuteTemplate(w, "health", nil)
}
