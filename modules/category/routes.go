package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes 注册分类相关路由
func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	categories := r.Group("/categories")
	{
		categories.POST("", handler.Create)
		categories.GET("", handler.List)
		categories.GET("/:id", handler.Get)
		categories.PUT("/:id", handler.Update)
		categories.DELETE("/:id", handler.Delete)
		categories.PATCH("/:id/toggle", handler.ToggleStatus)
		categories.GET("/:id/children", handler.GetChildren)
	}
}
