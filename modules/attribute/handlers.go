package attribute

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

// CreateAttribute 创建属性
// @Summary 创建属性
// @Description 创建新的属性
// @Tags 属性管理
// @Accept json
// @Produce json
// @Param attribute body Attribute true "属性信息"
// @Success 200 {object} response.Response{data=Attribute} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /attributes [post]
func (h *Handler) CreateAttribute(c *gin.Context) {
	var attribute Attribute
	if err := c.ShouldBindJSON(&attribute); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Create(&attribute).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建属性失败")
		return
	}

	response.Success(c, attribute)
}

// ListAttributes 获取属性列表
// @Summary 获取属性列表
// @Description 获取所有属性的列表
// @Tags 属性管理
// @Accept json
// @Produce json
// @Param category_id query string false "分类ID"
// @Param is_enabled query string false "是否启用"
// @Success 200 {object} response.Response{data=[]Attribute} "获取成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /attributes [get]
func (h *Handler) ListAttributes(c *gin.Context) {
	var attributes []Attribute
	query := h.db.Model(&Attribute{})

	// 支持按分类ID筛选
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// 支持按启用状态筛选
	if isEnabled := c.Query("is_enabled"); isEnabled != "" {
		query = query.Where("is_enabled = ?", isEnabled)
	}

	if err := query.Find(&attributes).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取属性列表失败")
		return
	}

	response.Success(c, attributes)
}

// GetAttribute 获取单个属性
// @Summary 获取单个属性
// @Description 根据ID获取属性信息
// @Tags 属性管理
// @Accept json
// @Produce json
// @Param id path int true "属性ID"
// @Success 200 {object} response.Response{data=Attribute} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "属性不存在"
// @Router /attributes/{id} [get]
func (h *Handler) GetAttribute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var attribute Attribute
	if err := h.db.First(&attribute, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "属性不存在")
		return
	}

	response.Success(c, attribute)
}

// UpdateAttribute 更新属性
// @Summary 更新属性
// @Description 更新属性信息
// @Tags 属性管理
// @Accept json
// @Produce json
// @Param id path int true "属性ID"
// @Param attribute body Attribute true "属性信息"
// @Success 200 {object} response.Response{data=Attribute} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "属性不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /attributes/{id} [put]
func (h *Handler) UpdateAttribute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var attribute Attribute
	if err := h.db.First(&attribute, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "属性不存在")
		return
	}

	if err := c.ShouldBindJSON(&attribute); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&attribute).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新属性失败")
		return
	}

	response.Success(c, attribute)
}

// DeleteAttribute 删除属性
// @Summary 删除属性
// @Description 删除属性
// @Tags 属性管理
// @Accept json
// @Produce json
// @Param id path int true "属性ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /attributes/{id} [delete]
func (h *Handler) DeleteAttribute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.db.Delete(&Attribute{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "删除属性失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

// ToggleAttributeStatus 切换属性状态
// @Summary 切换属性状态
// @Description 启用/禁用属性
// @Tags 属性管理
// @Accept json
// @Produce json
// @Param id path int true "属性ID"
// @Success 200 {object} response.Response{data=Attribute} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "属性不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /attributes/{id}/toggle [patch]
func (h *Handler) ToggleAttributeStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var attribute Attribute
	if err := h.db.First(&attribute, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "属性不存在")
		return
	}

	attribute.IsEnabled = !attribute.IsEnabled
	if err := h.db.Save(&attribute).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新状态失败")
		return
	}

	response.Success(c, attribute)
}

// CreateProductAttribute 创建商品属性值
// @Summary 创建商品属性值
// @Description 创建新的商品属性值
// @Tags 商品属性值管理
// @Accept json
// @Produce json
// @Param productAttribute body ProductAttribute true "商品属性值信息"
// @Success 200 {object} response.Response{data=ProductAttribute} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /product-attributes [post]
func (h *Handler) CreateProductAttribute(c *gin.Context) {
	var productAttribute ProductAttribute
	if err := c.ShouldBindJSON(&productAttribute); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Create(&productAttribute).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建商品属性值失败")
		return
	}

	response.Success(c, productAttribute)
}

// ListProductAttributes 获取商品属性值列表
// @Summary 获取商品属性值列表
// @Description 获取所有商品属性值的列表
// @Tags 商品属性值管理
// @Accept json
// @Produce json
// @Param product_id query string false "商品ID"
// @Param attribute_id query string false "属性ID"
// @Success 200 {object} response.Response{data=[]ProductAttribute} "获取成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /product-attributes [get]
func (h *Handler) ListProductAttributes(c *gin.Context) {
	var productAttributes []ProductAttribute
	query := h.db.Model(&ProductAttribute{})

	// 支持按商品ID筛选
	if productID := c.Query("product_id"); productID != "" {
		query = query.Where("product_id = ?", productID)
	}

	// 支持按属性ID筛选
	if attributeID := c.Query("attribute_id"); attributeID != "" {
		query = query.Where("attribute_id = ?", attributeID)
	}

	if err := query.Find(&productAttributes).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取商品属性值列表失败")
		return
	}

	response.Success(c, productAttributes)
}

// UpdateProductAttribute 更新商品属性值
// @Summary 更新商品属性值
// @Description 更新商品属性值信息
// @Tags 商品属性值管理
// @Accept json
// @Produce json
// @Param id path int true "商品属性值ID"
// @Param productAttribute body ProductAttribute true "商品属性值信息"
// @Success 200 {object} response.Response{data=ProductAttribute} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "商品属性值不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /product-attributes/{id} [put]
func (h *Handler) UpdateProductAttribute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var productAttribute ProductAttribute
	if err := h.db.First(&productAttribute, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "商品属性值不存在")
		return
	}

	if err := c.ShouldBindJSON(&productAttribute); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&productAttribute).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新商品属性值失败")
		return
	}

	response.Success(c, productAttribute)
}

// DeleteProductAttribute 删除商品属性值
// @Summary 删除商品属性值
// @Description 删除商品属性值
// @Tags 商品属性值管理
// @Accept json
// @Produce json
// @Param id path int true "商品属性值ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /product-attributes/{id} [delete]
func (h *Handler) DeleteProductAttribute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.db.Delete(&ProductAttribute{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "删除商品属性值失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}
