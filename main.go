package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := connectToMongoDB()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Now you can perform operations with the MongoDB client
}

func connectToMongoDB() (*mongo.Client, error) {
	// Set client options
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB successfully!")
	return client, nil
}
