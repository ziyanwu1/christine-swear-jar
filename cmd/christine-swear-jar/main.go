package main

import (
	"christineswearjar/internal/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load("config/.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v\n", err)
    }

    http.HandleFunc("/", handlers.ServeIndexHTML)
    http.HandleFunc("/api/increment", handlers.IncrementHandler)
    http.HandleFunc("/api/gettotal", handlers.GetTotalHandler)

    addr := "localhost:8080"
    err = http.ListenAndServe(addr, nil)
    if err != nil {
        fmt.Printf("oops server broke\n")
        fmt.Printf("error: %v\n", err)
    }
}