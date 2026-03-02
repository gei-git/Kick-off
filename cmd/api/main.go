package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gei-git/Kick-off/internal/config"
	"github.com/gei-git/Kick-off/internal/handler"
	"github.com/gei-git/Kick-off/internal/middleware"
	"github.com/gei-git/Kick-off/internal/model"
	"github.com/gei-git/Kick-off/internal/repository"
	"github.com/gei-git/Kick-off/internal/service"

	_ "github.com/gei-git/Kick-off/docs" // ← swag 生成的文档，必须保留
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable client_encoding=utf8 TimeZone=Asia/Shanghai",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 自动迁移（Task + User 都要迁移！）
	repo := repository.NewTaskRepository(db)
	if err := repo.AutoMigrate(); err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	// 迁移 User 表
	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic("User 表迁移失败: " + err.Error())
	}

	fmt.Println("✅ 数据库连接成功 + 表迁移完成！")
	// 生产环境建议加这行
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康检查端点（企业必备）
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "ok",
			"time":    gin.H{"now": "2026-02-25"}, // 后续会用真实时间
		})
	})

	// 后续会加 /api/v1/tasks 等路由
	// ... 原有数据库初始化代码保持不变 ...

	// 初始化 Service 和 Handler
	taskService := service.NewTaskService(db)
	taskHandler := handler.NewTaskHandler(taskService)

	authService := service.NewAuthService(db)
	authHandler := handler.NewAuthHandler(authService)

	// API 路由组（企业标准 v1）
	v1 := r.Group("/api/v1")
	{
		// Auth 路由（公开，无需登录）
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// 受保护的任务路由（必须带 JWT Token）
		tasks := v1.Group("/tasks")
		tasks.Use(middleware.JWTAuth())
		{
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("", taskHandler.ListTasks)
		}
	}
	fmt.Println("🚀 API 服务启动成功！端口:", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
	r.Run(":4567") // ← 改成 8080
}
