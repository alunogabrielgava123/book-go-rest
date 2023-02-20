package main

import (
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/factory"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	
	controller := factory.FacotoryControllerAlbum()
	

	router := gin.Default()
	router.Use(middleware.Logger())
	router.GET("/albums", controller.FindeAlbuns)
	router.POST("/albums", controller.AddAlbun)
	router.DELETE("/albums/:id", controller.DeleteAlbumById)
	router.Run("localhost:8080")

}
