package middleware

import (
	// "net/http" // 提示：返回 401 状态码需使用 net/http 包

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 【选做加分项】简易 Token 鉴权中间件
// 学习目标：掌握中间件拦截机制与 c.Abort() 终止调用链的正确用法
//
// TODO (选做加分项实现步骤):
// 1. 获取请求头：token := c.GetHeader("Authorization")
// 2. 检查凭证：若 token 为空或不合法（例如不等于 "tenzor-token"）：
//    - 返回 401 状态码：c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供有效凭证"})（需解开上方 net/http 包导入）
//    - 核心拦截：调用 c.Abort() 立即截断后续 Handler 执行链
//    - 函数返回：return
// 3. 凭证合法：调用 c.Next() 移交控制权
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请在此处补全 Token 拦截校验逻辑（约 6~8 行代码）

		c.Next()
	}
}
