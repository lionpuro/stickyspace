package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lionpuro/stickyspace/auth"
	"log"
	"net/http"
	"os"
)

func main() {
	env, err := getEnv(
		"SERVER_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_HOST_PORT",
		"POSTGRES_DB",
	)
	if err != nil {
		log.Fatal(err)
	}

	db, err := newDB(
		env["POSTGRES_USER"],
		env["POSTGRES_PASSWORD"],
		env["POSTGRES_HOST_PORT"],
		env["POSTGRES_DB"],
	)
	if err != nil {
		log.Fatalf("init database: %v", err)
	}

	as, err := auth.NewService()
	if err != nil {
		log.Fatal(err)
	}
	us := NewUserService(db)

	port := env["SERVER_PORT"]

	srv := newServer(as, us)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: srv,
	}

	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(server.ListenAndServe())
}

func newServer(as *auth.Service, us *UserService) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("POST /signup", handleSignUp(as, us))
	return r
}

func getEnv(vars ...string) (map[string]string, error) {
	envmap := map[string]string{}
	for _, key := range vars {
		val := os.Getenv(key)
		if val == "" {
			return nil, fmt.Errorf("unset env variable %s", key)
		}
		envmap[key] = val
	}
	return envmap, nil
}
