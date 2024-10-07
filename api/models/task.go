package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    TaskListID  primitive.ObjectID `json:"task_list_id" bson:"task_list_id"`
    Title       string             `json:"title" bson:"title"`
    Description string             `json:"description" bson:"description"`
    DueDate     time.Time          `json:"due_date" bson:"due_date"`
    Completed   bool               `json:"completed" bson:"completed"`
    CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
    UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

type CreateTask struct {
    TaskListID  primitive.ObjectID `json:"task_list_id"`
    Title       string             `json:"title"`
    Description string             `json:"description"`
    DueDate     time.Time          `json:"due_date"`
}

type UpdateTask struct {
    ID          primitive.ObjectID `json:"id"`
    Title       string             `json:"title"`
    Description string             `json:"description"`
    DueDate     time.Time          `json:"due_date"`
    Completed   bool               `json:"completed"`
}

type GetTask struct {
    ID          primitive.ObjectID `json:"id"`
    TaskListID  primitive.ObjectID `json:"task_list_id"`
    Title       string             `json:"title"`
    Description string             `json:"description"`
    DueDate     time.Time          `json:"due_date"`
    Completed   bool               `json:"completed"`
}
