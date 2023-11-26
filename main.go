package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CompileJob struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Status string `json:"status"`
}

var CompileJobs = []CompileJob{
	{ID: 1, Code: "test", Status: "pending"},
	{ID: 2, Code: "testing", Status: "complete"},
}

func getAllCompileJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Returning all jobs")
	json.NewEncoder(w).Encode(CompileJobs)
}

func getOneCompileJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := strconv.Atoi(params["jobId"])
	if err != nil {
		w.WriteHeader(400)
	}

	for _, job := range CompileJobs {
		if job.ID == idParam {
			json.NewEncoder(w).Encode(job)
		}
	}
	w.WriteHeader(400)
}

func createCompileJob(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}

	var newJob CompileJob
	json.Unmarshal(data, &newJob)

	CompileJobs = append(CompileJobs, newJob)
	json.NewEncoder(w).Encode(newJob)
}

func deleteCompileJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := strconv.Atoi(params["jobId"])
	if err != nil {
		w.WriteHeader(400)
		return
	}

	for i, job := range CompileJobs {
		if job.ID == idParam {
			CompileJobs = append(CompileJobs[:i], CompileJobs[i+1:]...)
			w.WriteHeader(204)
			return
		}
	}

	w.WriteHeader(400)
}

func updateCompileJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := strconv.Atoi(params["jobId"])
	if err != nil {
		w.WriteHeader(400)
		return
	}

	data, err := io.ReadAll((r.Body))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var update CompileJob
	json.Unmarshal(data, &update)

	for i, job := range CompileJobs {
		if job.ID == idParam {
			CompileJobs = append(CompileJobs[:i], CompileJobs[i+1:]...)
			CompileJobs = append(CompileJobs, update)
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(update)
			return
		}
	}
}

func main() {

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

}
