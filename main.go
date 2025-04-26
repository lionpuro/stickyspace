package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	router := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(server.ListenAndServe())
}
