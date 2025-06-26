package attribute

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes 注册属性相关路由
func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	// 属性管理路由
	attributes := r.Group("/attributes")
	{
		attributes.POST("", handler.CreateAttribute)
		attributes.GET("", handler.ListAttributes)
		attributes.GET("/:id", handler.GetAttribute)
		attributes.PUT("/:id", handler.UpdateAttribute)
		attributes.DELETE("/:id", handler.DeleteAttribute)
		attributes.PATCH("/:id/toggle", handler.ToggleAttributeStatus)
	}

	// 商品属性值路由
	productAttributes := r.Group("/product-attributes")
	{
		productAttributes.POST("", handler.CreateProductAttribute)
		productAttributes.GET("", handler.ListProductAttributes)
		productAttributes.PUT("/:id", handler.UpdateProductAttribute)
		productAttributes.DELETE("/:id", handler.DeleteProductAttribute)
	}
}
