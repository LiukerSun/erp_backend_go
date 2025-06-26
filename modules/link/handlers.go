package link

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"erp-backend/pkg/response"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// Create 创建链接
// @Summary 创建链接
// @Description 创建新的链接
// @Tags 链接管理
// @Accept json
// @Produce json
// @Param link body Link true "链接信息"
// @Success 200 {object} response.Response{data=Link} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /links [post]
func (h *Handler) Create(c *gin.Context) {
	var link Link
	if err := c.ShouldBindJSON(&link); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Create(&link).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建链接失败")
		return
	}

	response.Success(c, link)
}

// List 获取链接列表
// @Summary 获取链接列表
// @Description 获取所有链接的列表
// @Tags 链接管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]Link} "获取成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /links [get]
func (h *Handler) List(c *gin.Context) {
	var links []Link
	query := h.db.Model(&Link{})

	// 支持按店铺ID筛选
	if shopID := c.Query("shop_id"); shopID != "" {
		query = query.Where("shop_id = ?", shopID)
	}

	// 支持按分类ID筛选
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// 支持按启用状态筛选
	if isEnabled := c.Query("is_enabled"); isEnabled != "" {
		query = query.Where("is_enabled = ?", isEnabled)
	}

	if err := query.Find(&links).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取链接列表失败")
		return
	}

	response.Success(c, links)
}

// Get 获取单个链接
// @Summary 获取单个链接
// @Description 根据ID获取链接信息
// @Tags 链接管理
// @Accept json
// @Produce json
// @Param id path int true "链接ID"
// @Success 200 {object} response.Response{data=Link} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "链接不存在"
// @Router /links/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var link Link
	if err := h.db.First(&link, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "链接不存在")
		return
	}

	response.Success(c, link)
}

// Update 更新链接
// @Summary 更新链接
// @Description 更新链接信息
// @Tags 链接管理
// @Accept json
// @Produce json
// @Param id path int true "链接ID"
// @Param link body Link true "链接信息"
// @Success 200 {object} response.Response{data=Link} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "链接不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /links/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var link Link
	if err := h.db.First(&link, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "链接不存在")
		return
	}

	if err := c.ShouldBindJSON(&link); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&link).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新链接失败")
		return
	}

	response.Success(c, link)
}

// Delete 删除链接
// @Summary 删除链接
// @Description 删除链接
// @Tags 链接管理
// @Accept json
// @Produce json
// @Param id path int true "链接ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /links/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.db.Delete(&Link{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "删除链接失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

// ToggleStatus 切换链接状态
// @Summary 切换链接状态
// @Description 启用/禁用链接
// @Tags 链接管理
// @Accept json
// @Produce json
// @Param id path int true "链接ID"
// @Success 200 {object} response.Response{data=Link} "切换成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "链接不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /links/{id}/toggle [patch]
func (h *Handler) ToggleStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var link Link
	if err := h.db.First(&link, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "链接不存在")
		return
	}

	link.IsEnabled = !link.IsEnabled
	if err := h.db.Save(&link).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新状态失败")
		return
	}

	response.Success(c, link)
}
