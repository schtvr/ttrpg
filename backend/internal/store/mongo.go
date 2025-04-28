package store

import (
	"context"
	"log"
	"sync"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
	// uri         = "mongodb://root:example@localhost:27017/?authSource=admin"
)

func InitMongo(uri string) {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// Connect to MongoDB
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalf("Mongo connection failed: %v", err)
		}
		// Optional ping
		if err := client.Ping(ctx, nil); err != nil {
			log.Fatalf("Mongo ping failed: %v", err)
		}

		mongoClient = client
		log.Println("Mongo connected")
	})
}

func GetMongoClient() *mongo.Client {
	if mongoClient == nil {
		log.Fatal("Mongo client not initialized. Call InitMongo() first.")
	}
	return mongoClient
}
