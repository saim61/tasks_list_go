package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my homepage!\n")
	w.WriteHeader(http.StatusCreated)
}

func tasksList(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my tasks list!\n")
	w.WriteHeader(http.StatusCreated)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	taskId := r.FormValue("taskId")
	taskTitle := r.FormValue("taskTitle")
	taskDescription := r.FormValue("taskDescription")
	taskStatus := r.FormValue("taskStatus")

	payload := []byte(fmt.Sprintf(`{"message":"Successfully created task", "taskId":%s, "taskTitle":%s, "taskDescription":%s, "taskStatus":%s}`, taskId, taskTitle, taskDescription, taskStatus))

	w.WriteHeader(http.StatusCreated)
	w.Write(payload)
}

func editTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	taskId := r.FormValue("taskId")
	taskTitle := r.FormValue("taskTitle")
	taskDescription := r.FormValue("taskDescription")
	taskStatus := r.FormValue("taskStatus")

	payload := []byte(fmt.Sprintf(`{"message":"Successfully edited task", "taskId":%s, "taskTitle":%s, "taskDescription":%s, "taskStatus":%s}`, taskId, taskTitle, taskDescription, taskStatus))

	w.WriteHeader(http.StatusCreated)
	w.Write(payload)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	taskId := r.FormValue("taskId")

	payload := []byte(fmt.Sprintf(`{"message":"Successfully deleted task", "taskId":%s}`, taskId))

	w.WriteHeader(http.StatusCreated)
	w.Write(payload)
}

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", homepage).Methods(http.MethodGet)
	api.HandleFunc("/", homepage).Methods(http.MethodGet)
	api.HandleFunc("/tasks", tasksList).Methods(http.MethodGet)
	api.HandleFunc("/createTask", createTask).Methods(http.MethodPost)
	api.HandleFunc("/editTask", editTask).Methods(http.MethodPost)
	api.HandleFunc("/deleteTask", deleteTask).Methods(http.MethodDelete)

	http.ListenAndServe(":8000", router)
}
