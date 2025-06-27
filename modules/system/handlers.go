package system

import (
	"time"

	"github.com/gin-gonic/gin"

	"erp_backend/pkg/response"
)

// @Summary 健康检查
// @Description 检查服务是否正常运行
// @Tags 系统
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=map[string]string} "服务运行正常"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	response.Success(c, gin.H{
		"status": "ok",
		"time":   time.Now(),
	})
}

// SystemInfo 系统信息
func SystemInfo(c *gin.Context) {
	response.Success(c, gin.H{
		"name":    "ERP Backend",
		"version": "1.0.0",
		"env":     gin.Mode(),
	})
}
