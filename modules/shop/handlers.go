package shop

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"erp_backend/pkg/response"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// Create 创建店铺
// @Summary 创建店铺
// @Description 创建新的店铺
// @Tags 店铺管理
// @Accept json
// @Produce json
// @Param shop body Shop true "店铺信息"
// @Success 200 {object} response.Response{data=Shop} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /shops [post]
func (h *Handler) Create(c *gin.Context) {
	var shop Shop
	if err := c.ShouldBindJSON(&shop); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Create(&shop).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建店铺失败")
		return
	}

	response.Success(c, shop)
}

// List 获取店铺列表
// @Summary 获取店铺列表
// @Description 获取所有店铺的列表
// @Tags 店铺管理
// @Accept json
// @Produce json
// @Param supplier_id query string false "供应商ID"
// @Param is_enabled query string false "是否启用"
// @Success 200 {object} response.Response{data=[]Shop} "获取成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /shops [get]
func (h *Handler) List(c *gin.Context) {
	var shops []Shop
	query := h.db.Model(&Shop{})

	// 支持按供应商ID筛选
	if supplierID := c.Query("supplier_id"); supplierID != "" {
		query = query.Where("supplier_id = ?", supplierID)
	}

	// 支持按启用状态筛选
	if isEnabled := c.Query("is_enabled"); isEnabled != "" {
		query = query.Where("is_enabled = ?", isEnabled)
	}

	if err := query.Find(&shops).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取店铺列表失败")
		return
	}

	response.Success(c, shops)
}

// Get 获取单个店铺
// @Summary 获取单个店铺
// @Description 根据ID获取店铺信息
// @Tags 店铺管理
// @Accept json
// @Produce json
// @Param id path int true "店铺ID"
// @Success 200 {object} response.Response{data=Shop} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "店铺不存在"
// @Router /shops/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var shop Shop
	if err := h.db.First(&shop, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "店铺不存在")
		return
	}

	response.Success(c, shop)
}

// Update 更新店铺
// @Summary 更新店铺
// @Description 更新店铺信息
// @Tags 店铺管理
// @Accept json
// @Produce json
// @Param id path int true "店铺ID"
// @Param shop body Shop true "店铺信息"
// @Success 200 {object} response.Response{data=Shop} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "店铺不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /shops/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var shop Shop
	if err := h.db.First(&shop, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "店铺不存在")
		return
	}

	if err := c.ShouldBindJSON(&shop); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&shop).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新店铺失败")
		return
	}

	response.Success(c, shop)
}

// Delete 删除店铺
// @Summary 删除店铺
// @Description 删除店铺
// @Tags 店铺管理
// @Accept json
// @Produce json
// @Param id path int true "店铺ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /shops/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.db.Delete(&Shop{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "删除店铺失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

// ToggleStatus 切换店铺状态
// @Summary 切换店铺状态
// @Description 启用/禁用店铺
// @Tags 店铺管理
// @Accept json
// @Produce json
// @Param id path int true "店铺ID"
// @Success 200 {object} response.Response{data=Shop} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "店铺不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /shops/{id}/toggle [patch]
func (h *Handler) ToggleStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var shop Shop
	if err := h.db.First(&shop, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "店铺不存在")
		return
	}

	shop.IsEnabled = !shop.IsEnabled
	if err := h.db.Save(&shop).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新状态失败")
		return
	}

	response.Success(c, shop)
}
