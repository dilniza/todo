package mongodb

import (
	"todo/pkg/logger"
	"todo/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

// Storage struct contains all repositories for MongoDB.
type Storage struct {
	UserRepo     UserRepo
	LabelRepo    LabelRepo
	TaskRepo     TaskRepo
	TaskListRepo TaskListRepo
}

// NewStorage initializes a new Storage struct with the provided MongoDB database and logger.
func NewStorage(db *mongo.Database, log logger.ILogger) *Storage {
	return &Storage{
		UserRepo:     *NewUserRepo(db, log),
		LabelRepo:    *NewLabelRepo(db, log),
		TaskRepo:     *NewTaskRepo(db, log),
		TaskListRepo: *NewTaskListRepo(db, log),
	}
}

// Interface for User, Label, Task, and TaskList repositories.
var (
	_ storage.UserStorage     = &UserRepo{}
	_ storage.LabelStorage    = &LabelRepo{}
	_ storage.TaskStorage     = &TaskRepo{}
	_ storage.TaskListStorage = &TaskListRepo{}
)
