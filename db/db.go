package db

import "database/sql"

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

func getDSN() string {
	// dbHost := os.Getenv(DB_HOST)
	// dbPort := os.Getenv(DB_PORT)
	// dbUser := os.Getenv(DB_USER)
	// dbPassword := os.Getenv(DB_PASSWORD)
	// dbName := os.Getenv(DB_NAME)

	return DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME
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
