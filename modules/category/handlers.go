package category

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

// Create 创建分类
// @Summary 创建分类
// @Description 创建新的分类
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param category body Category true "分类信息"
// @Success 200 {object} response.Response{data=Category} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /categories [post]
func (h *Handler) Create(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Create(&category).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建分类失败")
		return
	}

	response.Success(c, category)
}

// List 获取分类列表
// @Summary 获取分类列表
// @Description 获取所有分类的列表
// @Tags 分类管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]Category} "获取成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /categories [get]
func (h *Handler) List(c *gin.Context) {
	var categories []Category
	if err := h.db.Find(&categories).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取分类列表失败")
		return
	}

	response.Success(c, categories)
}

// Get 获取单个分类
// @Summary 获取单个分类
// @Description 根据ID获取分类信息
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Success 200 {object} response.Response{data=Category} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "分类不存在"
// @Router /categories/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var category Category
	if err := h.db.First(&category, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "分类不存在")
		return
	}

	response.Success(c, category)
}

// Update 更新分类
// @Summary 更新分类
// @Description 更新分类信息
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Param category body Category true "分类信息"
// @Success 200 {object} response.Response{data=Category} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "分类不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /categories/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var category Category
	if err := h.db.First(&category, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "分类不存在")
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&category).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新分类失败")
		return
	}

	response.Success(c, category)
}

// Delete 删除分类
// @Summary 删除分类
// @Description 删除分类
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /categories/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	// 检查是否有子分类
	var count int64
	if err := h.db.Model(&Category{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "检查子分类失败")
		return
	}

	if count > 0 {
		response.Error(c, http.StatusBadRequest, "该分类下存在子分类，无法删除")
		return
	}

	if err := h.db.Delete(&Category{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "删除分类失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

// ToggleStatus 切换分类状态
// @Summary 切换分类状态
// @Description 启用/禁用分类
// @Tags 分类管理
// @Accept json
// @Produce json
// @Param id path int true "分类ID"
// @Success 200 {object} response.Response{data=Category} "切换成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "分类不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /categories/{id}/toggle [patch]
func (h *Handler) ToggleStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var category Category
	if err := h.db.First(&category, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "分类不存在")
		return
	}

	category.IsEnabled = !category.IsEnabled
	if err := h.db.Save(&category).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新状态失败")
		return
	}

	response.Success(c, category)
}

// GetChildren 获取子分类
func (h *Handler) GetChildren(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var categories []Category
	if err := h.db.Where("parent_id = ?", id).Find(&categories).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取子分类失败")
		return
	}

	response.Success(c, categories)
}
