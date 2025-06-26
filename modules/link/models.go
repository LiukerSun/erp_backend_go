package link

import (
	"time"
)

// Link 链接模型
// @Description 链接信息
type Link struct {
	ID         uint       `gorm:"primarykey" json:"id"`                                // 主键ID
	CreatedAt  time.Time  `json:"created_at"`                                          // 创建时间
	UpdatedAt  time.Time  `json:"updated_at"`                                          // 更新时间
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at"`                             // 删除时间
	Name       string     `gorm:"type:varchar(100);not null;comment:链接名称" json:"name"` // 链接名称
	URL        string     `gorm:"type:varchar(500);not null;comment:链接地址" json:"url"`  // 链接地址
	BaseRemark string     `gorm:"type:text;comment:基础备注" json:"base_remark"`           // 基础备注
	ShopID     uint       `gorm:"not null;comment:店铺ID" json:"shop_id"`                // 店铺ID
	CategoryID uint       `gorm:"not null;comment:类目ID" json:"category_id"`            // 类目ID
	Remark     string     `gorm:"type:text;comment:链接备注" json:"remark"`                // 链接备注
	IsEnabled  bool       `gorm:"default:true;comment:是否启用" json:"is_enabled"`         // 是否启用
}
