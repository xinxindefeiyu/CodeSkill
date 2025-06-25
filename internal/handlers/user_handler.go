package handlers

import (
	"codeskill/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser 创建用户接口 - 展示接口升级技巧
// 该接口既能处理旧版本的BasicUserRequest，也能处理新版本的DetailedUserRequest
func CreateUser(c *gin.Context) {
	var rawData map[string]interface{}
	if err := c.ShouldBindJSON(&rawData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数格式错误",
		})
		return
	}

	// 使用服务层处理用户创建逻辑
	userService := services.NewUserService()
	user, err := userService.CreateUser(rawData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
		"msg":  "用户创建成功",
	})
}
