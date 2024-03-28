package main

import (
	"fmt"
	"net/http"

	"github.com/saim61/tasks_list_go/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", routes.Homepage)
	api.HandleFunc("/", routes.Homepage)
	api.HandleFunc("/tasks", routes.TasksList).Methods(http.MethodGet)
	api.HandleFunc("/task", routes.GetTask).Methods(http.MethodGet)
	api.HandleFunc("/createTask", routes.CreateTask).Methods(http.MethodPost)
	api.HandleFunc("/editTask", routes.EditTask).Methods(http.MethodPost)
	api.HandleFunc("/editTaskStatus", routes.EditTaskStatus).Methods(http.MethodPost)
	api.HandleFunc("/deleteTask", routes.DeleteTask).Methods(http.MethodDelete)

	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", router)
}
