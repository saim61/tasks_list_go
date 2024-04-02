package routes

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/saim61/tasks_list_go/auth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	csrf "github.com/utrack/gin-csrf"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	store := cookie.NewStore([]byte(os.Getenv("COOKIE_STORE")))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	router.GET("/protected", func(c *gin.Context) {
		c.String(200, csrf.GetToken(c))
	})

	router.POST("/protected", func(c *gin.Context) {
		c.String(200, "CSRF token is valid")
	})

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
