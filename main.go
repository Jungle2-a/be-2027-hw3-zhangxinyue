package main

import (
	"be-2027-hw3-template/controller"
	// "be-2027-hw3-template/middleware" // 【选做加分项】引入中间件包
	"be-2027-hw3-template/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 【选做加分项 1】：若实现了耗时统计中间件，可在此全局挂载体验洋葱模型：
	// r.Use(middleware.CostTimeMiddleware())

	// 1. 依赖装配链：Service -> Controller
	userSrv := service.NewUserService()
	userCtr := controller.NewUserController(userSrv)

	// 2. 路由分组注册
	v1 := r.Group("/api/v1")
	{
		// 用户注册核心必做接口：POST /api/v1/users/register
		v1.POST("/users/register", userCtr.Register)

		// 【选做加分项 2】：若实现了 Token 鉴权中间件，可挂载保护特定路由：
		// authed := v1.Group("/protected").Use(middleware.AuthMiddleware())
		// {
		// 	authed.GET("/ping", func(c *gin.Context) {
		// 		c.JSON(200, gin.H{"message": "Token 鉴权通过"})
		// 	})
		// }
	}

	// 3. 启动服务，默认监听 8080 端口
	r.Run(":8080")
}
