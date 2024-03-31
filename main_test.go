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

var router *gin.Engine = routes.SetupRouter()

func TestGetAllTasks(t *testing.T) {
	router.GET("/test/tasks", routes.TasksList)
	req, _ := http.NewRequest("GET", "/test/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var tasks []tasks.Task
	json.Unmarshal(w.Body.Bytes(), &tasks)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, tasks)
}

func TestCreateTask(t *testing.T) {
	router.POST("/test/createTask", routes.CreateTask)
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

func TestGetTaskById(t *testing.T) {
	router.GET("/test/task/:id", routes.GetTask)
	req, _ := http.NewRequest("GET", "/test/task/:id?id=2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	router.DELETE("/test/deleteTask/:id", routes.DeleteTask)
	req, _ := http.NewRequest("DELETE", "/test/deleteTask/:id?id=5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEditTask(t *testing.T) {
	router.POST("/test/editTask", routes.EditTask)
	task := tasks.Task{
		Id:          2,
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

func TestEditTaskStatus(t *testing.T) {
	router.POST("/test/editTaskStatus", routes.EditTaskStatus)
	task := tasks.EditTaskStatusRequest{
		Id:          2,
		Status:      "anything",
	}

	jsonEditTask, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/test/editTaskStatus", bytes.NewBuffer(jsonEditTask))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
