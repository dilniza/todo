package service

import (
	"context"
	"todo/api/models"
	"todo/storage/mongodb"
)

type TaskService interface {
	CreateTask(ctx context.Context, req models.CreateTask) (models.Task, error)
	GetTaskByID(ctx context.Context, id string) (models.Task, error)
	UpdateTask(ctx context.Context, req models.UpdateTask) (models.Task, error)
	DeleteTask(ctx context.Context, id string) error
	ListTasks(ctx context.Context, taskListID string, page, limit uint64) ([]models.Task, int64, error)
}

type taskService struct {
	repo mongodb.TaskRepo
}

func NewTaskService(repo mongodb.TaskRepo) TaskService {
	return &taskService{repo: repo}
}

func (ts *taskService) CreateTask(ctx context.Context, req models.CreateTask) (models.Task, error) {
	return ts.repo.CreateTask(ctx, req)
}

func (ts *taskService) GetTaskByID(ctx context.Context, id string) (models.Task, error) {
	return ts.repo.GetTask(ctx, id)
}

func (ts *taskService) UpdateTask(ctx context.Context, req models.UpdateTask) (models.Task, error) {
	return ts.repo.UpdateTask(ctx, req)
}

func (ts *taskService) DeleteTask(ctx context.Context, id string) error {
	return ts.repo.DeleteTask(ctx, id)
}

func (ts *taskService) ListTasks(ctx context.Context, taskListID string, page, limit uint64) ([]models.Task, int64, error) {
	return ts.repo.GetAllTasks(ctx, taskListID, page, limit)
}
