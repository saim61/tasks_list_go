package main

import (
	"github.com/saim61/tasks_list_go/routes"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/saim61/tasks_list_go/docs"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Tasks List Go Documentation API
// @version 1.0
// @Description This is the documentation for your tasks list. It shows all the routes and whatever you can do with this service.
// @host localhost:8080
// @BasePath /api/v1
func main() {

	router := routes.SetupRouter()
	router.Use(cors.Default())
	router = routes.SetupAPIRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
