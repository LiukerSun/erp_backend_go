package shop

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes 注册店铺相关路由
func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	shops := r.Group("/shops")
	{
		shops.POST("", handler.Create)
		shops.GET("", handler.List)
		shops.GET("/:id", handler.Get)
		shops.PUT("/:id", handler.Update)
		shops.DELETE("/:id", handler.Delete)
		shops.PATCH("/:id/toggle", handler.ToggleStatus)
	}
}
