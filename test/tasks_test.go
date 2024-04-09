package test

import (
	"testing"

	"github.com/saim61/tasks_list_go/db"
	"github.com/saim61/tasks_list_go/tasks"
	"github.com/stretchr/testify/assert"
)

// Passing scenario for get all tasks
func TestGetAllTasks(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()
	_, _, allTasks := tasks.GetAllUserTasks(1, database)

	assert.NotEmpty(t, allTasks)
}

// Passing scenario for create task
func TestCreateTask(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()
	task := tasks.CreateTaskRequest{
		Title:       "Demo title",
		Description: "Demo Description",
		Status:      "open",
	}

	_, _, status := tasks.CreateTask(task, 1, database)

	assert.Equal(t, status, true)
}

// Passing scenario for get task by id
func TestGetTaskById(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()
	_, _, allTasks := tasks.GetAllUserTasks(1, database)

	assert.NotEmpty(t, allTasks)

	_, _, task, status := tasks.GetTask(allTasks[0].Id, 1, database)

	assert.Equal(t, status, true)
	assert.NotEmpty(t, task)
}

// Passing scenario for delete task
func TestDeleteTask(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()
	_, _, allTasks := tasks.GetAllUserTasks(1, database)

	assert.NotEmpty(t, allTasks)

	_, _, status := tasks.DeleteTask(allTasks[0].Id, 1, database)

	assert.Equal(t, status, true)
}

// Passing scenario for edit task
func TestEditTask(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()
	_, _, allTasks := tasks.GetAllUserTasks(1, database)

	assert.NotEmpty(t, allTasks)

	task := tasks.Task{
		Id:          allTasks[0].Id,
		Title:       "Edited title",
		Description: "Edited Description",
		Status:      "closed",
	}

	_, _, status := tasks.EditTask(task, 1, database)

	assert.Equal(t, status, status)
}

// Passing scenario for edit task status
func TestEditTaskStatus(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()
	_, _, allTasks := tasks.GetAllUserTasks(1, database)

	assert.NotEmpty(t, allTasks)

	task := tasks.EditTaskStatusRequest{
		Id:     allTasks[0].Id,
		Status: "anything",
	}

	_, _, status := tasks.EditTaskStatus(task, 1, database)

	assert.Equal(t, status, true)
}

// Fail scenario for create task
// A task cannot have empty values
func TestFailCreateTask(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()

	task := tasks.CreateTaskRequest{
		Title:       "",
		Description: "",
		Status:      "",
	}

	_, _, status := tasks.CreateTask(task, 1, database)

	assert.Equal(t, status, false)
}

// Fail scenario for edit task
// Id cannot be -1 or 0
// Title, Description or Status cannot be empty
func TestFailEditTask(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()

	task := tasks.Task{
		Id:          -1,
		Title:       "",
		Description: "",
		Status:      "",
	}

	_, _, status := tasks.EditTask(task, 1, database)

	assert.Equal(t, status, false)
}

// Fail scenario for edit task status
// Id cannot be -1 or 0
// Status cannot be empty
func TestFailEditTaskStatus(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()

	task := tasks.EditTaskStatusRequest{
		Id:     -1,
		Status: "",
	}

	_, _, status := tasks.EditTaskStatus(task, 1, database)

	assert.Equal(t, status, false)
}

// Fail scenario for delete task
// Should have an ID parameter in request with a valid ID
func TestFailDeleteTask(t *testing.T) {
	database := db.GetDatabaseObject("test")
	defer database.Close()

	// Invalid task ID request. Task ID should be greater than 1
	_, _, status := tasks.DeleteTask(-1, 1, database)

	assert.Equal(t, status, false)

	// Task with this ID doesnt exist in DB
	_, _, status = tasks.DeleteTask(9999999, 1, database)

	assert.Equal(t, status, false)
}
