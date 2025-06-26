package user

import (
	"errors"

	"erp-backend/pkg/database"

	"gorm.io/gorm"
)

// Repository 用户数据访问接口
type Repository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByName(name string) (*User, error)
	GetAll() ([]User, error)
	Update(user *User) error
	Delete(id uint) error
}

// repository 用户数据访问实现
type repository struct {
	db *gorm.DB
}

// NewRepository 创建用户repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create 创建用户
func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

// GetByID 根据ID获取用户
func (r *repository) GetByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱获取用户
func (r *repository) GetByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByName 根据名称获取用户
func (r *repository) GetByName(name string) (*User, error) {
	var user User
	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetAll 获取所有用户
func (r *repository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

// Update 更新用户
func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
