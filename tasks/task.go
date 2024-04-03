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
	UserId      int    `json:"user_id"`
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
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.UserId)
		if err != nil {
			return "000x2", err.Error(), nil
		}
		tasks = append(tasks, task)
	}

	return "", "", tasks
}

func GetAllUserTasks(userId int, database *sql.DB) (string, string, []Task) {
	rows, err := database.Query(db.GET_ALL_USER_TASKS_QUERY(), userId)
	if err != nil {
		panic("000x3: " + err.Error())
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.UserId)
		if err != nil {
			return "000x4", err.Error(), nil
		}
		tasks = append(tasks, task)
	}

	return "", "", tasks
}

func GetTask(taskId int, userIdArg int, database *sql.DB) (string, string, Task, bool) {
	id, title, description, status, userId := -1, "", "", "", -1
	row := database.QueryRow(db.GET_TASK_QUERY(), taskId, userIdArg)

	err := row.Scan(&id, &title, &description, &status, &userId)
	if err == sql.ErrNoRows {
		return "000x5", "No record found", Task{}, false
	} else {
		task := Task{Id: id, Title: title, Description: description, Status: status}
		return "", "", task, true
	}
}

func CreateTask(task CreateTaskRequest, userId int, database *sql.DB) (string, string, bool) {
	_, err := database.Exec(
		db.INSERT_TASK_QUERY(),
		task.Title,
		task.Description,
		task.Status,
		userId,
	)

	if err != nil {
		return "000x6", err.Error(), false
	}

	return "", "", true
}

func EditTask(taskParams Task, userIdArg int, database *sql.DB) (string, string, bool) {
	errorCode, errorString, _, status := GetTask(taskParams.Id, userIdArg, database)
	if status {
		_, err := database.Exec(db.EDIT_TASK_QUERY(), taskParams.Title, taskParams.Description, taskParams.Status, taskParams.Id, userIdArg)
		if err != nil {
			return "000x7", err.Error(), false
		}
		return "", "", true
	}

	return errorCode, errorString, false
}

func EditTaskStatus(taskParams EditTaskStatusRequest, userIdArg int, database *sql.DB) (string, string, bool) {
	errorCode, errorString, _, status := GetTask(taskParams.Id, userIdArg, database)
	if status {
		_, err := database.Exec(db.EDIT_TASK_STATUS_QUERY(), taskParams.Status, taskParams.Id, userIdArg)
		if err != nil {
			return "000x8", err.Error(), false
		}
		return "", "", true
	}

	return errorCode, errorString, false
}

func DeleteTask(taskId int, userIdArg int, database *sql.DB) (string, string, bool) {
	errorCode, errorString, _, status := GetTask(taskId, userIdArg, database)

	if status {
		_, err := database.Exec(
			db.DELETE_TASK_QUERY(),
			taskId,
		)

		if err != nil {
			return "000x9", err.Error(), false
		}
		return "", "", true
	}

	return errorCode, errorString, false
}
