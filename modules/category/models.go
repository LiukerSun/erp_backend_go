package category

import (
	"time"
)

// Category 分类模型
// @Description 分类信息
type Category struct {
	ID          uint       `gorm:"primarykey" json:"id"`                                // 主键ID
	CreatedAt   time.Time  `json:"created_at"`                                          // 创建时间
	UpdatedAt   time.Time  `json:"updated_at"`                                          // 更新时间
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`                             // 删除时间
	Name        string     `gorm:"type:varchar(100);not null;comment:分类名称" json:"name"` // 分类名称
	Description string     `gorm:"type:text;comment:分类描述" json:"description"`           // 分类描述
	ParentID    *uint      `gorm:"comment:父级分类ID" json:"parent_id"`                     // 父级分类ID
	LevelRemark string     `gorm:"type:text;comment:层级备注" json:"level_remark"`          // 层级备注
	IsEnabled   bool       `gorm:"default:true;comment:是否启用" json:"is_enabled"`         // 是否启用
}
