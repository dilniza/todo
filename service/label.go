package service

import (
	"context"
	"todo/api/models"
	"todo/storage/mongodb"
)

type LabelService interface {
	CreateLabel(ctx context.Context, req models.CreateLabel) (models.Label, error)
	GetLabelByID(ctx context.Context, id string) (models.Label, error)
	UpdateLabel(ctx context.Context, req models.UpdateLabel) (models.Label, error)
	DeleteLabel(ctx context.Context, id string) error
	ListLabels(ctx context.Context, page, limit uint64) ([]models.Label, int64, error)
}

type labelService struct {
	repo mongodb.LabelRepo
}

func NewLabelService(repo mongodb.LabelRepo) LabelService {
	return &labelService{repo: repo}
}

func (ls *labelService) CreateLabel(ctx context.Context, req models.CreateLabel) (models.Label, error) {
	return ls.repo.CreateLabel(ctx, req)
}

func (ls *labelService) GetLabelByID(ctx context.Context, id string) (models.Label, error) {
	return ls.repo.GetLabel(ctx, id)
}

func (ls *labelService) UpdateLabel(ctx context.Context, req models.UpdateLabel) (models.Label, error) {
	return ls.repo.UpdateLabel(ctx, req)
}

func (ls *labelService) DeleteLabel(ctx context.Context, id string) error {
	return ls.repo.DeleteLabel(ctx, id)
}

func (ls *labelService) ListLabels(ctx context.Context, page, limit uint64) ([]models.Label, int64, error) {
	return ls.repo.GetAllLabels(ctx, page, limit)
}
