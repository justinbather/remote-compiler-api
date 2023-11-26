package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

func Connect() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading env vars")
		log.Fatal(err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	fmt.Println(MONGO_URI)

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		fmt.Println("Error creating mongo client")
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting to mongo client")
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error pinging atlas cluster")
		log.Fatal(err)
	}
	return client

}
