package utils

import (
	"fmt"
	"log"
	"net/mail"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/saim61/tasks_list_go/tasks"
)

type ErrorResponse struct {
	ErrorCode   string
	ErrorString string
	Message     string
}

type SuccessResponse struct {
	Message string
}

type LoginSuccessResponse struct {
	Message string
	Token   string
}

func NewErrorResponse(errorCode string, errorString string, message string) ErrorResponse {
	return ErrorResponse{
		ErrorCode: errorCode, ErrorString: errorString, Message: message,
	}
}

func NewSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{
		Message: message,
	}
}

func NewLoginSuccessResponse(message string, token string) LoginSuccessResponse {
	return LoginSuccessResponse{
		Message: message,
		Token:   token,
	}
}

func IsValidIdForTask(id int) bool {
	return id > 0
}

func IsValidEditTask(task tasks.Task) bool {
	if task.Description == "" || task.Title == "" || task.Status == "" {
		return false
	}

	if !IsValidIdForTask(task.Id) {
		return false
	}
	return true
}

func IsValidEditTaskStatus(task tasks.EditTaskStatusRequest) bool {
	if task.Status == "" {
		return false
	}

	if !IsValidIdForTask(task.Id) {
		return false
	}
	return true
}

func IsValidCreateTask(task tasks.CreateTaskRequest) bool {
	if task.Description == "" || task.Title == "" || task.Status == "" {
		return false
	}
	return true
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func GetJWKTokenFromJWT(authParts []string) (*jwt.Token, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	return token, err
}

func GetUserEmailFromJWT(authHeader string) string {
	authParts := strings.Split(authHeader, " ")
	token, _ := GetJWKTokenFromJWT(authParts)

	return token.Claims.(jwt.MapClaims)["Email"].(string)
}

func PrintTask(task tasks.Task) {
	log.Println("Id: ", task.Id)
	log.Println("Title: ", task.Id)
	log.Println("Description: ", task.Id)
	log.Println("Status: ", task.Id)
}

func PrintTasks(tasks []tasks.Task) {
	for i := 0; i < len(tasks); i++ {
		PrintTask(tasks[i])
	}
}
