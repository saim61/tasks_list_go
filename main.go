package main

import (
	"log"

	"github.com/saim61/tasks_list_go/routes"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/saim61/tasks_list_go/docs"
)

// @Title Tasks List Go Documentation API
// @Version 1.0

// @Contact.name Saeem Mehmood
// @Contact.email saimmahmood61@gmail.com
// @Contact.url https://github.com/saim61

// @SecurityDefinitions.apiKey bearerToken
// @In header
// @Name Authorization

// @Description This is the documentation for your tasks list. It shows all the routes and whatever you can do with this service. Once you create your account, login and click on Authorize button at the top and add your bearer toke. Be sure to add `Bearer` before your token otherwise requests would fail.
// @Host localhost:8080
// @BasePath /api/v1
func main() {

	router := routes.SetupAPIRoutes()

	log.Fatal(router.Run(":8080"))
}
