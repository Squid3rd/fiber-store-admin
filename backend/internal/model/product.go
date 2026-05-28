package model

import "time"

type Product struct {
	ID        string    `bson:"_id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Price     float64   `bson:"price" json:"price"`
	Stock     int       `bson:"stock" json:"stock"`
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type CreateProductInput struct {
	Name   string  `json:"name" validate:"required"`
	Price  float64 `json:"price" validate:"required,min=0"`
	Stock  int     `json:"stock" validate:"required,min=0"`
	Status string  `json:"status" validate:"required,oneof=active inactive"`
}

type UpdateProductInput struct {
	Name   string  `json:"name" validate:"required"`
	Price  float64 `json:"price" validate:"required,min=0"`
	Stock  int     `json:"stock" validate:"required,min=0"`
	Status string  `json:"status" validate:"required,oneof=active inactive"`
}

type ProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
