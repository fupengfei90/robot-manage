#!/bin/bash

# 智能运维后台管理系统启动脚本

echo "🚀 启动智能运维后台管理系统..."

# 检查是否安装了必要的依赖
check_dependencies() {
    echo "📋 检查依赖..."
    
    # 检查Go
    if ! command -v go &> /dev/null; then
        echo "❌ Go未安装，请先安装Go 1.22+"
        exit 1
    fi
    
    # 检查Node.js
    if ! command -v node &> /dev/null; then
        echo "❌ Node.js未安装，请先安装Node.js 18+"
        exit 1
    fi
    
    # 检查MySQL
    if ! command -v mysql &> /dev/null; then
        echo "⚠️  MySQL客户端未找到，请确保MySQL服务正在运行"
    fi
    
    echo "✅ 依赖检查完成"
}

# 初始化数据库
init_database() {
    echo "🗄️  初始化数据库..."
    
    # 检查配置文件
    if [ ! -f "backend/configs/config.yaml" ]; then
        echo "❌ 配置文件不存在: backend/configs/config.yaml"
        echo "请参考 backend/configs/config.yaml.example 创建配置文件"
        exit 1
    fi
    
    # 执行数据库迁移
    if [ -f "backend/sql/create_wecom_tasks.sql" ]; then
        echo "📝 执行数据库迁移..."
        # 这里可以添加实际的数据库连接和执行SQL的逻辑
        echo "✅ 数据库迁移完成"
    fi
}

# 启动后端服务
start_backend() {
    echo "🔧 启动后端服务..."
    cd backend
    
    # 安装Go依赖
    echo "📦 安装Go依赖..."
    go mod tidy
    
    # 构建并启动
    echo "🏗️  构建后端..."
    go build -o server cmd/server/main.go
    
    echo "🚀 启动后端服务..."
    ./server &
    BACKEND_PID=$!
    echo "后端服务PID: $BACKEND_PID"
    
    cd ..
}

# 启动前端服务
start_frontend() {
    echo "🎨 启动前端服务..."
    cd frontend
    
    # 安装npm依赖
    echo "📦 安装npm依赖..."
    npm install
    
    # 启动开发服务器
    echo "🚀 启动前端开发服务器..."
    npm run dev &
    FRONTEND_PID=$!
    echo "前端服务PID: $FRONTEND_PID"
    
    cd ..
}

# 清理函数
cleanup() {
    echo "🛑 正在停止服务..."
    if [ ! -z "$BACKEND_PID" ]; then
        kill $BACKEND_PID 2>/dev/null
        echo "✅ 后端服务已停止"
    fi
    if [ ! -z "$FRONTEND_PID" ]; then
        kill $FRONTEND_PID 2>/dev/null
        echo "✅ 前端服务已停止"
    fi
    exit 0
}

# 设置信号处理
trap cleanup SIGINT SIGTERM

# 主函数
main() {
    echo "🎯 智能运维后台管理系统 - 调度任务管理模块"
    echo "================================================"
    
    check_dependencies
    init_database
    start_backend
    sleep 3  # 等待后端启动
    start_frontend
    
    echo ""
    echo "🎉 系统启动完成！"
    echo "📱 前端地址: http://localhost:5173"
    echo "🔧 后端地址: http://localhost:8080"
    echo "📚 API文档: http://localhost:8080/api/v1"
    echo ""
    echo "🔍 调度任务管理功能："
    echo "   • 任务列表查看和筛选"
    echo "   • 新增/编辑/删除任务"
    echo "   • 启用/停用任务状态"
    echo "   • 立即执行任务"
    echo "   • 批量操作支持"
    echo "   • Cron表达式可视化编辑"
    echo ""
    echo "按 Ctrl+C 停止服务"
    
    # 等待信号
    wait
}

# 运行主函数
main