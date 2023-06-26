// doc-extractor/pkg/repositories/setup.go

package repositories

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	db         *mongo.Database
	clientOnce sync.Once
)

// GetMongoClient returns a MongoDB client instance
func GetMongoClient() *mongo.Client {
	clientOnce.Do(func() {
		_ = ConnectMongoDB()
	})
	return client
}

// ConnectMongoDB connects to MongoDB and initializes the client and database
func ConnectMongoDB() error {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")

	// Assign the created client and database to the package-level variables
	client = mongoClient
	db = client.Database("your_database_name")

	return nil
}

// GetMongoDB returns the MongoDB database instance
func GetMongoDB() *mongo.Database {
	return db
}

// CloseMongoClient closes the MongoDB client connection
func CloseMongoClient() error {
	if client != nil {
		err := client.Disconnect(context.TODO())
		if err != nil {
			return err
		}
		log.Println("MongoDB client disconnected.")
	}
	return nil
}
