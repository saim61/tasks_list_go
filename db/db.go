package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func INSERT_TASK_QUERY() string {
	return "INSERT into tasks (title, description, status, user_id) VALUES (?, ?, ?, ?)"
}
func GET_ALL_TASKS_QUERY() string      { return "SELECT * FROM tasks" }
func GET_ALL_USER_TASKS_QUERY() string { return "SELECT * FROM tasks WHERE user_id = ?" }
func GET_TASK_QUERY() string           { return "SELECT * FROM tasks WHERE id = ? AND user_id = ?" }
func DELETE_TASK_QUERY() string        { return "DELETE FROM tasks WHERE id = ?" }
func EDIT_TASK_QUERY() string {
	return "UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ? AND user_id = ?"
}
func EDIT_TASK_STATUS_QUERY() string {
	return "UPDATE tasks SET status = ? WHERE id = ? AND user_id = ?"
}

func INSERT_USER_QUERY() string {
	return "INSERT into users (email, password) VALUES (?, ?)"
}
func GET_USER_QUERY() string {
	return "SELECT * FROM users WHERE email = ?"
}
func EDIT_USER_QUERY() string {
	return "UPDATE users SET email = ?, password = ? WHERE id = ?"
}

func getDSN(mode string) string {
	var dbUser, dbPassword, dbHost, dbPort, dbName string
	if mode == "dev" {
		dbUser = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	} else {
		dbUser = "root"
		dbPassword = "saeem"
		dbHost = "localhost"
		dbPort = "3306"
		dbName = "test_tasks_list_go"
	}

	return dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

func GetDatabaseObject(mode string) *sql.DB {
	database, err := sql.Open("mysql", getDSN(mode))
	if err != nil {
		panic(err.Error())
	}

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}

	return database
}
