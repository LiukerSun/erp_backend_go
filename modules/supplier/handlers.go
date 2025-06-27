package supplier

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

// Create 创建供应商
// @Summary 创建供应商
// @Description 创建新的供应商
// @Tags 供应商管理
// @Accept json
// @Produce json
// @Param supplier body Supplier true "供应商信息"
// @Success 200 {object} response.Response{data=Supplier} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /suppliers [post]
func (h *Handler) Create(c *gin.Context) {
	var supplier Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Create(&supplier).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建供应商失败")
		return
	}

	response.Success(c, supplier)
}

// List 获取供应商列表
// @Summary 获取供应商列表
// @Description 获取所有供应商的列表
// @Tags 供应商管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]Supplier} "获取成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /suppliers [get]
func (h *Handler) List(c *gin.Context) {
	var suppliers []Supplier
	if err := h.db.Find(&suppliers).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取供应商列表失败")
		return
	}

	response.Success(c, suppliers)
}

// Get 获取单个供应商
// @Summary 获取单个供应商
// @Description 根据ID获取供应商信息
// @Tags 供应商管理
// @Accept json
// @Produce json
// @Param id path int true "供应商ID"
// @Success 200 {object} response.Response{data=Supplier} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "供应商不存在"
// @Router /suppliers/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var supplier Supplier
	if err := h.db.First(&supplier, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "供应商不存在")
		return
	}

	response.Success(c, supplier)
}

// Update 更新供应商
// @Summary 更新供应商
// @Description 更新供应商信息
// @Tags 供应商管理
// @Accept json
// @Produce json
// @Param id path int true "供应商ID"
// @Param supplier body Supplier true "供应商信息"
// @Success 200 {object} response.Response{data=Supplier} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "供应商不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /suppliers/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var supplier Supplier
	if err := h.db.First(&supplier, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "供应商不存在")
		return
	}

	if err := c.ShouldBindJSON(&supplier); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&supplier).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新供应商失败")
		return
	}

	response.Success(c, supplier)
}

// Delete 删除供应商
// @Summary 删除供应商
// @Description 删除供应商
// @Tags 供应商管理
// @Accept json
// @Produce json
// @Param id path int true "供应商ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /suppliers/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.db.Delete(&Supplier{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "删除供应商失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

// ToggleStatus 切换供应商状态
// @Summary 切换供应商状态
// @Description 启用/禁用供应商
// @Tags 供应商管理
// @Accept json
// @Produce json
// @Param id path int true "供应商ID"
// @Success 200 {object} response.Response{data=Supplier} "切换成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "供应商不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /suppliers/{id}/toggle [patch]
func (h *Handler) ToggleStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var supplier Supplier
	if err := h.db.First(&supplier, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "供应商不存在")
		return
	}

	supplier.IsEnabled = !supplier.IsEnabled
	if err := h.db.Save(&supplier).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新状态失败")
		return
	}

	response.Success(c, supplier)
}
