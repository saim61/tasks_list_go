package routes

import (
	"fmt"
	"io"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my homepage!\n")
	w.WriteHeader(http.StatusCreated)
}

func TasksList(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my tasks list!\n")
	w.WriteHeader(http.StatusCreated)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	taskId := r.FormValue("taskId")
	taskTitle := r.FormValue("taskTitle")
	taskDescription := r.FormValue("taskDescription")
	taskStatus := r.FormValue("taskStatus")

	payload := []byte(fmt.Sprintf(`{"message":"Successfully created task", "taskId":%s, "taskTitle":%s, "taskDescription":%s, "taskStatus":%s}`, taskId, taskTitle, taskDescription, taskStatus))

	w.WriteHeader(http.StatusCreated)
	w.Write(payload)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	taskId := r.FormValue("taskId")
	taskTitle := r.FormValue("taskTitle")
	taskDescription := r.FormValue("taskDescription")
	taskStatus := r.FormValue("taskStatus")

	payload := []byte(fmt.Sprintf(`{"message":"Successfully edited task", "taskId":%s, "taskTitle":%s, "taskDescription":%s, "taskStatus":%s}`, taskId, taskTitle, taskDescription, taskStatus))

	w.WriteHeader(http.StatusCreated)
	w.Write(payload)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	taskId := r.FormValue("taskId")

	payload := []byte(fmt.Sprintf(`{"message":"Successfully deleted task", "taskId":%s}`, taskId))

	w.WriteHeader(http.StatusCreated)
	w.Write(payload)
}
