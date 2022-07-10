package main

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharmamanan796/gin-gonic/gin/service"
	"gitlab.com/sharmamanan796/golang-gin-poc/controller"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
