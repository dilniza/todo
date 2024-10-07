package service

import (
	"context"
	"todo/api/models" // Import the password package
	"todo/pkg/password"
	"todo/storage/mongodb"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (models.User, error)
}

type authService struct {
	userRepo mongodb.UserRepo
}

func NewAuthService(userRepo mongodb.UserRepo) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(ctx context.Context, username, password string) (models.User, error) {
	user, err := s.userRepo.GetUser(ctx, username)
	if err != nil {
		return models.User{}, err
	}

	if err :=  password.CompareHashAndPassword(user.Password, password); err != nil {
		return models.User{}, models.ErrInvalidCredentials // Use your defined error constant
	}

	return user, nil
}
