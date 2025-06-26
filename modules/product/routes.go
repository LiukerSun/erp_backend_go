package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes 注册商品相关路由
func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	products := r.Group("/products")
	{
		products.POST("", handler.Create)
		products.GET("", handler.List)
		products.GET("/:id", handler.Get)
		products.PUT("/:id", handler.Update)
		products.DELETE("/:id", handler.Delete)
		products.PATCH("/:id/toggle", handler.ToggleStatus)
		products.PATCH("/:id/stock", handler.UpdateStock)
		products.PATCH("/:id/price", handler.UpdatePrice)
	}
}
