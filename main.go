package main

import (
	"codeskill/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建gin引擎
	r := gin.Default()

	// 设置路由
	router.SetupRoutes(r)

	// 启动服务器
	log.Println("CodeSkill服务器启动在 :8080 端口")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
