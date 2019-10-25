package db
import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
)


func SetupDB() (*mongo.Client, error) {
	// Set client options
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
		log.Fatal(err)
		return nil, err
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
		log.Fatal(err)
		return nil, err
}

fmt.Println("Connected to MongoDB!")
return client, nil
}