package routes

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/saim61/tasks_list_go/db"
	"github.com/saim61/tasks_list_go/tasks"
	userPkg "github.com/saim61/tasks_list_go/user"
	"github.com/saim61/tasks_list_go/utils"

	"github.com/gin-gonic/gin"
)

var errorResponse utils.ErrorResponse
var successResponse utils.SuccessResponse
var mode string

// TasksList godoc
// @Summary Get tasks list of all users
// @description Get and view all the tasks by all users in this route.
// @security bearerToken
// @scheme bearer
// @Tags Tasks
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {array} tasks.Task
// @failure 403 {object} utils.ErrorResponse
// @Router /tasks [get]
func TasksList(g *gin.Context) {
	path := g.Request.URL.Path
	if strings.Contains(path, "/test/") {
		mode = "test"
	} else {
		mode = "dev"
	}

	tasksRequestCounter.Inc()
	log.Println("Request to get Tasks List")
	database := db.GetDatabaseObject(mode)
	defer database.Close()

	errorCode, errorString, tasks := tasks.GetAllTasks(database)

	if tasks != nil {
		log.Println(utils.NewSuccessResponse("Successfully fetched tasks"))
		utils.PrintTasks(tasks)
		g.JSON(http.StatusOK, tasks)
	} else {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Failed to get tasks")
		log.Println(errorResponse)
		g.JSON(http.StatusForbidden, errorResponse)
	}
}

// TasksList godoc
// @Summary Get tasks list of current user
// @description Get and view all the tasks created by you.
// @security bearerToken
// @scheme bearer
// @Tags Tasks
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {array} tasks.Task
// @failure 403 {object} utils.ErrorResponse
// @Router /user_tasks [get]
func UserTasksList(g *gin.Context) {
	path := g.Request.URL.Path
	if strings.Contains(path, "/test/") {
		mode = "test"
	} else {
		mode = "dev"
	}

	userTasksRequestCounter.Inc()
	log.Println("Request to get Tasks List of current user")
	database := db.GetDatabaseObject(mode)
	defer database.Close()

	userEmail := utils.GetUserEmailFromJWT(g.GetHeader("Authorization"))
	_, _, user, _ := userPkg.GetUser(userEmail, database)

	errorCode, errorString, tasks := tasks.GetAllUserTasks(user.Id, database)

	if tasks != nil {
		log.Println(utils.NewSuccessResponse("Successfully fetched tasks"))
		utils.PrintTasks(tasks)
		g.JSON(http.StatusOK, tasks)
	} else {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Failed to get tasks")
		log.Println(errorResponse)
		g.JSON(http.StatusForbidden, errorResponse)
	}
}

// GetTask get a specific task
// @Summary Get a task by its id
// @Description Retreive your task by its id
// @security bearerToken
// @scheme bearer
// @Tags Tasks
// @Param id query int true "Required task id"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} tasks.Task
// @Failure 400 {object} utils.ErrorResponse
// @Router /task/:id [get]
func GetTask(g *gin.Context) {
	path := g.Request.URL.Path
	if strings.Contains(path, "/test/") {
		mode = "test"
	} else {
		mode = "dev"
	}

	taskRequestCounter.Inc()
	log.Println("Request to get Task by ID")
	database := db.GetDatabaseObject(mode)
	defer database.Close()

	userEmail := utils.GetUserEmailFromJWT(g.GetHeader("Authorization"))
	_, _, user, _ := userPkg.GetUser(userEmail, database)

	theParams := g.Request.URL.Query()
	id := theParams["id"]
	idConverted, err := strconv.Atoi(id[0])
	if err != nil {
		errorResponse = utils.NewErrorResponse("000x8", err.Error(), "Failed to get task")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
	}

	errorCode, errorString, task, status := tasks.GetTask(idConverted, user.Id, database)
	if status {
		log.Println(utils.NewSuccessResponse("Successfully fetched task"))
		utils.PrintTask(task)
		g.JSON(http.StatusOK, task)
	} else {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Failed to get task")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
	}
}

// DeleteTask delete a specific task
// @Summary Delete a task by its id
// @Description Delete your task by its id
// @security bearerToken
// @scheme bearer
// @Tags Tasks
// @Param id query int true "Required task id"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /deleteTask/:id [delete]
func DeleteTask(g *gin.Context) {
	path := g.Request.URL.Path
	if strings.Contains(path, "/test/") {
		mode = "test"
	} else {
		mode = "dev"
	}

	deleteTasksRequestCounter.Inc()
	log.Println("Request to delete task")
	database := db.GetDatabaseObject(mode)
	defer database.Close()

	userEmail := utils.GetUserEmailFromJWT(g.GetHeader("Authorization"))
	_, _, user, _ := userPkg.GetUser(userEmail, database)

	theParams := g.Request.URL.Query()
	id := theParams["id"]
	idConverted, err := strconv.Atoi(id[0])
	if err != nil {
		errorResponse = utils.NewErrorResponse("000x9", err.Error(), "Failed to delete task")
		log.Println(errorResponse)
		g.JSON(http.StatusForbidden, errorResponse)
	}

	errorCode, errorString, status := tasks.DeleteTask(idConverted, user.Id, database)
	if status {
		successResponse = utils.NewSuccessResponse("Successfully deleted task # " + id[0])
		log.Println(successResponse)
		g.JSON(http.StatusOK, successResponse)
	} else {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Failed to delete task")
		log.Println(errorResponse)
		g.JSON(http.StatusForbidden, errorResponse)
	}
}

// CreateTask create a task
// @Summary Create a task
// @Description Create a task as per your liking
// @security bearerToken
// @scheme bearer
// @Tags Tasks
// @Param task body tasks.CreateTaskRequest true "Required create task parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /createTask [post]
func CreateTask(g *gin.Context) {
	path := g.Request.URL.Path
	if strings.Contains(path, "/test/") {
		mode = "test"
	} else {
		mode = "dev"
	}

	createTasksRequestCounter.Inc()
	log.Println("Request to create task")
	database := db.GetDatabaseObject(mode)
	defer database.Close()

	userEmail := utils.GetUserEmailFromJWT(g.GetHeader("Authorization"))
	_, _, user, _ := userPkg.GetUser(userEmail, database)

	var task tasks.CreateTaskRequest
	err := g.ShouldBindJSON(&task)
	if !utils.IsValidCreateTask(task) || err != nil {
		errorResponse = utils.NewErrorResponse("000x10", "Invalid parameters", "Request body not parsed")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorCode, errorString, status := tasks.CreateTask(task, user.Id, database)
	if status {
		successResponse = utils.NewSuccessResponse("Successfully created task")
		log.Println(successResponse)
		g.JSON(http.StatusCreated, successResponse)
	} else {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Failed to create task")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
	}

}

// EditTask edit a task
// @Summary Edit a task
// @Description Edit a task as per your liking. Add the task id and the other parameters
// @security bearerToken
// @scheme bearer
// @Tags Tasks
// @Param task body tasks.Task true "Required edit task parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /editTask [patch]
func EditTask(g *gin.Context) {
	path := g.Request.URL.Path
	if strings.Contains(path, "/test/") {
		mode = "test"
	} else {
		mode = "dev"
	}

	editTaskRequestCounter.Inc()
	log.Println("Request to edit task")
	database := db.GetDatabaseObject(mode)
	defer database.Close()

	userEmail := utils.GetUserEmailFromJWT(g.GetHeader("Authorization"))
	_, _, user, _ := userPkg.GetUser(userEmail, database)

	var task tasks.Task
	err := g.ShouldBindJSON(&task)
	if !utils.IsValidEditTask(task) || err != nil {
		errorResponse = utils.NewErrorResponse("000x11", "Invalid parameters", "Request body not parsed")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorCode, errorString, status := tasks.EditTask(task, user.Id, database)
	if status {
		successResponse = utils.NewSuccessResponse("Successfully edited task")
		log.Println(successResponse)
		g.JSON(http.StatusOK, successResponse)
	} else {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Failed to edit task")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
	}
}

// EditTaskStatus edit a task status
// @Summary Edit a task status
// @Description Edit a task status
// @security bearerToken
// @scheme bearer
// @Tags Tasks
// @Param task body tasks.EditTaskStatusRequest true "Required edit task status parameters"
// @Param X-CSRF-token header string true "Insert your CSRF token. Access the GET /protected route to get it"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /editTaskStatus [patch]
func EditTaskStatus(g *gin.Context) {
	path := g.Request.URL.Path
	if strings.Contains(path, "/test/") {
		mode = "test"
	} else {
		mode = "dev"
	}

	editTaskStatusRequestCounter.Inc()
	log.Println("Request to edit a task's status")
	database := db.GetDatabaseObject(mode)
	defer database.Close()

	userEmail := utils.GetUserEmailFromJWT(g.GetHeader("Authorization"))
	_, _, user, _ := userPkg.GetUser(userEmail, database)

	var task tasks.EditTaskStatusRequest
	err := g.ShouldBindJSON(&task)
	if !utils.IsValidEditTaskStatus(task) || err != nil {
		errorResponse = utils.NewErrorResponse("000x12", "Invalid parameters", "Request body not parsed")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	errorCode, errorString, status := tasks.EditTaskStatus(task, user.Id, database)
	if status {
		successResponse = utils.NewSuccessResponse("Successfully edited task status")
		log.Println(successResponse)
		g.JSON(http.StatusOK, successResponse)
	} else {
		errorResponse = utils.NewErrorResponse(errorCode, errorString, "Failed to edit task status")
		log.Println(errorResponse)
		g.JSON(http.StatusBadRequest, errorResponse)
	}
}
