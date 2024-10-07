package mongodb

import (
	"context"
	"todo/api/models"
	"todo/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewTaskRepo(db *mongo.Database, log logger.ILogger) *TaskRepo {
	return &TaskRepo{db: db, log: log}
}

// CreateTask creates a new task in the database.
func (tr *TaskRepo) CreateTask(ctx context.Context, req models.CreateTask) (models.Task, error) {
	task := models.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		UserID:      req.UserID,
		Status:      req.Status,
	}

	_, err := tr.db.Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		tr.log.Error("Error creating task: ", err)
		return models.Task{}, err
	}

	return task, nil
}

// GetTask retrieves a task by ID.
func (tr *TaskRepo) GetTask(ctx context.Context, taskID string) (models.Task, error) {
	var task models.Task
	err := tr.db.Collection("tasks").FindOne(ctx, bson.M{"id": taskID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, nil // Return an empty task if not found
		}
		tr.log.Error("Error retrieving task: ", err)
		return models.Task{}, err
	}

	return task, nil
}

// UpdateTask updates an existing task.
func (tr *TaskRepo) UpdateTask(ctx context.Context, req models.UpdateTask) (models.Task, error) {
	filter := bson.M{"id": req.ID}
	update := bson.M{
		"$set": bson.M{
			"title":       req.Title,
			"description": req.Description,
			"user_id":     req.UserID,
			"status":      req.Status,
		},
	}

	_, err := tr.db.Collection("tasks").UpdateOne(ctx, filter, update)
	if err != nil {
		tr.log.Error("Error updating task: ", err)
		return models.Task{}, err
	}

	return tr.GetTask(ctx, req.ID)
}

// DeleteTask removes a task from the database.
func (tr *TaskRepo) DeleteTask(ctx context.Context, taskID string) error {
	_, err := tr.db.Collection("tasks").DeleteOne(ctx, bson.M{"id": taskID})
	if err != nil {
		tr.log.Error("Error deleting task: ", err)
		return err
	}
	return nil
}

// GetAllTasks retrieves all tasks for a user with pagination.
func (tr *TaskRepo) GetAllTasks(ctx context.Context, userID string, search string, page, limit uint64) ([]models.Task, int64, error) {
	tasks := []models.Task{}
	filter := bson.M{"user_id": userID}

	if search != "" {
		filter["title"] = bson.M{"$regex": search, "$options": "i"}
	}

	opts := options.Find().
		SetSkip((page - 1) * limit).
		SetLimit(limit)

	cursor, err := tr.db.Collection("tasks").Find(ctx, filter, opts)
	if err != nil {
		tr.log.Error("Error retrieving tasks: ", err)
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			tr.log.Error("Error decoding task: ", err)
			continue
		}
		tasks = append(tasks, task)
	}

	count, err := tr.db.Collection("tasks").CountDocuments(ctx, filter)
	if err != nil {
		tr.log.Error("Error counting tasks: ", err)
		return nil, 0, err
	}

	return tasks, count, nil
}
