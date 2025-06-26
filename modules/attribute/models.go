package attribute

import (
	"time"
)

// Attribute 属性模型
// @Description 属性信息
type Attribute struct {
	ID         uint       `gorm:"primarykey" json:"id"`                                    // 主键ID
	CreatedAt  time.Time  `json:"created_at"`                                              // 创建时间
	UpdatedAt  time.Time  `json:"updated_at"`                                              // 更新时间
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at"`                                 // 删除时间
	Name       string     `gorm:"type:varchar(100);not null;comment:属性名称" json:"name"`     // 属性名称
	DataType   string     `gorm:"type:varchar(50);not null;comment:数据类型" json:"data_type"` // 数据类型
	CategoryID uint       `gorm:"not null;comment:所属分类ID" json:"category_id"`              // 所属分类ID
	IsRequired bool       `gorm:"default:false;comment:是否必填" json:"is_required"`           // 是否必填
	Remark     string     `gorm:"type:text;comment:属性备注" json:"remark"`                    // 属性备注
	IsEnabled  bool       `gorm:"default:true;comment:是否启用" json:"is_enabled"`             // 是否启用
}

// ProductAttribute 商品属性值模型
// @Description 商品属性值信息
type ProductAttribute struct {
	ID          uint       `gorm:"primarykey" json:"id"`                      // 主键ID
	CreatedAt   time.Time  `json:"created_at"`                                // 创建时间
	UpdatedAt   time.Time  `json:"updated_at"`                                // 更新时间
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`                   // 删除时间
	ProductID   uint       `gorm:"not null;comment:商品ID" json:"product_id"`   // 商品ID
	AttributeID uint       `gorm:"not null;comment:属性ID" json:"attribute_id"` // 属性ID
	Value       string     `gorm:"type:text;comment:属性值" json:"value"`        // 属性值
}
