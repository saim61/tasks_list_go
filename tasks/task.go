package tasks

import (
	"database/sql"

	"github.com/saim61/tasks_list_go/db"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type EditTaskStatusRequest struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

func GetAllTasks(database *sql.DB) (string, string, []Task) {
	rows, err := database.Query(db.GET_ALL_TASKS_QUERY())
	if err != nil {
		panic("000x1: " + err.Error())
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return "000x2", err.Error(), nil
		}
		tasks = append(tasks, task)
	}

	return "", "", tasks
}

func GetTask(taskId int, database *sql.DB) (string, string, Task, bool) {
	id, title, description, status := -1, "", "", ""
	row := database.QueryRow(db.GET_TASK_QUERY(), taskId)

	err := row.Scan(&id, &title, &description, &status)
	if err == sql.ErrNoRows {
		return "000x3", "No record found", Task{}, false
	} else {
		task := Task{Id: id, Title: title, Description: description, Status: status}
		return "", "", task, true
	}
}

func CreateTask(task CreateTaskRequest, database *sql.DB) (string, string, bool) {
	_, err := database.Exec(
		db.INSERT_TASK_QUERY(),
		task.Title,
		task.Description,
		task.Status,
	)

	if err != nil {
		return "000x4", err.Error(), false
	}

	return "", "", true
}

func EditTask(taskParams Task, database *sql.DB) (string, string, bool) {
	errorCode, errorString, _, status := GetTask(taskParams.Id, database)
	if status {
		_, err := database.Exec(db.EDIT_TASK_QUERY(), taskParams.Title, taskParams.Description, taskParams.Status, taskParams.Id)
		if err != nil {
			return "000x5", err.Error(), false
		}
		return "", "", true
	}

	return errorCode, errorString, false
}

func EditTaskStatus(taskParams EditTaskStatusRequest, database *sql.DB) (string, string, bool) {
	errorCode, errorString, _, status := GetTask(taskParams.Id, database)
	if status {
		_, err := database.Exec(db.EDIT_TASK_STATUS_QUERY(), taskParams.Status, taskParams.Id)
		if err != nil {
			return "000x6", err.Error(), false
		}
		return "", "", true
	}

	return errorCode, errorString, false
}

func DeleteTask(taskId int, database *sql.DB) (string, string, bool) {
	errorCode, errorString, _, status := GetTask(taskId, database)

	if status {
		_, err := database.Exec(
			db.DELETE_TASK_QUERY(),
			taskId,
		)

		if err != nil {
			return "000x7", err.Error(), false
		}
		return "", "", true
	}

	return errorCode, errorString, false
}
