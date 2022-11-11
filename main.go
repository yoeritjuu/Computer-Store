package main

import (
	"net/http"

	"github.com/yoeritjuu/Computer-Store/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.GetPartsHandler)
	http.ListenAndServe("localhost:8000", nil)
}
