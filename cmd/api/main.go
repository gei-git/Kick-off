package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 生产环境建议加这行
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 健康检查端点（企业必备）
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "ok",
			"time":    gin.H{"now": "2026-02-25"}, // 后续会用真实时间
		})
	})

	// 后续会加 /api/v1/tasks 等路由

	r.Run(":4567") // ← 改成 8080
}
