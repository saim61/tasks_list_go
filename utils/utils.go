package utils

import "github.com/saim61/tasks_list_go/tasks"

type ErrorResponse struct {
	ErrorCode   string
	ErrorString string
	Message     string
}

type SuccessResponse struct {
	Message string
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

func IsValidIdForTask(id int) bool {
	if id > 0 {
		return true
	}
	return false
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
