package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Label struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" bson:"name"`
	Color     string             `json:"color" bson:"color"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

type CreateLabel struct {
	UserID primitive.ObjectID `json:"user_id"`
	Name   string             `json:"name"`
	Color  string             `json:"color"`
}

type UpdateLabel struct {
	ID    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Color string             `json:"color"`
}

type GetLabel struct {
	ID     primitive.ObjectID `json:"id"`
	UserID primitive.ObjectID `json:"user_id"`
	Name   string             `json:"name"`
	Color  string             `json:"color"`
}
