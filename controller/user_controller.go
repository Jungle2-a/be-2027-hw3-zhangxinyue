package controller

import (
	"net/http"

	"be-2027-hw3-template/service"
	"github.com/gin-gonic/gin"
)

// UserController 用户表现层控制器
// 核心职责：协议解析、参数校验、调用下游 Service、映射 HTTP 状态码
type UserController struct {
	userSrv *service.UserService
}

func NewUserController(userSrv *service.UserService) *UserController {
	return &UserController{userSrv: userSrv}
}

// RegisterRequest 请求数据传输对象 (DTO)
// TODO 1: 请完善结构体字段与 Tag 规则：
// - Username: 必填，长度 4 ~ 16 位
// - Password: 必填，长度 6 ~ 32 位
// - Email:    必填，必须符合邮箱格式
// - Age:      选填（omitempty），范围 1 ~ 120
type RegisterRequest struct {
	// 请在下方补充字段（提示：首字母必须大写导出）
	Username string `json:"username" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=6,max=32"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"omitempty,gte=1,lte=120"`
}

// Register 处理用户注册请求 POST /api/v1/users/register
func (u *UserController) Register(c *gin.Context) {
	// TODO 2: 声明 req 并使用 c.ShouldBindJSON 绑定与校验参数
	// 若校验失败，返回 400 Bad Request 与具体错误信息
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO 3: 调用下游 u.userSrv.Register 业务方法
	// 若返回错误（用户名重复），返回 409 Conflict 与错误信息
	if err := u.userSrv.Register(req.Username, req.Password, req.Email, req.Age);err!= nil{
	c.JSON(http.StatusConflict,gin.H{
		"error":err.Error(),
	})
	return
}
	// TODO 4: 业务执行成功，返回 200 OK 与成功提示 JSON
	c.JSON(http.StatusOK,gin.H{
		"message":"注册成功",
		"username":req.Username,
	})


	c.JSON(http.StatusNotImplemented, gin.H{
		"code":    501,
		"message": "TODO: 尚未实现 Controller Register 接口",
	})
}
