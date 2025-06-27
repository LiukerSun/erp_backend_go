# ERP 系统后端

这是一个基于 Go 语言和 Gin 框架开发的 ERP 系统后端服务。

## 技术栈

- Go 1.21+
- Gin Web 框架
- GORM ORM 框架
- JWT 认证
- Swagger API 文档

## 项目结构

```
backend/
├── docs/               # Swagger 文档
├── modules/           # 业务模块
│   ├── attribute/    # 属性管理模块
│   ├── category/     # 分类管理模块
│   ├── link/         # 链接管理模块
│   ├── product/      # 商品管理模块
│   ├── shop/         # 店铺管理模块
│   ├── supplier/     # 供应商管理模块
│   ├── system/       # 系统模块
│   └── user/         # 用户模块
├── pkg/              # 公共包
│   ├── config/       # 配置
│   ├── database/     # 数据库
│   ├── middleware/   # 中间件
│   └── response/     # 响应处理
├── .env.example      # 环境变量示例
├── go.mod           # Go 模块文件
├── go.sum           # Go 依赖版本文件
├── main.go          # 主程序入口
└── README.md        # 项目说明文档
```

## 快速开始（二进制运行）

### 1. 下载

从 [Releases](https://github.com/your-username/erp_backend/releases) 页面下载适合您系统的最新版本：

- Windows: `erp_backend_windows_amd64.exe`
- Linux: `erp_backend_linux_amd64`
- macOS: `erp_backend_darwin_amd64`

### 2. 配置

1. 在二进制文件同目录下创建 `.env` 文件：
```env
# 服务器配置
PORT=8080
GIN_MODE=debug

# 数据库配置
DB_HOST=localhost
DB_PORT=5432
DB_NAME=erp_db
DB_USER=postgres
DB_PASSWORD=your_password
DB_SSLMODE=disable

# JWT配置
JWT_SECRET=your-secret-key
JWT_EXPIRE_HOURS=24
```

2. 确保您有一个运行中的 PostgreSQL (9.6+) 数据库服务
3. 创建数据库：
```sql
CREATE DATABASE erp_db;
```

### 3. 运行

#### Windows
```bash
# 直接运行
erp_backend_windows_amd64.exe

# 或者在命令提示符中运行
start erp_backend_windows_amd64.exe

# 后台运行
start /b erp_backend_windows_amd64.exe
```

#### Linux/macOS
```bash
# 添加执行权限
chmod +x erp_backend_linux_amd64

# 直接运行
./erp_backend_linux_amd64

# 后台运行
nohup ./erp_backend_linux_amd64 &

# 查看日志
tail -f nohup.out
```

服务器默认运行在 8080 端口。首次运行时会自动：
- 创建必要的数据库表
- 创建默认管理员账号

### 4. 默认管理员账号

系统初始化时会创建一个默认管理员账号：
- 用户名：evansun
- 密码：test1234

**请在首次登录后立即修改默认密码！**

## API 文档

服务启动后，访问 Swagger 文档：
```
http://localhost:8080/swagger/index.html
```

## 主要功能模块

### 1. 用户管理模块 (user)
- 用户认证与授权
- 用户信息管理
- 角色权限控制

### 2. 供应商管理模块 (supplier)
- 供应商信息管理
- 供应商状态控制
- 供应商关联管理

### 3. 店铺管理模块 (shop)
- 店铺信息管理
- 店铺与供应商关联
- 店铺状态控制

### 4. 链接管理模块 (link)
- 链接信息管理
- 链接分类关联
- 链接状态控制

### 5. 商品管理模块 (product)
- 商品基本信息管理
- 商品SKU管理
- 商品价格和库存管理
- 商品属性管理

### 6. 分类管理模块 (category)
- 商品分类管理
- 分类层级关系
- 分类属性关联

### 7. 属性管理模块 (attribute)
- 属性定义管理
- 属性值管理
- 属性与分类关联

### 8. 系统模块 (system)
- 系统配置管理
- 系统日志
- 基础数据维护

## 数据库架构设计

### 1. 供应商管理表 (suppliers)
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | Int | 主键(PK) |
| name | String | 供应商名称 |
| remark | String | 供应商备注 |
| is_active | Boolean | 供应商状态 |
| created_at | DateTime | 创建时间 |
| updated_at | DateTime | 更新时间 |

### 2. 店铺管理表 (shops)
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | Int | 主键(PK) |
| supplier_id | Int | 供应商ID(FK) |
| name | String | 店铺名称 |
| remark | String | 店铺备注 |
| is_active | Boolean | 店铺状态 |
| created_at | DateTime | 创建时间 |
| updated_at | DateTime | 更新时间 |

### 3. 链接管理表 (links)
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | Int | 主键(PK) |
| name | String | 链接名称 |
| url | String | 链接地址 |
| base_remark | String | 基础备注 |
| shop_id | Int | 店铺ID(FK) |
| category_id | Int | 分类ID(FK) |
| remark | String | 补充说明 |
| is_active | Boolean | 链接状态 |
| created_at | DateTime | 创建时间 |
| updated_at | DateTime | 更新时间 |

### 4. 商品管理表 (products)
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | Int | 主键(PK) |
| supplier_id | Int | 供应商ID(FK) |
| category_id | Int | 分类ID(FK) |
| name | String | 商品名称 |
| sku | String | 商品SKU |
| type | Int | 商品类型 |
| price | Decimal(10,2) | 商品价格 |
| stock | Int | 库存数量 |
| attributes | JSON | 动态属性 |
| remark | String | 商品备注 |
| is_active | Boolean | 商品状态 |
| created_at | DateTime | 创建时间 |
| updated_at | DateTime | 更新时间 |

### 5. 分类管理表 (categories)
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | Int | 主键(PK) |
| name | String | 分类名称 |
| parent_id | Int | 父分类ID |
| level | Int | 分类层级 |
| sort_order | Int | 排序序号 |
| is_active | Boolean | 分类状态 |
| created_at | DateTime | 创建时间 |
| updated_at | DateTime | 更新时间 |

### 6. 属性管理表 (attributes)
| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | Int | 主键(PK) |
| name | String | 属性名称 |
| type | String | 属性类型 |
| is_required | Boolean | 是否必填 |
| default_value | String | 默认值 |
| is_active | Boolean | 属性状态 |
| created_at | DateTime | 创建时间 |
| updated_at | DateTime | 更新时间 |

## 开发团队

- 后端开发：EvanSun
- 技术支持：RealX Team

## 版权信息

Copyright © 2024 RealX Team. All rights reserved.