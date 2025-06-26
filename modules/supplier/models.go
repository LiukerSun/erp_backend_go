package supplier

import (
	"time"
)

// Supplier 供应商模型
// @Description 供应商信息
type Supplier struct {
	ID        uint       `gorm:"primarykey" json:"id"`                                 // 主键ID
	CreatedAt time.Time  `json:"created_at"`                                           // 创建时间
	UpdatedAt time.Time  `json:"updated_at"`                                           // 更新时间
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`                              // 删除时间
	Name      string     `gorm:"type:varchar(100);not null;comment:供应商名称" json:"name"` // 供应商名称
	Remark    string     `gorm:"type:text;comment:供应商备注" json:"remark"`                // 供应商备注
	IsEnabled bool       `gorm:"default:true;comment:是否启用" json:"is_enabled"`          // 是否启用
}
