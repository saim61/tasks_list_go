package user

import (
	"database/sql"

	"github.com/saim61/tasks_list_go/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditUserRequest struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUser(emailArg string, database *sql.DB) (string, string, User, bool) {
	id, email, password := -1, "", ""
	row := database.QueryRow(db.GET_USER_QUERY(), emailArg)

	err := row.Scan(&id, &email, &password)
	if err == sql.ErrNoRows {
		return "000x20", "No record found", User{}, false
	} else {
		user := User{Id: id, Email: email, Password: password}
		return "", "", user, true
	}
}

func RegisterUser(user RegisterUserRequest, database *sql.DB) (string, string, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "000x21", err.Error(), false
	}

	_, err = database.Exec(
		db.INSERT_USER_QUERY(),
		user.Email,
		hashedPassword,
	)

	if err != nil {
		return "000x22", err.Error(), false
	}

	return "", "", true
}

func EditUser(user EditUserRequest, database *sql.DB) (string, string, bool) {
	errorCode, errorString, _, status := GetUser(user.Email, database)
	if status {
		_, err := database.Exec(db.EDIT_USER_QUERY(), user.Id, user.Email, user.Password)
		if err != nil {
			return "000x23", err.Error(), false
		}
		return "", "", true
	}

	return errorCode, errorString, false
}
