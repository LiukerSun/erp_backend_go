package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"erp_backend/pkg/middleware"
	"erp_backend/pkg/response"
)

type Handler struct {
	db *gorm.DB
}

// LoginResponse 登录响应
// @Description 登录成功后的响应数据
type LoginResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."` // JWT令牌
	User  UserResponse `json:"user"`                                                    // 用户信息
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录并返回JWT token
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param data body LoginRequest true "登录信息"
// @Success 200 {object} response.Response{data=LoginResponse} "登录成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "用户名或密码错误"
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var loginData LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	var user User
	if err := h.db.Where("name = ?", loginData.Username).First(&user).Error; err != nil {
		response.Error(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		response.Error(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// 生成JWT token
	tokenString, err := middleware.GenerateToken(user.ID, user.UserType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "生成token失败")
		return
	}

	response.Success(c, LoginResponse{
		Token: tokenString,
		User:  user.ToResponse(),
	})
}

// Register 用户注册
// @Summary 用户注册
// @Description 注册新用户
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param data body CreateUserRequest true "用户信息"
// @Success 200 {object} response.Response{data=UserResponse} "注册成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	// 检查用户名是否已存在
	var count int64
	h.db.Model(&User{}).Where("name = ?", user.Name).Count(&count)
	if count > 0 {
		response.Error(c, http.StatusBadRequest, "用户名已存在")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "密码加密失败")
		return
	}
	user.Password = string(hashedPassword)

	// 设置默认用户类型
	if user.UserType == "" {
		user.UserType = "user"
	}

	if err := h.db.Create(&user).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建用户失败")
		return
	}

	response.Success(c, user)
}

// List 获取用户列表
// @Summary 获取用户列表
// @Description 获取所有用户的列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=[]UserResponse} "获取成功"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /users [get]
func (h *Handler) List(c *gin.Context) {
	var users []User
	if err := h.db.Find(&users).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "获取用户列表失败")
		return
	}

	response.Success(c, users)
}

// Get 获取单个用户
// @Summary 获取单个用户
// @Description 根据ID获取用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response{data=UserResponse} "获取成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "用户不存在"
// @Router /users/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var user User
	if err := h.db.First(&user, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	response.Success(c, user)
}

// Create 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body CreateUserRequest true "用户信息"
// @Success 200 {object} response.Response{data=UserResponse} "创建成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /users [post]
func (h *Handler) Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "密码加密失败")
		return
	}
	user.Password = string(hashedPassword)

	if err := h.db.Create(&user).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "创建用户失败")
		return
	}

	response.Success(c, user)
}

// Update 更新用户
// @Summary 更新用户
// @Description 更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Param data body UpdateUserRequest true "用户信息"
// @Success 200 {object} response.Response{data=UserResponse} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "用户不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /users/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	var user User
	if err := h.db.First(&user, id).Error; err != nil {
		response.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&user).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新用户失败")
		return
	}

	response.Success(c, user)
}

// Delete 删除用户
// @Summary 删除用户
// @Description 删除指定用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} response.Response "删除成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /users/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "无效的ID")
		return
	}

	if err := h.db.Delete(&User{}, id).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "删除用户失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

// GetProfile 获取用户个人资料
// @Summary 获取个人资料
// @Description 获取当前登录用户的个人资料
// @Tags 个人中心
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=UserResponse} "获取成功"
// @Failure 401 {object} response.Response "未登录"
// @Failure 404 {object} response.Response "用户不存在"
// @Router /users/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "未认证")
		return
	}

	var user User
	if err := h.db.First(&user, userID).Error; err != nil {
		response.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	response.Success(c, user)
}

// UpdateProfile 更新个人资料
// @Summary 更新个人资料
// @Description 更新当前登录用户的个人资料
// @Tags 个人中心
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body UpdateUserRequest true "用户信息"
// @Success 200 {object} response.Response{data=UserResponse} "更新成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "未登录"
// @Failure 404 {object} response.Response "用户不存在"
// @Router /users/profile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "未认证")
		return
	}

	var user User
	if err := h.db.First(&user, userID).Error; err != nil {
		response.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	if err := h.db.Save(&user).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新个人资料失败")
		return
	}

	response.Success(c, user)
}

// UpdatePassword 修改密码
// @Summary 修改密码
// @Description 修改当前登录用户的密码
// @Tags 个人中心
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body UpdatePasswordRequest true "密码信息"
// @Success 200 {object} response.Response "修改成功"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 401 {object} response.Response "未登录或旧密码错误"
// @Failure 404 {object} response.Response "用户不存在"
// @Router /users/password [put]
func (h *Handler) UpdatePassword(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "未认证")
		return
	}

	var passwordData struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&passwordData); err != nil {
		response.Error(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	var user User
	if err := h.db.First(&user, userID).Error; err != nil {
		response.Error(c, http.StatusNotFound, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordData.OldPassword)); err != nil {
		response.Error(c, http.StatusBadRequest, "原密码错误")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordData.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "密码加密失败")
		return
	}

	user.Password = string(hashedPassword)
	if err := h.db.Save(&user).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "更新密码失败")
		return
	}

	response.Success(c, gin.H{"message": "密码更新成功"})
}
