package middleware

import (
	// "log"  // 提示：打印耗时日志需使用 log 包
	// "time" // 提示：计算耗时需使用 time 包

	"github.com/gin-gonic/gin"
)

// CostTimeMiddleware 【选做加分项】自定义耗时统计中间件
// 学习目标：掌握中间件洋葱模型执行时序（前置打点 -> c.Next() -> 后置耗时计算与日志审计）
//
// TODO (选做加分项实现步骤):
// 1. 记录开始时间：start := time.Now()（需解开上方 log 与 time 包导入）
// 2. 移交控制权：调用 c.Next() 让下游 Handler 处理业务
// 3. 洋葱回溯打印耗时：log.Printf("[耗时统计] %s | %v", c.Request.URL.Path, time.Since(start))
func CostTimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请在此处补全耗时统计洋葱模型逻辑（约 3 行代码）

		c.Next()
	}
}
