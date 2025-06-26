package product

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

// Create 创建商品
// @Summary 创建商品
// @Description 创建新的商品
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param product body Product true "商品信息"
// @Success 200 {object} response.Response{data=Product} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /products [post]
func (h *Handler) Create(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Create(&product).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建商品失败")
		return
	}

	response.Success(c, product)
}

// List 获取商品列表
// @Summary 获取商品列表
// @Description 获取所有商品的列表
// @Tags 商品管理
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]Product} "获取成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /products [get]
func (h *Handler) List(c *gin.Context) {
	var products []Product
	query := h.db.Model(&Product{})

	// 支持按供应商ID筛选
	if supplierID := c.Query("supplier_id"); supplierID != "" {
		query = query.Where("supplier_id = ?", supplierID)
	}

	// 支持按分类ID筛选
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// 支持按商品类型筛选
	if productType := c.Query("type"); productType != "" {
		query = query.Where("type = ?", productType)
	}

	// 支持按启用状态筛选
	if isEnabled := c.Query("is_enabled"); isEnabled != "" {
		query = query.Where("is_enabled = ?", isEnabled)
	}

	if err := query.Find(&products).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取商品列表失败")
		return
	}

	response.Success(c, products)
}

// Get 获取单个商品
// @Summary 获取单个商品
// @Description 根据ID获取商品信息
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} response.Response{data=Product} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "商品不存在"
// @Router /products/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var product Product
	if err := h.db.First(&product, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	response.Success(c, product)
}

// Update 更新商品
// @Summary 更新商品
// @Description 更新商品信息
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Param product body Product true "商品信息"
// @Success 200 {object} response.Response{data=Product} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "商品不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /products/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var product Product
	if err := h.db.First(&product, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&product).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新商品失败")
		return
	}

	response.Success(c, product)
}

// Delete 删除商品
// @Summary 删除商品
// @Description 删除商品
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /products/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.db.Delete(&Product{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "删除商品失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

// ToggleStatus 切换商品状态
// @Summary 切换商品状态
// @Description 启用/禁用商品
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} response.Response{data=Product} "切换成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "商品不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /products/{id}/toggle [patch]
func (h *Handler) ToggleStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var product Product
	if err := h.db.First(&product, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "商品不存在")
		return
	}

	product.IsEnabled = !product.IsEnabled
	if err := h.db.Save(&product).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新状态失败")
		return
	}

	response.Success(c, product)
}

// UpdateStock 更新库存
func (h *Handler) UpdateStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var stockUpdate struct {
		Stock int `json:"stock" binding:"required"`
	}

	if err := c.ShouldBindJSON(&stockUpdate); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Model(&Product{}).Where("id = ?", id).Update("stock", stockUpdate.Stock).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新库存失败")
		return
	}

	response.Success(c, gin.H{"message": "更新库存成功"})
}

// UpdatePrice 更新价格
func (h *Handler) UpdatePrice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var priceUpdate struct {
		Price float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&priceUpdate); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Model(&Product{}).Where("id = ?", id).Update("price", priceUpdate.Price).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新价格失败")
		return
	}

	response.Success(c, gin.H{"message": "更新价格成功"})
}
