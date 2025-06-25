package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "CodeSkill服务正常运行",
	})
}

// GetSkills 获取所有编程技巧
func GetSkills(c *gin.Context) {
	skills := []map[string]interface{}{
		{
			"id":          1,
			"title":       "接口参数类型升级技巧",
			"description": "通过接口实现向后兼容的参数类型升级",
			"category":    "接口设计",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": skills,
		"msg":  "获取成功",
	})
}
