package system

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes 注册系统相关路由
func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"time":   time.Now(),
		})
	})

	// 系统信息
	r.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "ERP Backend",
			"version": "1.0.0",
			"env":     gin.Mode(),
		})
	})
}
