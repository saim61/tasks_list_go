package db

import (
	"database/sql"
	"os"
)

const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_PASSWORD = "saeem"
const DB_USER = "root"
const DB_NAME = "tasks_list_go"

func INSERT_TASK_QUERY() string {
	return "INSERT into tasks (title, description, status) VALUES (?, ?, ?)"
}
func GET_ALL_TASKS_QUERY() string { return "SELECT * FROM tasks" }
func GET_TASK_QUERY() string      { return "SELECT * FROM tasks WHERE id = ?" }
func DELETE_TASK_QUERY() string   { return "DELETE FROM tasks WHERE id = ?" }
func EDIT_TASK_QUERY() string {
	return "UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?"
}
func EDIT_TASK_STATUS_QUERY() string {
	return "UPDATE tasks SET status = ? WHERE id = ?"
}

func INSERT_USER_QUERY() string {
	return "INSERT into users (email, password) VALUES (?, ?)"
}
func GET_USER_QUERY() string {
	return "SELECT * FROM users WHERE email = ?"
}
func EDIT_USER_QUERY() string {
	return "UPDATE user SET email = ?, password = ?"
}

func getDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

func GetDatabaseObject() *sql.DB {
	database, err := sql.Open("mysql", getDSN())
	if err != nil {
		panic(err.Error())
	}

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}

	return database
}
