package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", HomeController)
	server.Run()
}

func HomeController(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello amigos!",
	})
}
