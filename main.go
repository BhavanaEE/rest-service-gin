package main

import (
	"github.com/gin-gonic/gin"
	"web-service-gin/rest"
)

func main() {

	router := gin.Default()
	router.GET("/albums", rest.GetAlbums)
	router.POST("/albums", rest.PostAlbums)
	router.GET("/albums/:id", rest.GetAlbumByID)
	router.Run("localhost:8080")
}
