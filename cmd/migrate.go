package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// runMigrations performs database migrations
func runMigrations(dbURI string) error {
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	database := client.Database("todoapp")

	usersCollection := database.Collection("users")

	_, err = usersCollection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("failed to create index: %v", err)
	}

	fmt.Println("Migrations completed.")
	return nil
}
