package utils

import (
	"fmt"

	"github.com/saim61/tasks_list_go/tasks"
)

func CreatePayload(errorString string, message string) []byte {
	if errorString != "" {
		return []byte(fmt.Sprintf(`{"message": %s, "error": %s}`, message, errorString))
	}

	return []byte(fmt.Sprintf(`{"message": %s}`, message))
}

func CheckValidEditTaskRequest(task tasks.Task) bool {
	if task.Id == -1 {
		return false
	} else if task.Title == "" || task.Description == "" || task.Status == "" {
		return false
	}
	return true
}

func CheckValidCreateRequest(task tasks.Task) bool {
	if task.Title == "" || task.Description == "" || task.Status == "" {
		return false
	}
	return true
}

func CheckValidEditTaskStatusRequest(task tasks.Task) bool {
	if task.Id == -1 {
		return false
	} else if task.Status == "" {
		return false
	}
	return true
}
