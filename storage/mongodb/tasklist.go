package mongodb

import (
	"context"
	"todo/api/models"
	"todo/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskListRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewTaskListRepo(db *mongo.Database, log logger.ILogger) *TaskListRepo {
	return &TaskListRepo{db: db, log: log}
}

// CreateTaskList creates a new task list in the database.
func (tlr *TaskListRepo) CreateTaskList(ctx context.Context, req models.CreateTaskList) (models.TaskList, error) {
	taskList := models.TaskList{
		ID:     req.ID,
		Title:  req.Title,
		UserID: req.UserID,
	}

	_, err := tlr.db.Collection("task_lists").InsertOne(ctx, taskList)
	if err != nil {
		tlr.log.Error("Error creating task list: ", err)
		return models.TaskList{}, err
	}

	return taskList, nil
}

// GetTaskList retrieves a task list by ID.
func (tlr *TaskListRepo) GetTaskList(ctx context.Context, taskListID string) (models.TaskList, error) {
	var taskList models.TaskList
	err := tlr.db.Collection("task_lists").FindOne(ctx, bson.M{"id": taskListID}).Decode(&taskList)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.TaskList{}, nil // Return an empty task list
		}
		tlr.log.Error("Error retrieving task list: ", err)
		return models.TaskList{}, err
	}

	return taskList, nil
}

// UpdateTaskList updates an existing task list.
func (tlr *TaskListRepo) UpdateTaskList(ctx context.Context, req models.UpdateTaskList) (models.TaskList, error) {
	filter := bson.M{"id": req.ID}
	update := bson.M{
		"$set": bson.M{
			"title": req.Title,
		},
	}

	_, err := tlr.db.Collection("task_lists").UpdateOne(ctx, filter, update)
	if err != nil {
		tlr.log.Error("Error updating task list: ", err)
		return models.TaskList{}, err
	}

	return tlr.GetTaskList(ctx, req.ID)
}

// DeleteTaskList removes a task list from the database.
func (tlr *TaskListRepo) DeleteTaskList(ctx context.Context, taskListID string) error {
	_, err := tlr.db.Collection("task_lists").DeleteOne(ctx, bson.M{"id": taskListID})
	if err != nil {
		tlr.log.Error("Error deleting task list: ", err)
		return err
	}
	return nil
}

// GetAllTaskLists retrieves all task lists for a user with pagination.
func (tlr *TaskListRepo) GetAllTaskLists(ctx context.Context, userID string, page, limit uint64) ([]models.TaskList, int64, error) {
	taskLists := []models.TaskList{}
	filter := bson.M{"user_id": userID}

	opts := options.Find().
		SetSkip((page - 1) * limit).
		SetLimit(limit)

	cursor, err := tlr.db.Collection("task_lists").Find(ctx, filter, opts)
	if err != nil {
		tlr.log.Error("Error retrieving task lists: ", err)
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var taskList models.TaskList
		if err := cursor.Decode(&taskList); err != nil {
			tlr.log.Error("Error decoding task list: ", err)
			continue
		}
		taskLists = append(taskLists, taskList)
	}

	count, err := tlr.db.Collection("task_lists").CountDocuments(ctx, filter)
	if err != nil {
		tlr.log.Error("Error counting task lists: ", err)
		return nil, 0, err
	}

	return taskLists, count, nil
}
