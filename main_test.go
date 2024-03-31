package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/saim61/tasks_list_go/routes"
	"github.com/saim61/tasks_list_go/tasks"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine = routes.SetupTestingRoutes()

// Passing scenario for get all tasks
func TestGetAllTasks(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var tasks []tasks.Task
	json.Unmarshal(w.Body.Bytes(), &tasks)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, tasks)
}

// Passing scenario for create task
func TestCreateTask(t *testing.T) {
	task := tasks.CreateTaskRequest{
		Title:       "Demo title",
		Description: "Demo Description",
		Status:      "open",
	}
	jsonTask, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/test/createTask", bytes.NewBuffer(jsonTask))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// Passing scenario for get task by id
func TestGetTaskById(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test/task/:id?id=1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Passing scenario for delete task
func TestDeleteTask(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/test/deleteTask/:id?id=5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Passing scenario for edit task
func TestEditTask(t *testing.T) {
	task := tasks.Task{
		Id:          1,
		Title:       "Edited title",
		Description: "Edited Description",
		Status:      "closed",
	}

	jsonEditTask, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/test/editTask", bytes.NewBuffer(jsonEditTask))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Passing scenario for edit task status
func TestEditTaskStatus(t *testing.T) {
	task := tasks.EditTaskStatusRequest{
		Id:     2,
		Status: "anything",
	}

	jsonEditTask, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/test/editTaskStatus", bytes.NewBuffer(jsonEditTask))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Fail scenario for create task
// A task cannot have empty values
func TestFailCreateTask(t *testing.T) {
	task := tasks.CreateTaskRequest{
		Title:       "",
		Description: "",
		Status:      "",
	}
	jsonTask, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/test/createTask", bytes.NewBuffer(jsonTask))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Fail scenario for edit task
// Id cannot be -1 or 0
// Title, Description or Status cannot be empty
func TestFailEditTask(t *testing.T) {
	task := tasks.Task{
		Id:          -1,
		Title:       "",
		Description: "",
		Status:      "",
	}

	jsonEditTask, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/test/editTask", bytes.NewBuffer(jsonEditTask))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Fail scenario for edit task status
// Id cannot be -1 or 0
// Status cannot be empty
func TestFailEditTaskStatus(t *testing.T) {
	task := tasks.EditTaskStatusRequest{
		Id:     -1,
		Status: "",
	}

	jsonEditTask, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/test/editTaskStatus", bytes.NewBuffer(jsonEditTask))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Fail scenario for delete task
// Should have an ID parameter in request with a valid ID
func TestFailDeleteTask(t *testing.T) {
	// Invalid ID request
	req, _ := http.NewRequest("DELETE", "/test/deleteTask/:id?id=-1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)

	// Task with this ID doesnt exist in DB
	req, _ = http.NewRequest("DELETE", "/test/deleteTask/:id?id=99999999", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
