package storage

import (
	"context"
	"todo/api/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	userCollection     *mongo.Collection
	taskCollection     *mongo.Collection
	taskListCollection *mongo.Collection
	labelCollection    *mongo.Collection
}

// NewStorage initializes a new MongoDB instance with the required collections.
func NewStorage(dbURI string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ensure MongoDB connection is established
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	db := client.Database("todoapp")
	return &MongoDB{
		userCollection:     db.Collection("users"),
		taskCollection:     db.Collection("tasks"),
		taskListCollection: db.Collection("task_lists"),
		labelCollection:    db.Collection("labels"),
	}, nil
}

// UserStorage defines the methods for user storage operations.
type UserStorage interface {
	CreateUser(ctx context.Context, req models.CreateUser) (models.User, error)
	GetUser(ctx context.Context, userID string) (models.User, error)
	UpdateUser(ctx context.Context, req models.UpdateUser) (models.User, error)
	DeleteUser(ctx context.Context, userID string) error
	GetAllUsers(ctx context.Context, page, limit uint64) ([]models.User, int64, error)
}

// LabelStorage defines the methods for label storage operations.
type LabelStorage interface {
	CreateLabel(ctx context.Context, req models.CreateLabel) (models.Label, error)
	GetLabel(ctx context.Context, labelID string) (models.Label, error)
	UpdateLabel(ctx context.Context, req models.UpdateLabel) (models.Label, error)
	DeleteLabel(ctx context.Context, labelID string) error
	GetAllLabels(ctx context.Context, userID string, search string, page, limit uint64) ([]models.Label, int64, error)
}

// TaskStorage defines the methods for task storage operations.
type TaskStorage interface {
	CreateTask(ctx context.Context, req models.CreateTask) (models.Task, error)
	GetTask(ctx context.Context, taskID string) (models.Task, error)
	UpdateTask(ctx context.Context, req models.UpdateTask) (models.Task, error)
	DeleteTask(ctx context.Context, taskID string) error
	GetAllTasks(ctx context.Context, userID string, search string, page, limit uint64) ([]models.Task, int64, error)
}

// TaskListStorage defines the methods for task list storage operations.
type TaskListStorage interface {
	CreateTaskList(ctx context.Context, req models.CreateTaskList) (models.TaskList, error)
	GetTaskList(ctx context.Context, taskListID string) (models.TaskList, error)
	UpdateTaskList(ctx context.Context, req models.UpdateTaskList) (models.TaskList, error)
	DeleteTaskList(ctx context.Context, taskListID string) error
	GetAllTaskLists(ctx context.Context, userID string, page, limit uint64) ([]models.TaskList, int64, error)
}
