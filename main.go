package main

import (
	"log"
	"os"

	"erp-backend/modules/attribute"
	"erp-backend/modules/category"
	"erp-backend/modules/link"
	"erp-backend/modules/product"
	"erp-backend/modules/shop"
	"erp-backend/modules/supplier"
	"erp-backend/modules/system"
	"erp-backend/modules/user"
	"erp-backend/pkg/database"
	"erp-backend/pkg/middleware"
	"erp-backend/pkg/response"

	_ "erp-backend/docs" // 导入 swagger docs

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @title ERP 系统 API
// @version 1.0
// @description ERP 系统后端 API 文档
// @host localhost:8080
// @BasePath /api/v1

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description 请在此输入 Bearer token: Bearer {token}
func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用默认配置")
	}

	// 设置Gin模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化数据库
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 自动迁移数据库结构
	if err := autoMigrate(db); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化种子数据
	if err := seedData(db); err != nil {
		log.Fatal("种子数据初始化失败:", err)
	}

	// 创建Gin引擎
	r := gin.Default()

	// 添加中间件
	r.Use(middleware.CORSMiddleware())

	// 设置路由
	setupRoutes(r, db)

	// 添加 Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 获取端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务器
	log.Printf("服务器启动在端口 %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}

// autoMigrate 自动迁移数据库结构
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&supplier.Supplier{},
		&shop.Shop{},
		&product.Product{},
		&category.Category{},
		&link.Link{},
		&attribute.Attribute{},
		&attribute.ProductAttribute{},
	)
}

// seedData 初始化种子数据
func seedData(db *gorm.DB) error {
	log.Println("开始初始化种子数据...")

	// 检查是否已有数据
	var userCount int64
	db.Model(&user.User{}).Count(&userCount)
	if userCount > 0 {
		log.Println("用户表已有数据，跳过种子数据初始化")
		return nil
	}

	// 创建管理员用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("test1234"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("密码加密失败: %v", err)
		return err
	}

	adminUser := user.User{
		Name:     "evansun",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
		UserType: "admin",
		IsDelete: false,
	}

	if err := db.Create(&adminUser).Error; err != nil {
		log.Printf("创建管理员用户失败: %v", err)
		return err
	}

	log.Println("管理员用户创建成功")
	return nil
}

// setupRoutes 设置路由
func setupRoutes(r *gin.Engine, db *gorm.DB) {
	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 系统模块路由
		system.RegisterRoutes(v1, db)

		// 用户模块路由
		user.RegisterRoutes(v1, db)

		// 供应商模块路由
		supplier.RegisterRoutes(v1, db)

		// 店铺模块路由
		shop.RegisterRoutes(v1, db)

		// 商品模块路由
		product.RegisterRoutes(v1, db)

		// 分类模块路由
		category.RegisterRoutes(v1, db)

		// 链接模块路由
		link.RegisterRoutes(v1, db)

		// 属性模块路由
		attribute.RegisterRoutes(v1, db)
	}

	// 根路径
	r.GET("/", func(c *gin.Context) {
		response.Success(c, gin.H{
			"message": "ERP后端服务运行正常",
			"version": "1.0.0",
		})
	})
}
