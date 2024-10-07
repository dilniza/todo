package mongodb

import (
	"context"
	"fmt"
	"todo/api/models"
	"todo/pkg"
	"todo/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LabelRepo struct {
	collection *mongo.Collection
	logger     logger.ILogger
}

// NewLabelRepo initializes a new LabelRepo with a MongoDB collection and logger.
func NewLabelRepo(db *mongo.Database, log logger.ILogger) *LabelRepo {
	return &LabelRepo{
		collection: db.Collection("labels"),
		logger:     log,
	}
}

// CreateLabel creates a new label in the database.
func (lr *LabelRepo) CreateLabel(ctx context.Context, req models.CreateLabel) (models.Label, error) {
	label := models.Label{
		Id:     pkg.GenerateUUID(),
		Name:   req.Name,
		UserID: req.UserID,
	}

	_, err := lr.collection.InsertOne(ctx, label)
	if err != nil {
		lr.logger.Error("error while creating label in db", logger.Error(err))
		return models.Label{}, fmt.Errorf("error while creating label: %w", err)
	}
	return label, nil
}

// GetLabel retrieves a label by its ID.
func (lr *LabelRepo) GetLabel(ctx context.Context, labelID string) (models.Label, error) {
	var label models.Label
	err := lr.collection.FindOne(ctx, bson.M{"id": labelID}).Decode(&label)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Label{}, fmt.Errorf("label not found: %w", err)
		}
		lr.logger.Error("error while getting label from db", logger.Error(err))
		return models.Label{}, fmt.Errorf("error while getting label: %w", err)
	}
	return label, nil
}

// UpdateLabel updates the label information.
func (lr *LabelRepo) UpdateLabel(ctx context.Context, req models.UpdateLabel) (models.Label, error) {
	filter := bson.M{"id": req.Id}
	update := bson.M{"$set": bson.M{"name": req.Name}}

	_, err := lr.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		lr.logger.Error("error while updating label in db", logger.Error(err))
		return models.Label{}, fmt.Errorf("error while updating label: %w", err)
	}

	return models.Label{
		Id:     req.Id,
		Name:   req.Name,
		UserID: req.UserID,
	}, nil
}

// DeleteLabel removes a label by its ID.
func (lr *LabelRepo) DeleteLabel(ctx context.Context, labelID string) error {
	_, err := lr.collection.DeleteOne(ctx, bson.M{"id": labelID})
	if err != nil {
		lr.logger.Error("error while deleting label from db", logger.Error(err))
		return fmt.Errorf("error while deleting label: %w", err)
	}
	return nil
}

// GetAllLabels retrieves all labels for a specific user with pagination support.
func (lr *LabelRepo) GetAllLabels(ctx context.Context, userID string, search string, page, limit uint64) ([]models.Label, int64, error) {
	var labels []models.Label
	count, err := lr.collection.CountDocuments(ctx, bson.M{"user_id": userID, "name": bson.M{"$regex": search, "$options": "i"}})
	if err != nil {
		lr.logger.Error("error while counting labels in db", logger.Error(err))
		return nil, 0, fmt.Errorf("error while counting labels: %w", err)
	}

	cursor, err := lr.collection.Find(ctx, bson.M{"user_id": userID, "name": bson.M{"$regex": search, "$options": "i"}}, options.Find().SetSkip((page-1)*limit).SetLimit(limit))
	if err != nil {
		lr.logger.Error("error while getting labels from db", logger.Error(err))
		return nil, 0, fmt.Errorf("error while getting labels: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var label models.Label
		if err := cursor.Decode(&label); err != nil {
			lr.logger.Error("error while decoding label in db", logger.Error(err))
			return nil, 0, fmt.Errorf("error while decoding label: %w", err)
		}
		labels = append(labels, label)
	}

	return labels, count, nil
}
