package service

import (
	"context"
	"todo/api/models"
	"todo/storage/mongodb"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUser) (models.User, error)
	GetUserByID(ctx context.Context, id string) (models.User, error)
	UpdateUser(ctx context.Context, req models.UpdateUser) (models.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, page, limit uint64) ([]models.User, int64, error)
}

type userService struct {
	repo mongodb.UserRepo
}

func NewUserService(repo mongodb.UserRepo) UserService {
	return &userService{repo: repo}
}

func (us *userService) CreateUser(ctx context.Context, req models.CreateUser) (models.User, error) {
	return us.repo.CreateUser(ctx, req)
}

func (us *userService) GetUserByID(ctx context.Context, id string) (models.User, error) {
	return us.repo.GetUser(ctx, id)
}

func (us *userService) UpdateUser(ctx context.Context, req models.UpdateUser) (models.User, error) {
	return us.repo.UpdateUser(ctx, req)
}

func (us *userService) DeleteUser(ctx context.Context, id string) error {
	return us.repo.DeleteUser(ctx, id)
}

func (us *userService) ListUsers(ctx context.Context, page, limit uint64) ([]models.User, int64, error) {
	return us.repo.GetAllUsers(ctx, page, limit)
}
