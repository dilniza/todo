package service

import (
	"todo/storage/mongodb"
)

type Service struct {
	AuthService     AuthService
	UserService     UserService
	TaskService     TaskService
	TaskListService TaskListService
	LabelService    LabelService
}

func NewService(userRepo mongodb.UserRepo, taskRepo mongodb.TaskRepo, taskListRepo mongodb.TaskListRepo, labelRepo mongodb.LabelRepo) *Service {
	return &Service{
		AuthService:     NewAuthService(userRepo),
		UserService:     NewUserService(userRepo),
		TaskService:     NewTaskService(taskRepo),
		TaskListService: NewTaskListService(taskListRepo),
		LabelService:    NewLabelService(labelRepo),
	}
}
