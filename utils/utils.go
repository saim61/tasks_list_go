package utils

import (
	"log"

	"github.com/saim61/tasks_list_go/tasks"
)

type ErrorResponse struct {
	ErrorCode   string
	ErrorString string
	Message     string
}

type SuccessResponse struct {
	Message string
}

type LoginSuccessResponse struct {
	Message string
	Token   string
}

func NewErrorResponse(errorCode string, errorString string, message string) ErrorResponse {
	return ErrorResponse{
		ErrorCode: errorCode, ErrorString: errorString, Message: message,
	}
}

func NewSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{
		Message: message,
	}
}

func NewLoginSuccessResponse(message string, token string) LoginSuccessResponse {
	return LoginSuccessResponse{
		Message: message,
		Token:   token,
	}
}

func IsValidIdForTask(id int) bool {
	return id > 0
}

func IsValidEditTask(task tasks.Task) bool {
	if task.Description == "" || task.Title == "" || task.Status == "" {
		return false
	}

	if !IsValidIdForTask(task.Id) {
		return false
	}
	return true
}

func IsValidEditTaskStatus(task tasks.EditTaskStatusRequest) bool {
	if task.Status == "" {
		return false
	}

	if !IsValidIdForTask(task.Id) {
		return false
	}
	return true
}

func IsValidCreateTask(task tasks.CreateTaskRequest) bool {
	if task.Description == "" || task.Title == "" || task.Status == "" {
		return false
	}
	return true
}

func PrintTask(task tasks.Task) {
	log.Println("Id: ", task.Id)
	log.Println("Title: ", task.Id)
	log.Println("Description: ", task.Id)
	log.Println("Status: ", task.Id)
}

func PrintTasks(tasks []tasks.Task) {
	for i := 0; i < len(tasks); i++ {
		PrintTask(tasks[i])
	}
}
