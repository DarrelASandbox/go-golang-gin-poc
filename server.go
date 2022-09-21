package main

import (
	"io"
	"net/http"
	"os"

	"github.com/DarrelASandbox/go-golang-gin-poc/controller"
	"github.com/DarrelASandbox/go-golang-gin-poc/middlewares"
	"github.com/DarrelASandbox/go-golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()
	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())
	// server.Use(gin.Recovery(), gin.Logger()) is same as server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
		}
	})

	server.Run(":4000")
}
