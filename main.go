package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/justinbather/remote-compiler-api/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type CompileJob struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Category   string             `bson:"category"`
	Difficulty string             `bson:"difficulty"`
}

func getAllCompileJobs(w http.ResponseWriter, r *http.Request) {
}

func getOneCompileJob(w http.ResponseWriter, r *http.Request) {
}

func createCompileJob(w http.ResponseWriter, r *http.Request) {
}

func deleteCompileJob(w http.ResponseWriter, r *http.Request) {
}

func updateCompileJob(w http.ResponseWriter, r *http.Request) {
}

func main() {
	client := internal.Connect()

	coll := client.Database("test").Collection("problems")

	var result CompileJob
	filter := bson.D{{"title", "Two Sum"}}
	err = coll.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {

		fmt.Println("no documents found")
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	dbList, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error fetching list of database names")
		log.Fatal(err)
	}
	fmt.Println(dbList)

	r := mux.NewRouter()
	r.HandleFunc("/jobs", getAllCompileJobs).Methods("GET")
	r.HandleFunc("/jobs", createCompileJob).Methods("POST")
	r.HandleFunc("/jobs/{jobId}", deleteCompileJob).Methods("DELETE")
	r.HandleFunc("/jobs/{jobId}", getOneCompileJob).Methods("GET")
	r.HandleFunc("/jobs/{jobId}", updateCompileJob).Methods("PUT", "PATCH")

	fmt.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("error starting server", err)
	}

	defer client.Disconnect(ctx)
}
