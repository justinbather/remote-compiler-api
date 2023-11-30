package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var client *mongo.Client

func Connect() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading env vars")
		log.Fatal(err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err = mongo.Connect(context.Background(), clientOptions)
	return err

}

func GetClient() *mongo.Client {
	return client
}
