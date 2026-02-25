package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() //创建gin引擎
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.Run(":4567") //开启服务器，默认监听localhost:8080
}
