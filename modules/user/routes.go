package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @title ERP系统 API
// @version 1.0
// @description ERP系统后端API文档
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// RegisterRoutes 注册用户相关路由
func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	// 认证相关路由
	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)       // @Summary 用户登录
		auth.POST("/register", handler.Register) // @Summary 用户注册
	}

	// 用户管理路由
	users := r.Group("/users")
	{
		users.GET("", handler.List)                    // @Summary 获取用户列表
		users.GET("/:id", handler.Get)                 // @Summary 获取单个用户
		users.POST("", handler.Create)                 // @Summary 创建用户
		users.PUT("/:id", handler.Update)              // @Summary 更新用户
		users.DELETE("/:id", handler.Delete)           // @Summary 删除用户
		users.GET("/profile", handler.GetProfile)      // @Summary 获取个人资料
		users.PUT("/profile", handler.UpdateProfile)   // @Summary 更新个人资料
		users.PUT("/password", handler.UpdatePassword) // @Summary 修改密码
	}
}
