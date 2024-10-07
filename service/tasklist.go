package service

import (
	"context"
	"todo/api/models"
	"todo/storage/mongodb"
)

type TaskListService interface {
	CreateTaskList(ctx context.Context, req models.CreateTaskList) (models.TaskList, error)
	GetTaskListByID(ctx context.Context, id string) (models.TaskList, error)
	UpdateTaskList(ctx context.Context, req models.UpdateTaskList) (models.TaskList, error)
	DeleteTaskList(ctx context.Context, id string) error
	ListTaskLists(ctx context.Context, userID string, page, limit uint64) ([]models.TaskList, int64, error)
}

type taskListService struct {
	repo mongodb.TaskListRepo
}

func NewTaskListService(repo mongodb.TaskListRepo) TaskListService {
	return &taskListService{repo: repo}
}

func (tls *taskListService) CreateTaskList(ctx context.Context, req models.CreateTaskList) (models.TaskList, error) {
	return tls.repo.CreateTaskList(ctx, req)
}

func (tls *taskListService) GetTaskListByID(ctx context.Context, id string) (models.TaskList, error) {
	return tls.repo.GetTaskList(ctx, id)
}

func (tls *taskListService) UpdateTaskList(ctx context.Context, req models.UpdateTaskList) (models.TaskList, error) {
	return tls.repo.UpdateTaskList(ctx, req)
}

func (tls *taskListService) DeleteTaskList(ctx context.Context, id string) error {
	return tls.repo.DeleteTaskList(ctx, id)
}

func (tls *taskListService) ListTaskLists(ctx context.Context, userID string, page, limit uint64) ([]models.TaskList, int64, error) {
	return tls.repo.GetAllTaskLists(ctx, userID, page, limit)
}
