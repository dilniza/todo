package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskList struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

type CreateTaskList struct {
	UserID      primitive.ObjectID `json:"user_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
}

type UpdateTaskList struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
}

type GetTaskList struct {
	ID          primitive.ObjectID `json:"id"`
	UserID      primitive.ObjectID `json:"user_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
}
