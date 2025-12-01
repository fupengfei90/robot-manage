# Docker 容器化部署指南

## 📋 前置要求

- Docker 20.10+
- Docker Compose 2.0+
- 至少 2GB 可用内存
- 至少 10GB 可用磁盘空间

## 🚀 快速开始

### 1. 克隆项目（如果还没有）

```bash
git clone <repository-url>
cd robot-manage
```

### 2. 配置环境变量

```bash
# 复制环境变量示例文件
cp .env.example .env

# 编辑 .env 文件，修改数据库密码和 JWT 密钥
vim .env
```

### 3. 一键部署

```bash
./deploy.sh
```

### 4. 访问应用

- 前端地址: http://localhost
- 后端API: http://localhost:8088
- 数据库: localhost:3306
- Redis: localhost:6379

## 📦 服务说明

### 服务列表

| 服务名 | 端口 | 说明 |
|--------|------|------|
| frontend | 80 | Vue3 前端应用 |
| backend | 8088 | Go 后端服务 |
| mysql | 3306 | MySQL 8.0 数据库 |
| redis | 6379 | Redis 缓存 |

### 数据持久化

所有数据都通过 Docker Volume 持久化：

- `mysql_data`: MySQL 数据
- `redis_data`: Redis 数据
- `backend_logs`: 后端日志

## 🔧 常用命令

### 启动服务

```bash
docker-compose up -d
```

### 停止服务

```bash
docker-compose down
```

### 查看日志

```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f backend
docker-compose logs -f frontend
```

### 重启服务

```bash
# 重启所有服务
docker-compose restart

# 重启特定服务
docker-compose restart backend
```

### 查看服务状态

```bash
docker-compose ps
```

### 进入容器

```bash
# 进入后端容器
docker-compose exec backend sh

# 进入数据库容器
docker-compose exec mysql mysql -u xedv -p
```

### 重新构建镜像

```bash
docker-compose build --no-cache
docker-compose up -d
```

## 🔄 更新部署

```bash
# 拉取最新代码
git pull

# 重新构建并启动
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

## 🗄️ 数据库管理

### 备份数据库

```bash
docker-compose exec mysql mysqldump -u xedv -p xedv > backup_$(date +%Y%m%d_%H%M%S).sql
```

### 恢复数据库

```bash
docker-compose exec -T mysql mysql -u xedv -p xedv < backup.sql
```

### 初始化数据库

数据库会在首次启动时自动执行 `backend/sql/` 目录下的 SQL 文件进行初始化。

## 🐛 故障排查

### 查看容器状态

```bash
docker-compose ps
```

### 查看详细日志

```bash
docker-compose logs -f --tail=100 backend
```

### 检查网络连接

```bash
docker network inspect robot-manage_robot-network
```

### 清理并重新部署

```bash
# 停止并删除所有容器、网络
docker-compose down

# 删除所有数据卷（注意：会删除所有数据）
docker-compose down -v

# 重新部署
./deploy.sh
```

## 🔒 生产环境建议

1. **修改默认密码**：在 `.env` 文件中修改所有默认密码
2. **使用 HTTPS**：配置 SSL 证书，使用 nginx 反向代理
3. **限制端口暴露**：只暴露必要的端口（80/443）
4. **定期备份**：设置定时任务备份数据库
5. **监控日志**：配置日志收集和监控系统
6. **资源限制**：在 docker-compose.yml 中添加资源限制

## 📊 性能优化

### 添加资源限制

在 `docker-compose.yml` 中为每个服务添加：

```yaml
services:
  backend:
    # ...
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 1G
        reservations:
          cpus: '0.5'
          memory: 512M
```

### 使用生产环境配置

确保后端使用 `config.docker.yaml` 配置文件。

## 🆘 获取帮助

如遇到问题，请：

1. 查看日志：`docker-compose logs -f`
2. 检查服务状态：`docker-compose ps`
3. 查看容器资源使用：`docker stats`
4. 提交 Issue 到项目仓库
