package auth

import (
	"context"
	"errors"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/errorutils"
	"google.golang.org/api/option"

	_ "github.com/joho/godotenv/autoload"
)

type User struct {
	ID            string
	Email         string
	EmailVerified bool
}

type Service struct {
	FirebaseApp *firebase.App
	Client      *auth.Client
}

func NewService() (*Service, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase.json")
	fb, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	cl, err := fb.Auth(ctx)
	if err != nil {
		return nil, err
	}

	as := &Service{
		FirebaseApp: fb,
		Client:      cl,
	}
	return as, nil
}

func (s *Service) VerifyToken(ctx context.Context, token string) (*auth.Token, error) {
	tkn, err := s.Client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return tkn, nil
}

func (s *Service) CreateUser(ctx context.Context, email, password, name string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		DisplayName(name).
		Disabled(false)
	u, err := s.Client.CreateUser(ctx, params)
	if err != nil {
		return nil, errors.New(errorMsg(err))
	}
	return u, nil
}

func errorMsg(err error) string {
	switch {
	case errorutils.IsInternal(err):
		return "Internal server error"
	case auth.IsEmailAlreadyExists(err):
		return "User already exists"
	default:
		return err.Error()
	}
}
