package supplier

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes 注册供应商相关路由
func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	suppliers := r.Group("/suppliers")
	{
		suppliers.POST("", handler.Create)
		suppliers.GET("", handler.List)
		suppliers.GET("/:id", handler.Get)
		suppliers.PUT("/:id", handler.Update)
		suppliers.DELETE("/:id", handler.Delete)
		suppliers.PATCH("/:id/toggle", handler.ToggleStatus)
	}
}
