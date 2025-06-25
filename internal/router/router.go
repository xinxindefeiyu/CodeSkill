package router

import (
	"codeskill/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有路由
func SetupRoutes(r *gin.Engine) {
	// 健康检查接口
	r.GET("/health", handlers.HealthCheck)

	// API版本分组
	v1 := r.Group("/api/v1")
	{
		// 接口升级技巧示例
		v1.POST("/user/create", handlers.CreateUser)
		v1.GET("/skills", handlers.GetSkills)
	}
}
