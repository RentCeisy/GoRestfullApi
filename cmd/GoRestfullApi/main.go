package main

import (
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	port, exists := os.LookupEnv("PORT")
	if exists {
		log.Println("Listening server on port " + port)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./../../assets/"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/health", HealthHandler)

	http.ListenAndServe(":" + port, nil)
}

func IndexHandler(w http.ResponseWriter, r * http.Request) {
	res, err := template.ParseFiles("../../templates/index.html", "../../templates/header.html", "../../templates/footer.html")
	if err != nil {
		log.Println(err.Error())
		return
	}

	res.ExecuteTemplate(w, "index", nil)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Server")
	w.WriteHeader(200)
}
