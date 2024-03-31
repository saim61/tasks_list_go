package routes

import (
	"net/http"
	"strconv"

	"github.com/saim61/tasks_list_go/db"
	"github.com/saim61/tasks_list_go/tasks"
	"github.com/saim61/tasks_list_go/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	return router
}

func SetupAPIRoutes() *gin.Engine {
	router := SetupRouter()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", TasksList)
		v1.GET("/task/:id", GetTask)

		v1.POST("/createTask", CreateTask)
		v1.POST("/editTask", EditTask)
		v1.POST("/editTaskStatus", EditTaskStatus)

		v1.DELETE("/deleteTask/:id", DeleteTask)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func SetupTestingRoutes() *gin.Engine {
	router := SetupRouter()
	testingRoutes := router.Group("/test")
	{
		testingRoutes.GET("/tasks", TasksList)
		testingRoutes.GET("/task/:id", GetTask)

		testingRoutes.POST("/createTask", CreateTask)
		testingRoutes.POST("/editTask", EditTask)
		testingRoutes.POST("/editTaskStatus", EditTaskStatus)

		testingRoutes.DELETE("/deleteTask/:id", DeleteTask)
	}
	return router
}

// TasksList godoc
// @Summary Get tasks list
// @description Get and view all your tasks in this route.
// @Tags Tasks
// @Success 200 {array} tasks.Task
// @failure 403 {object} utils.ErrorResponse
// @Router /tasks [get]
func TasksList(g *gin.Context) {
	database := db.GetDatabaseObject()
	defer database.Close()

	errorCode, errorString, tasks := tasks.GetAllTasks(database)

	if tasks != nil {
		g.JSON(http.StatusOK, tasks)
	} else {
		g.JSON(http.StatusForbidden, utils.NewErrorResponse(errorCode, errorString, "Failed to get tasks"))
	}
}

// GetTask get a specific task
// @Summary Get a task by its id
// @Description Retreive your task by its id
// @Tags Tasks
// @Param id query int true "Required task id"
// @Success 200 {object} tasks.Task
// @Failure 400 {object} utils.ErrorResponse
// @Router /task/:id [get]
func GetTask(g *gin.Context) {
	database := db.GetDatabaseObject()
	defer database.Close()

	theParams := g.Request.URL.Query()
	id := theParams["id"]
	idConverted, err := strconv.Atoi(id[0])
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.NewErrorResponse("000x8", err.Error(), "Failed to get task"))
	}

	errorCode, errorString, task, status := tasks.GetTask(idConverted, database)
	if status {
		g.JSON(http.StatusOK, task)
	} else {
		g.JSON(http.StatusBadRequest, utils.NewErrorResponse(errorCode, errorString, "Failed to get task"))
	}
}

// DeleteTask delete a specific task
// @Summary Delete a task by its id
// @Description Delete your task by its id
// @Tags Tasks
// @Param id query int true "Required task id"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /deleteTask/:id [delete]
func DeleteTask(g *gin.Context) {
	database := db.GetDatabaseObject()
	defer database.Close()

	theParams := g.Request.URL.Query()
	id := theParams["id"]
	idConverted, err := strconv.Atoi(id[0])
	if err != nil {
		g.JSON(http.StatusForbidden, utils.NewErrorResponse("000x9", err.Error(), "Failed to delete task"))
	}

	errorCode, errorString, status := tasks.DeleteTask(idConverted, database)
	if status {
		g.JSON(http.StatusOK, utils.NewSuccessResponse("Successfully deleted task"))
	} else {
		g.JSON(http.StatusForbidden, utils.NewErrorResponse(errorCode, errorString, "Failed to delete task"))
	}
}

// CreateTask create a task
// @Summary Create a task
// @Description Create a task as per your liking
// @Tags Tasks
// @Param task body tasks.CreateTaskRequest true "Required create task parameters"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /createTask [post]
func CreateTask(g *gin.Context) {
	database := db.GetDatabaseObject()
	defer database.Close()

	var task tasks.CreateTaskRequest
	err := g.ShouldBindJSON(&task)
	if !utils.IsValidCreateTask(task) || err != nil {
		g.JSON(http.StatusBadRequest, utils.NewErrorResponse("000x10", "Invalid parameters", "Request body not parsed"))
		return
	}

	errorCode, errorString, status := tasks.CreateTask(task, database)
	if status {
		g.JSON(http.StatusCreated, utils.NewSuccessResponse("Successfully created task"))
	} else {
		g.JSON(http.StatusBadRequest, utils.NewErrorResponse(errorCode, errorString, "Failed to create task"))
	}

}

// EditTask edit a task
// @Summary Edit a task
// @Description Edit a task as per your liking. Add the task id and the other parameters
// @Tags Tasks
// @Param task body tasks.Task true "Required edit task parameters"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /editTask [post]
func EditTask(g *gin.Context) {
	database := db.GetDatabaseObject()
	defer database.Close()

	var task tasks.Task
	err := g.ShouldBindJSON(&task)
	if !utils.IsValidEditTask(task) || err != nil {
		g.JSON(http.StatusBadRequest, utils.NewErrorResponse("000x11", "Invalid parameters", "Request body not parsed"))
		return
	}

	errorCode, errorString, status := tasks.EditTask(task, database)
	if status {
		g.JSON(http.StatusOK, utils.NewSuccessResponse("Successfully edited task"))
	} else {
		g.JSON(http.StatusBadRequest, utils.NewErrorResponse(errorCode, errorString, "Failed to edit task"))
	}
}

// EditTaskStatus edit a task status
// @Summary Edit a task status
// @Description Edit a task status
// @Tags Tasks
// @Param task body tasks.EditTaskStatusRequest true "Required edit task status parameters"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /editTaskStatus [post]
func EditTaskStatus(g *gin.Context) {
	database := db.GetDatabaseObject()
	defer database.Close()

	var task tasks.EditTaskStatusRequest
	err := g.ShouldBindJSON(&task)
	if !utils.IsValidEditTaskStatus(task) || err != nil {
		g.JSON(http.StatusBadRequest, utils.NewErrorResponse("000x12", "Invalid parameters", "Request body not parsed"))
		return
	}

	errorCode, errorString, status := tasks.EditTaskStatus(task, database)
	if status {
		g.JSON(http.StatusOK, utils.NewSuccessResponse("Successfully edited task status"))
	} else {
		g.JSON(http.StatusBadRequest, utils.NewErrorResponse(errorCode, errorString, "Failed to edit task status"))
	}
}
