package repository

import (
	"context"
	"time"

	"github.com/squid3rd/fiber-store-admin/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepo struct {
	db *mongo.Collection
}

func NewProductRepo(db *mongo.Collection) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) FindAll(ctx context.Context, page, limit int64) ([]model.Product, int64, error) {
	skip := (page - 1) * limit
	opts := options.Find().SetSkip(skip).SetLimit(limit)

	cur, err := r.db.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}

	defer cur.Close(ctx)

	var items []model.Product
	if err := cur.All(ctx, &items); err != nil {
		return nil, 0, err
	}

	total, err := r.db.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return items, total, err

}

func (r *ProductRepo) Create(ctx context.Context, in model.CreateProductInput) (*model.Product, error) {
	now := time.Now()

	p := model.Product{
		ID:        primitive.NewObjectID().Hex(),
		Name:      in.Name,
		Price:     in.Price,
		Stock:     in.Stock,
		Status:    in.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := r.db.InsertOne(ctx, p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
