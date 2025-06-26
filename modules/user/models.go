package user

import (
	"time"
)

// User 用户模型
// @Description 用户信息
type User struct {
	ID        uint       `gorm:"primarykey" json:"id"`                                           // 主键ID
	CreatedAt time.Time  `json:"created_at"`                                                     // 创建时间
	UpdatedAt time.Time  `json:"updated_at"`                                                     // 更新时间
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`                                        // 删除时间
	Name      string     `gorm:"type:varchar(100);not null;uniqueIndex;comment:用户名" json:"name"` // 用户名
	Email     string     `gorm:"type:varchar(100);uniqueIndex;comment:邮箱" json:"email"`          // 邮箱
	Password  string     `gorm:"type:varchar(100);not null;comment:密码" json:"-"`                 // 密码
	UserType  string     `gorm:"type:varchar(20);default:user;comment:用户类型" json:"user_type"`    // 用户类型（管理员、供应商、员工）
	IsDelete  bool       `gorm:"default:false;comment:是否删除" json:"is_delete"`                    // 是否删除
	Phone     string     `gorm:"size:20;comment:电话号码" json:"phone"`                              // 电话号码
}

// LoginRequest 登录请求
// @Description 用户登录的请求参数
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`       // 用户名
	Password string `json:"password" binding:"required" example:"password123"` // 密码
}

// CreateUserRequest 创建用户请求
// @Description 创建用户的请求参数
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" example:"张三"`                          // 用户名
	UserType string `json:"user_type" binding:"required,oneof=管理员 供应商 员工" example:"员工"`    // 用户类型
	Password string `json:"password" binding:"required,min=6" example:"123456"`            // 密码
	Email    string `json:"email" binding:"required,email" example:"zhangsan@example.com"` // 邮箱
	Phone    string `json:"phone" binding:"required" example:"13800138000"`                // 电话号码
}

// UpdateUserRequest 更新用户请求
// @Description 更新用户的请求参数
type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required" example:"张三"`                          // 用户名
	UserType string `json:"user_type" binding:"required,oneof=管理员 供应商 员工" example:"员工"`    // 用户类型
	Password string `json:"password,omitempty" binding:"omitempty,min=6" example:"123456"` // 密码（可选）
	Email    string `json:"email" binding:"required,email" example:"zhangsan@example.com"` // 邮箱
	Phone    string `json:"phone" binding:"required" example:"13800138000"`                // 电话号码
}

// UpdatePasswordRequest 修改密码请求
// @Description 修改密码的请求参数
type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" example:"123456"`       // 旧密码
	NewPassword string `json:"new_password" binding:"required,min=6" example:"654321"` // 新密码
}

// UserResponse 用户响应
// @Description 用户信息的响应格式
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`                                 // 用户ID
	Name      string    `json:"name" example:"张三"`                              // 用户名
	UserType  string    `json:"user_type" example:"员工"`                         // 用户类型
	Email     string    `json:"email" example:"zhangsan@example.com"`           // 邮箱
	Phone     string    `json:"phone" example:"13800138000"`                    // 电话号码
	IsDelete  bool      `json:"is_delete" example:"false"`                      // 是否删除
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00+08:00"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-01T00:00:00+08:00"` // 更新时间
}

// ToResponse 转换为响应格式
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		UserType:  u.UserType,
		Email:     u.Email,
		Phone:     u.Phone,
		IsDelete:  u.IsDelete,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
