package main

import (
	"shortener/src/controller"
	"shortener/src/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.Connect()

	repository.RunMigrations()

	router := gin.Default()

	router.POST("/shortener", controller.Create)
	router.GET("/r/:slug", controller.RedirectBySlug)

	router.Run(":8080")
}
