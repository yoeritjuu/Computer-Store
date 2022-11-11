package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/yoeritjuu/Computer-Store/pkg/handlers"
)

func main() {
	godotenv.Load()
	http.HandleFunc("/", handlers.GetPartsHandler)
	http.ListenAndServe("localhost:8000", nil)
}
