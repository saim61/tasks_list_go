package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"tasks_list_go/db"
	"tasks_list_go/tasks"
	"tasks_list_go/utils"
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

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload(err.Error(), "Request body not parsed"))
		return
	}

	var payload []byte
	taskId, err := strconv.Atoi(r.FormValue("taskId"))

	if err == nil {
		errorString, task, status := tasks.GetTask(taskId, database)
		if status {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task)
		} else {
			w.WriteHeader(http.StatusNotFound)
			payload = utils.CreatePayload(errorString, "Failed to get task")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		payload = utils.CreatePayload("Please add a task ID in request", "Failed to get task")
	}

	w.Write(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload(err.Error(), "Request body not parsed"))
		return
	}

	var payload []byte
	if len(r.Form) != 3 {
		w.WriteHeader(http.StatusForbidden)
		payload = utils.CreatePayload("Missing parameters", "Failed to create task")
	} else {
		task := tasks.Task{
			Title:       r.FormValue("taskTitle"),
			Description: r.FormValue("taskDescription"),
			Status:      r.FormValue("taskStatus"),
		}

		errorString, status := tasks.CreateTask(task, database)
		if status {
			w.WriteHeader(http.StatusCreated)
			payload = utils.CreatePayload("", "Successfully created task")
		} else {
			w.WriteHeader(http.StatusForbidden)
			payload = utils.CreatePayload(errorString, "Failed to create task")
		}
	}

	w.Write(payload)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	var payload []byte
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload(err.Error(), "Request body not parsed"))
		return
	}

	// To ensure we have all the parameters required to edit the task
	if len(r.Form) != 4 {
		payload = utils.CreatePayload("Missing parameters", "Failed to edit task")
		w.WriteHeader(http.StatusForbidden)

	} else {
		taskId, _ := strconv.Atoi(r.FormValue("taskId"))
		task := tasks.Task{
			Id:          taskId,
			Title:       r.FormValue("taskTitle"),
			Description: r.FormValue("taskDescription"),
			Status:      r.FormValue("taskStatus"),
		}

		errorString, status := tasks.EditTask(task, database)
		if status {
			w.WriteHeader(http.StatusOK)
			payload = utils.CreatePayload("", "Successfully edited task")
		} else {
			w.WriteHeader(http.StatusForbidden)
			payload = utils.CreatePayload(errorString, "Failed to edit task")
		}
	}

	w.Write(payload)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.GetDatabaseObject()
	defer database.Close()

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.CreatePayload("", "Request body not parsed"))
		return
	}

	var payload []byte
	taskId, err := strconv.Atoi(r.FormValue("taskId"))

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		payload = utils.CreatePayload("Please add task ID", "Failed to delete task")
	} else {
		errorString, status := tasks.DeleteTask(taskId, database)
		if status {
			w.WriteHeader(http.StatusOK)
			payload = utils.CreatePayload("", "Successfully deleted task")
		} else {
			w.WriteHeader(http.StatusForbidden)
			payload = utils.CreatePayload(errorString, "Failed to deleted task")
		}
	}

	w.Write(payload)
}
