package service

import (
	"context"

	"github.com/squid3rd/fiber-store-admin/internal/model"
	"github.com/squid3rd/fiber-store-admin/internal/repository"
)

type AuthService struct {
	repo *repository.AuthRepo
}

func NewAuthService(repo *repository.AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, in model.CreateUserInput) (*model.User, error) {
	return s.repo.CreateUser(ctx, in)
}

func (s *AuthService) FindUserByEmail(ctx context.Context, email string) (*model.UserRecord, error) {
	return s.repo.FindUserByEmail(ctx, email)
}
