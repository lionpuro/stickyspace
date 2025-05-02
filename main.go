package main

import (
	"fmt"
	"github.com/lionpuro/stickyspace/auth"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	as, err := auth.NewService()
	if err != nil {
		log.Fatal(err)
	}

	srv := newServer(as)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: srv,
	}

	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(server.ListenAndServe())
}

func newServer(as *auth.Service) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("POST /signup", handleSignUp(as))
	return r
}
