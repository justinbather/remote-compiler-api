package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-compiler-api/db"
	"go-compiler-api/models"
	"net/http"
	"time"
)

func GetAllCompileJobs(w http.ResponseWriter, r *http.Request) {
}

func GetOneCompileJob(w http.ResponseWriter, r *http.Request) {
}

func CreateCompileJob(w http.ResponseWriter, r *http.Request) {
	client := db.GetClient()

	coll := client.Database("test").Collection("public_compile")
	var newCompile models.CompileJob

	err := json.NewDecoder(r.Body).Decode(&newCompile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := coll.InsertOne(ctx, newCompile)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCompile)
	fmt.Println(res)
}

func DeleteCompileJob(w http.ResponseWriter, r *http.Request) {
}

func UpdateCompileJob(w http.ResponseWriter, r *http.Request) {
}
