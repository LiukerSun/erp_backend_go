package database

import (
	"log"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(models ...interface{}) error {
	log.Println("开始数据库迁移...")

	// 获取数据库实例
	db := GetDB()

	// 迁移所有模型
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Printf("迁移表失败: %v", err)
			return err
		}
	}

	log.Println("数据库迁移完成")
	return nil
}
