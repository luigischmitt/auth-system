package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Service struct {
	Database *mongo.Database
}

func New() *Service {
	// Load .env file from current directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Read environment variables inside the function
	dbURI := os.Getenv("MONGODB_URL")
	dbName := os.Getenv("MONGODB_DATABASE")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(dbURI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	return &Service{
		Database: client.Database(dbName),
	}
}
