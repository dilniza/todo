package mongodb

import (
	"context"
	"todo/api/models"
	"todo/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	db  *mongo.Database
	log logger.ILogger
}

func NewUserRepo(db *mongo.Database, log logger.ILogger) *UserRepo {
	return &UserRepo{db: db, log: log}
}

// CreateUser creates a new user in the database.
func (ur *UserRepo) CreateUser(ctx context.Context, req models.CreateUser) (models.User, error) {
	user := models.User{
		ID:    req.ID,
		Username:  req.Username,
		Email: req.Email,
	}

	_, err := ur.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		ur.log.Error("Error creating user: ", err)
		return models.User{}, err
	}

	return user, nil
}

// GetUser retrieves a user by ID.
func (ur *UserRepo) GetUser(ctx context.Context, userID string) (models.User, error) {
	var user models.User
	err := ur.db.Collection("users").FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, nil // Return an empty user if not found
		}
		ur.log.Error("Error retrieving user: ", err)
		return models.User{}, err
	}

	return user, nil
}

// UpdateUser updates an existing user.
func (ur *UserRepo) UpdateUser(ctx context.Context, req models.UpdateUser) (models.User, error) {
	filter := bson.M{"id": req.ID}
	update := bson.M{
		"$set": bson.M{
			"name":  req.Name,
			"email": req.Email,
		},
	}

	_, err := ur.db.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		ur.log.Error("Error updating user: ", err)
		return models.User{}, err
	}

	return ur.GetUser(ctx, req.ID)
}

// DeleteUser removes a user from the database.
func (ur *UserRepo) DeleteUser(ctx context.Context, userID string) error {
	_, err := ur.db.Collection("users").DeleteOne(ctx, bson.M{"id": userID})
	if err != nil {
		ur.log.Error("Error deleting user: ", err)
		return err
	}
	return nil
}

// GetAllUsers retrieves all users with pagination.
func (ur *UserRepo) GetAllUsers(ctx context.Context, page, limit uint64) ([]models.User, int64, error) {
	users := []models.User{}
	opts := options.Find().
		SetSkip((page - 1) * limit).
		SetLimit(limit)

	cursor, err := ur.db.Collection("users").Find(ctx, bson.M{}, opts)
	if err != nil {
		ur.log.Error("Error retrieving users: ", err)
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			ur.log.Error("Error decoding user: ", err)
			continue
		}
		users = append(users, user)
	}

	count, err := ur.db.Collection("users").CountDocuments(ctx, bson.M{})
	if err != nil {
		ur.log.Error("Error counting users: ", err)
		return nil, 0, err
	}

	return users, count, nil
}
