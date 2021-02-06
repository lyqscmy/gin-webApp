package main

import (
	"log"

	//Package gin implements a HTTP web framework called gin
	"github.com/gin-gonic/gin"

	"api/handlers"
	"api/middleware"
	"api/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	//create a jwt middlware
	jwt, err := middleware.JwtMiddleware()
	if err != nil {
		log.Fatalf("faile to init jwt: %s", err)
	}

	//use localhost:8080/login to get token
	//username: admin
	//password: password
	r.POST("/login", jwt.LoginHandler)

	//check token
	auth := r.Group("/", jwt.MiddlewareFunc())

	auth.GET("/users", handlers.GetPeople)
	auth.GET("/user/:id", handlers.GetPerson)

	r.Run(":8080")
}
