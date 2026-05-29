package repository

import (
	"context"
	"time"

	"github.com/squid3rd/fiber-store-admin/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct {
	db *mongo.Collection
}

func NewAuthRepo(db *mongo.Collection) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(ctx context.Context, in model.CreateUserInput) (*model.User, error) {
	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	record := model.UserRecord{
		User: model.User{
			ID:        primitive.NewObjectID().Hex(),
			Name:      in.Name,
			Email:     in.Email,
			CreatedAt: now,
			UpdatedAt: now,
		},
		PasswordHash: string(hashedPassword),
	}

	_, err = r.db.InsertOne(ctx, record)

	if err != nil {
		return nil, err
	}

	return &record.User, nil
}

func (r *AuthRepo) FindUserByEmail(ctx context.Context, email string) (*model.UserRecord, error) {
	var record model.UserRecord

	err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}
