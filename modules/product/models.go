package product

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// DynamicAttributes 动态属性类型
type DynamicAttributes map[string]interface{}

// Value 实现 driver.Valuer 接口
func (da DynamicAttributes) Value() (driver.Value, error) {
	return json.Marshal(da)
}

// Scan 实现 sql.Scanner 接口
func (da *DynamicAttributes) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("类型断言失败")
	}

	return json.Unmarshal(bytes, &da)
}

// Product 商品模型
// @Description 商品信息
type Product struct {
	ID           uint              `gorm:"primarykey" json:"id"`                                  // 主键ID
	CreatedAt    time.Time         `json:"created_at"`                                            // 创建时间
	UpdatedAt    time.Time         `json:"updated_at"`                                            // 更新时间
	DeletedAt    *time.Time        `gorm:"index" json:"deleted_at"`                               // 删除时间
	SupplierID   uint              `gorm:"not null;comment:供应商ID" json:"supplier_id"`             // 供应商ID
	CategoryID   uint              `gorm:"not null;comment:分类ID" json:"category_id"`              // 分类ID
	Name         string            `gorm:"type:varchar(200);not null;comment:商品名称" json:"name"`   // 商品名称
	SKU          string            `gorm:"type:varchar(50);uniqueIndex;comment:商品SKU" json:"sku"` // 商品SKU
	Type         int               `gorm:"comment:商品类型" json:"type"`                              // 商品类型
	Price        float64           `gorm:"type:decimal(10,2);comment:商品价格" json:"price"`          // 商品价格
	Stock        int               `gorm:"comment:商品库存" json:"stock"`                             // 商品库存
	DynamicAttrs DynamicAttributes `gorm:"type:json;comment:动态属性" json:"dynamic_attrs"`           // 动态属性
	Remark       string            `gorm:"type:text;comment:商品备注" json:"remark"`                  // 商品备注
	IsEnabled    bool              `gorm:"default:true;comment:是否启用" json:"is_enabled"`           // 是否启用
}
