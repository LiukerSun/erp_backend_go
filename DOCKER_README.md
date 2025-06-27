# Docker 部署说明

## 快速开始

### 1. 构建并启动服务
```bash
docker-compose up --build
```

### 2. 后台运行
```bash
docker-compose up -d --build
```

### 3. 停止服务
```bash
docker-compose down
```

## 环境变量配置

### 必需配置
在 `docker-compose.yml` 中修改以下数据库配置：

```yaml
environment:
  - DB_HOST=your-database-host
  - DB_PORT=5432
  - DB_USER=your-database-user
  - DB_PASSWORD=your-database-password
  - DB_NAME=your-database-name
  - DB_SSLMODE=disable
```

### 可选配置

#### 跳过数据库迁移
如果您的数据库表结构已经存在，可以跳过迁移：
```yaml
environment:
  - SKIP_MIGRATION=true
```

#### 跳过种子数据初始化
如果您的数据库中已有初始数据，可以跳过种子数据创建：
```yaml
environment:
  - SKIP_SEED=true
```

#### 同时跳过迁移和种子数据
```yaml
environment:
  - SKIP_MIGRATION=true
  - SKIP_SEED=true
```

## 常见使用场景

### 1. 全新部署
```bash
# 默认行为：执行迁移和种子数据初始化
docker-compose up --build
```

### 2. 已有数据库结构，只需要种子数据
```yaml
environment:
  - SKIP_MIGRATION=true
  # 不设置 SKIP_SEED，会执行种子数据初始化
```

### 3. 已有完整数据，跳过所有初始化
```yaml
environment:
  - SKIP_MIGRATION=true
  - SKIP_SEED=true
```

### 4. 开发环境（使用本地数据库）
```yaml
environment:
  - DB_HOST=host.docker.internal  # 连接宿主机数据库
  - DB_PORT=5432
  - DB_USER=postgres
  - DB_PASSWORD=your-password
  - DB_NAME=erp_db
  - GIN_MODE=debug
```

## 日志查看

### 查看实时日志
```bash
docker-compose logs -f erp_backend
```

### 查看特定时间段的日志
```bash
docker-compose logs --since="2024-01-01T00:00:00" erp_backend
```

## 故障排除

### 1. 数据库连接失败
- 检查数据库配置是否正确
- 确保数据库服务正在运行
- 检查防火墙设置

### 2. 端口冲突
如果8080端口被占用，修改端口映射：
```yaml
ports:
  - "8081:8080"  # 使用8081端口
```

### 3. 权限问题
确保Docker有足够的权限访问数据库和文件系统。

## 生产环境建议

1. **使用强密码**：修改所有默认密码
2. **设置JWT密钥**：使用强随机字符串
3. **配置SSL**：设置 `DB_SSLMODE=require`
4. **限制网络访问**：只允许必要的端口
5. **定期备份**：设置数据库备份策略 