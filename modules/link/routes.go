package link

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes 注册链接相关路由
func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	links := r.Group("/links")
	{
		links.POST("", handler.Create)
		links.GET("", handler.List)
		links.GET("/:id", handler.Get)
		links.PUT("/:id", handler.Update)
		links.DELETE("/:id", handler.Delete)
		links.PATCH("/:id/toggle", handler.ToggleStatus)
	}
}
