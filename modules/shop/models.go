package shop

import (
	"time"
)

// Shop 店铺模型
// @Description 店铺信息
type Shop struct {
	ID         uint       `gorm:"primarykey" json:"id"`                                // 主键ID
	CreatedAt  time.Time  `json:"created_at"`                                          // 创建时间
	UpdatedAt  time.Time  `json:"updated_at"`                                          // 更新时间
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at"`                             // 删除时间
	SupplierID uint       `gorm:"not null;comment:所属供应商ID" json:"supplier_id"`         // 所属供应商ID
	Name       string     `gorm:"type:varchar(100);not null;comment:店铺名称" json:"name"` // 店铺名称
	Remark     string     `gorm:"type:text;comment:店铺备注" json:"remark"`                // 店铺备注
	IsEnabled  bool       `gorm:"default:true;comment:是否启用" json:"is_enabled"`         // 是否启用
}
