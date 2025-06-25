package services

import (
	"codeskill/internal/models"
	"encoding/json"
	"fmt"
	"time"
)

type UserService struct {
	// 这里可以添加数据库连接等依赖
}

func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser 创建用户 - 核心技巧展示
// 该方法能够处理不同版本的用户请求，实现向后兼容
func (s *UserService) CreateUser(rawData map[string]interface{}) (*models.User, error) {
	// 尝试解析为不同版本的请求结构
	userReq, err := s.parseUserRequest(rawData)
	if err != nil {
		return nil, fmt.Errorf("解析用户请求失败: %v", err)
	}

	// 使用接口方法获取通用字段，无论是哪个版本的请求
	user := &models.User{
		ID:        generateUserID(), // 模拟生成ID
		Name:      userReq.GetName(),
		Email:     userReq.GetEmail(),
		Age:       userReq.GetAge(),
		Phone:     userReq.GetPhone(),
		Address:   userReq.GetAddress(),
		CreatedAt: time.Now(),
	}

	// 如果是新版本请求，获取额外信息
	if detailedReq, ok := userReq.(*models.DetailedUserRequest); ok {
		user.Company = detailedReq.GetCompany()
	}

	return user, nil
}

// parseUserRequest 智能解析用户请求
// 编程技巧：通过字段检测来判断请求类型，实现灵活的类型转换
func (s *UserService) parseUserRequest(rawData map[string]interface{}) (models.UserRequest, error) {
	// 检查是否包含新版本特有的字段
	if _, hasPhone := rawData["phone"]; hasPhone {
		if _, hasAddress := rawData["address"]; hasAddress {
			// 包含phone和address字段，说明是新版本请求
			return s.parseAsDetailedRequest(rawData)
		}
	}

	// 否则按旧版本处理
	return s.parseAsBasicRequest(rawData)
}

func (s *UserService) parseAsBasicRequest(rawData map[string]interface{}) (*models.BasicUserRequest, error) {
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		return nil, err
	}

	var req models.BasicUserRequest
	if err := json.Unmarshal(jsonData, &req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (s *UserService) parseAsDetailedRequest(rawData map[string]interface{}) (*models.DetailedUserRequest, error) {
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		return nil, err
	}

	var req models.DetailedUserRequest
	if err := json.Unmarshal(jsonData, &req); err != nil {
		return nil, err
	}

	return &req, nil
}

// 模拟ID生成
func generateUserID() int {
	return int(time.Now().Unix()) % 100000
}
