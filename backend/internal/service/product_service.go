package service

import (
	"context"

	"github.com/squid3rd/fiber-store-admin/internal/model"
	"github.com/squid3rd/fiber-store-admin/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepo
}

func NewProductService(repo *repository.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) FindAll(ctx context.Context, page, limit int64) ([]model.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	return s.repo.FindAll(ctx, page, limit)
}

func (s *ProductService) Create(ctx context.Context, in model.CreateProductInput) (*model.Product, error) {
	if in.Status == "" {
		in.Status = "active"
	}

	return s.repo.Create(ctx, in)
}
