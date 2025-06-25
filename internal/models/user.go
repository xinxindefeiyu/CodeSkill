package models

import "time"

// UserRequest 通用用户请求接口
// 编程技巧：通过接口实现参数类型的向后兼容升级
type UserRequest interface {
	GetName() string
	GetEmail() string
	GetAge() int
	GetPhone() string
	GetAddress() string
}

// BasicUserRequest 基础用户请求结构（旧版本）
type BasicUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int    `json:"age"`
}

// 实现UserRequest接口的方法
func (b *BasicUserRequest) GetName() string {
	return b.Name
}

func (b *BasicUserRequest) GetEmail() string {
	return b.Email
}

func (b *BasicUserRequest) GetAge() int {
	return b.Age
}

func (b *BasicUserRequest) GetPhone() string {
	return "" // 旧版本没有电话信息
}

func (b *BasicUserRequest) GetAddress() string {
	return "" // 旧版本没有地址信息
}

// DetailedUserRequest 详细用户请求结构（新版本）
type DetailedUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Company string `json:"company"` // 新增字段
}

// 实现UserRequest接口的方法
func (d *DetailedUserRequest) GetName() string {
	return d.Name
}

func (d *DetailedUserRequest) GetEmail() string {
	return d.Email
}

func (d *DetailedUserRequest) GetAge() int {
	return d.Age
}

func (d *DetailedUserRequest) GetPhone() string {
	return d.Phone
}

func (d *DetailedUserRequest) GetAddress() string {
	return d.Address
}

// GetCompany 新版本特有的方法
func (d *DetailedUserRequest) GetCompany() string {
	return d.Company
}

// User 用户实体
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	Company   string    `json:"company,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
