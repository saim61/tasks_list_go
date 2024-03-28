package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/saim61/tasks_list_go/tasks"
	"github.com/saim61/tasks_list_go/utils"

	"github.com/saim61/tasks_list_go/db"
)

func instructions(w http.ResponseWriter) {
	io.WriteString(w, "*********************************************************\n")
	io.WriteString(w, "Welcome to tasks list. You can do the following here.\n")
	io.WriteString(w, "/tasks to view your tasks\n")
	io.WriteString(w, "/createTask to create a task\n")
	io.WriteString(w, "/editTask to edit a task\n")
	io.WriteString(w, "/deleteTask to delete a task\n")
	io.WriteString(w, "*********************************************************\n")
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	instructions(w)
	w.WriteHeader(http.StatusCreated)
}

func TasksList(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "These are your tasks!\n")
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	errorString, tasks := tasks.GetAllTasks(database)

	if tasks != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write(utils.CreatePayload(errorString, "Failed to get tasksssssss"))
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	t := tasks.NewTask()
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload("ID not provided in request", "Request body not parsed"))
		return
	}

	var payload []byte

	errorString, task, status := tasks.GetTask(t.Id, database)
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	} else {
		w.WriteHeader(http.StatusNotFound)
		payload = utils.CreatePayload(errorString, "Failed to get task")
	}

	w.Write(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	task := tasks.NewTask()
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil || !utils.CheckValidCreateRequest(task) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload("Invalid parameters", "Request body not parsed"))
		return
	}

	var payload []byte
	errorString, status := tasks.CreateTask(task, database)
	if status {
		w.WriteHeader(http.StatusCreated)
		payload = utils.CreatePayload("", "Successfully created task")
	} else {
		w.WriteHeader(http.StatusForbidden)
		payload = utils.CreatePayload(errorString, "Failed to create task")
	}

	w.Write(payload)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	var payload []byte
	task := tasks.NewTask()

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil || !utils.CheckValidEditTaskRequest(task) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload("Invalid parameters", "Request body not parsed"))
		return
	}

	errorString, status := tasks.EditTask(task, database)
	if status {
		w.WriteHeader(http.StatusOK)
		payload = utils.CreatePayload("", "Successfully edited task")
	} else {
		w.WriteHeader(http.StatusForbidden)
		payload = utils.CreatePayload(errorString, "Failed to edit task")
	}

	w.Write(payload)
}

func EditTaskStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	var payload []byte
	task := tasks.NewTask()

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil || !utils.CheckValidEditTaskStatusRequest(task) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload("Invalid parameters", "Request body not parsed"))
		return
	}

	errorString, status := tasks.EditTaskStatus(task, database)
	if status {
		w.WriteHeader(http.StatusOK)
		payload = utils.CreatePayload("", "Successfully edited task status")
	} else {
		w.WriteHeader(http.StatusForbidden)
		payload = utils.CreatePayload(errorString, "Failed to edit task status")
	}

	w.Write(payload)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	t := tasks.NewTask()
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload("ID not provided in request", "Request body not parsed"))
		return
	}

	var payload []byte
	errorString, status := tasks.DeleteTask(t.Id, database)
	if status {
		w.WriteHeader(http.StatusOK)
		payload = utils.CreatePayload("", "Successfully deleted task")
	} else {
		w.WriteHeader(http.StatusForbidden)
		payload = utils.CreatePayload(errorString, "Failed to deleted task")
	}

	w.Write(payload)
}
