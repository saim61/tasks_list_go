package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saim61/tasks_list_go/auth"
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
		v1.POST("/user", auth.AuthMiddleware(), GetUser)
		v1.POST("/register", RegisterUser)
		v1.POST("/login", LoginUser)

		v1.GET("/tasks", auth.AuthMiddleware(), TasksList)
		v1.GET("/task/:id", auth.AuthMiddleware(), GetTask)

		v1.POST("/createTask", auth.AuthMiddleware(), CreateTask)
		v1.PATCH("/editTask", auth.AuthMiddleware(), EditTask)
		v1.PATCH("/editTaskStatus", auth.AuthMiddleware(), EditTaskStatus)

		v1.DELETE("/deleteTask/:id", auth.AuthMiddleware(), DeleteTask)
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
