// mongo.go

package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client      *mongo.Client
	playersColl *mongo.Collection
)

const (
	databaseName   = "nba-stats"
	collectionName = "players"
)

func InitMongo(ctx context.Context) {
	mongoURI := os.Getenv("MONGO_DB_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_DB_URI not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	if err = c.Ping(ctx, nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	client = c
	playersColl = client.Database(databaseName).Collection(collectionName)
}
