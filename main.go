package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-compiler-api/db"
	"go-compiler-api/handlers"
	"log"
	"net/http"
)

func main() {

	if err := db.Connect(); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/jobs", handlers.GetAllCompileJobs).Methods("GET")
	r.HandleFunc("/jobs", handlers.CreateCompileJob).Methods("POST")
	r.HandleFunc("/jobs/{jobId}", handlers.DeleteCompileJob).Methods("DELETE")
	r.HandleFunc("/jobs/{jobId}", handlers.GetOneCompileJob).Methods("GET")
	r.HandleFunc("/jobs/{jobId}", handlers.UpdateCompileJob).Methods("PUT", "PATCH")

	fmt.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("error starting server", err)
	}

}
