package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"` // MongoDB ObjectID
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUser struct {
	ID       primitive.ObjectID `json:"id"` // Returning the ID as an ObjectID to maintain consistency
	Username string             `json:"username"`
	Email    string             `json:"email"`
}

type UpdateUser struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
}

type GetAllUsersRequest struct {
	Search string `json:"search"`
	Page   int64  `json:"page"`
	Limit  int64  `json:"limit"`
}

type GetAllUsersResponse struct {
	Users []User `json:"users"` // Corrected "customers" to "users"
	Count int64  `json:"count"`
}

type GetTaskListResponse struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Tasks       []GetTaskResponse `json:"tasks"`
}

type GetTaskResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     time.Time          `json:"due_date"`
	Completed   bool               `json:"completed"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type GetUserTaskListsResponse struct {
	UserTaskLists []GetTaskListResponse `json:"user_tasklists"`
	Count         int64                 `json:"count"`
}
