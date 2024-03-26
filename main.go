package main

import (
	"net/http"
	"tasks_list_go/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", routes.Homepage)
	api.HandleFunc("/", routes.Homepage)
	api.HandleFunc("/tasks", routes.TasksList).Methods(http.MethodGet)
	api.HandleFunc("/task", routes.GetTask).Methods(http.MethodGet)
	api.HandleFunc("/createTask", routes.CreateTask).Methods(http.MethodPost)
	api.HandleFunc("/editTask", routes.EditTask).Methods(http.MethodPost)
	api.HandleFunc("/deleteTask", routes.DeleteTask).Methods(http.MethodDelete)

	http.ListenAndServe(":8000", router)
}
