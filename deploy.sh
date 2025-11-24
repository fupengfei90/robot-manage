#!/bin/bash

# 容器化部署脚本

set -e

echo "🚀 开始部署 Robot Manage 系统..."

# 检查 Docker 是否安装
if ! command -v docker &> /dev/null; then
    echo "❌ Docker 未安装，请先安装 Docker"
    exit 1
fi

# 检查 Docker Compose 是否安装
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose 未安装，请先安装 Docker Compose"
    exit 1
fi

# 创建 .env 文件（如果不存在）
if [ ! -f .env ]; then
    echo "📝 创建 .env 文件..."
    cp .env.example .env
    echo "⚠️  请编辑 .env 文件配置生产环境参数"
fi

# 停止并删除旧容器
echo "🛑 停止旧容器..."
docker-compose down

# 构建镜像
echo "🔨 构建 Docker 镜像..."
docker-compose build --no-cache

# 启动服务
echo "▶️  启动服务..."
docker-compose up -d

# 等待服务启动
echo "⏳ 等待服务启动..."
sleep 10

# 检查服务状态
echo "✅ 检查服务状态..."
docker-compose ps

echo ""
echo "🎉 部署完成！"
echo "📱 前端访问地址: http://localhost"
echo "🔧 后端API地址: http://localhost:8088"
echo ""
echo "📊 查看日志: docker-compose logs -f"
echo "🛑 停止服务: docker-compose down"
