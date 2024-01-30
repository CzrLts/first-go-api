package main

import (
	"net/http"

	"first-api.com/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", GetEvents)
	server.POST("/events", CreateEvent)
	server.Run()
}

func GetEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}
	event.ID = 1
	event.UserId = 1231
	event.Save()
	context.JSON(http.StatusOK, gin.H{"message": "created!"})
}
