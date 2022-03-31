package main

import (
	"github.com/gin-gonic/gin"
	"web-service-gin/rest"
)

func main() {

	router := gin.Default()
	router.GET("/albums", rest.GetAlbums)
	router.Run("localhost:8080")
}