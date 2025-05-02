package main

import (
	"encoding/json"
	"errors"
	"github.com/lionpuro/stickyspace/auth"
	"net/http"
	"net/mail"
	"unicode"

	_ "github.com/joho/godotenv/autoload"
)

func handleSignUp(as *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var u newUser
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil || u.Email == "" || u.Name == "" || u.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := validateNewUser(u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		if _, err := as.CreateUser(r.Context(), u.Email, u.Password, u.Name); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode("success")
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type newUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func validateNewUser(u newUser) error {
	if len(u.Name) < 2 {
		return errors.New("Name must be at least 2 characters long")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("Invalid email")
	}
	if ok := validatePassword(u.Password); !ok {
		return errors.New("Password must be at least 8 characters long, include at least one number and at least one uppercase letter")
	}
	return nil
}

func validatePassword(pw string) bool {
	chars := 0
	hasNumber := false
	hasUpper := false
	for _, c := range pw {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
			chars++
		case unicode.IsUpper(c):
			hasUpper = true
			chars++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			chars++
		case unicode.IsLetter(c):
			chars++
		}
	}
	switch {
	case chars < 8:
		return false
	case !hasNumber:
		return false
	case !hasUpper:
		return false
	}
	return true
}
