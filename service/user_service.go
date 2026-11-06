package service

import (
	"errors"
	"sync"
)

// UserService 用户业务逻辑服务
// 【架构红线】：纯 Go 业务逻辑，面向单元测试与多协议复用，严禁引入 *gin.Context
type UserService struct {
	storage sync.Map // 内存模拟数据库存储（并发安全）
}

func NewUserService() *UserService {
	return &UserService{}
}

// Register 执行用户注册核心业务
// TODO: 请在此处实现纯 Go 业务规则：
// 1. 数据排重：检查 username 是否已被占用，若已存在则返回错误 errors.New("用户名已被占用")
// 2. 数据持久化：若不存在，将用户信息存入 storage，返回 nil
func (s *UserService) Register(username, password, email string, age int) error {
	// 请在下方补全业务逻辑（代码仅需 5~8 行）

	return errors.New("TODO: 尚未实现 Register 业务逻辑")
}
