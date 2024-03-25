package main

import (
	"fmt"
	"net/http"

	"tasks_list_go/routes"

	"github.com/gorilla/mux"
)

func instructions() {
	fmt.Println("*********************************************************")
	fmt.Println("Welcome to tasks list. You can do the following here.")
	fmt.Println("/tasks to view your tasks")
	fmt.Println("/createTask to create a task")
	fmt.Println("/editTask to edit a task")
	fmt.Println("/deletetTask to delete a task")
	fmt.Println("*********************************************************")
}

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", routes.Homepage)
	api.HandleFunc("/", routes.Homepage)
	api.HandleFunc("/tasks", routes.TasksList).Methods(http.MethodGet)
	api.HandleFunc("/createTask", routes.CreateTask).Methods(http.MethodPost)
	api.HandleFunc("/editTask", routes.EditTask).Methods(http.MethodPost)
	api.HandleFunc("/deleteTask", routes.DeleteTask).Methods(http.MethodDelete)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println("The server has now started! Please continue with the following paths to progress.")
		instructions()
	}
}
