package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/didip/tollbooth"
	limiter "github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"

	"github.com/saim61/tasks_list_go/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	csrf "github.com/utrack/gin-csrf"
)

func init() {
	setupPrometheusCounters()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func SetupAPIRoutes() *gin.Engine {
	router := SetupRouter()

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

	rps := os.Getenv("RATE_LIMIT")

	// Setting up the rate limiter for all the requests
	// For now, its set to 100 requests/second
	// The number of requests can be changed from the .env file and time can be changed from here
	allowedRequests, _ := strconv.ParseFloat(rps, 64)

	tbOptions := &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour}
	limiterNew := tollbooth.NewLimiter(allowedRequests, tbOptions)
	limiterNew.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}).
		SetMethods([]string{"GET", "POST", "PATCH", "DELETE"})

	router.Use(tollbooth_gin.LimitHandler(limiterNew))
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user", middleware.AuthMiddleware(), GetUser)
		v1.POST("/register", RegisterUser)
		v1.POST("/login", LoginUser)
		v1.PATCH("/editUser", middleware.AuthMiddleware(), EditUser)

		v1.GET("/tasks", middleware.AuthMiddleware(), TasksList)
		v1.GET("/user_tasks", middleware.AuthMiddleware(), UserTasksList)
		v1.GET("/task/:id", middleware.AuthMiddleware(), GetTask)

		v1.POST("/createTask", middleware.AuthMiddleware(), CreateTask)
		v1.PATCH("/editTask", middleware.AuthMiddleware(), EditTask)
		v1.PATCH("/editTaskStatus", middleware.AuthMiddleware(), EditTaskStatus)

		v1.DELETE("/deleteTask/:id", middleware.AuthMiddleware(), DeleteTask)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/protected", func(g *gin.Context) {
		g.String(200, csrf.GetToken(g))
	})

	router.POST("/protected", func(g *gin.Context) {
		g.String(200, "CSRF token is valid")
	})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}
