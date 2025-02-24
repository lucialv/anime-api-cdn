package main

import (
	"log"
	"net/http"

	"github.com/lucialv/anime-api-cdn/endpoints"
)

func main() {
	http.HandleFunc("/hugs", endpoints.HugsHandler)
	http.HandleFunc("/pats", endpoints.PatsHandler)

	log.Println("Servidor local en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
