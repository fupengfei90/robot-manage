#!/bin/bash

# 构建脚本
set -e

echo "🚀 开始构建 Robot Manage 项目..."

# 1. 构建前端
echo "📦 构建前端..."
cd frontend
npm install
npm run build
cd ..

# 2. 构建后端
echo "🔨 构建后端..."
cd backend
go mod download
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o robot-manage ./cmd/main.go
cd ..

# 3. 创建部署目录
echo "📁 创建部署包..."
DEPLOY_DIR="deploy_$(date +%Y%m%d_%H%M%S)"
mkdir -p $DEPLOY_DIR

# 4. 复制文件
cp -r frontend/dist $DEPLOY_DIR/dist
cp backend/robot-manage $DEPLOY_DIR/
cp -r backend/configs $DEPLOY_DIR/
cp -r backend/sql $DEPLOY_DIR/

# 5. 创建启动脚本
cat > $DEPLOY_DIR/start.sh << 'EOF'
#!/bin/bash
nohup ./robot-manage > robot-manage.log 2>&1 &
echo $! > robot-manage.pid
echo "✅ 服务已启动，PID: $(cat robot-manage.pid)"
EOF

# 6. 创建停止脚本
cat > $DEPLOY_DIR/stop.sh << 'EOF'
#!/bin/bash
if [ -f robot-manage.pid ]; then
    PID=$(cat robot-manage.pid)
    kill $PID
    rm robot-manage.pid
    echo "✅ 服务已停止"
else
    echo "❌ 未找到PID文件"
fi
EOF

chmod +x $DEPLOY_DIR/start.sh
chmod +x $DEPLOY_DIR/stop.sh
chmod +x $DEPLOY_DIR/robot-manage

# 7. 打包
echo "📦 打包部署文件..."
tar -czf ${DEPLOY_DIR}.tar.gz $DEPLOY_DIR

echo "✅ 构建完成！"
echo "📦 部署包: ${DEPLOY_DIR}.tar.gz"
echo ""
echo "部署步骤："
echo "1. 上传 ${DEPLOY_DIR}.tar.gz 到服务器"
echo "2. 解压: tar -xzf ${DEPLOY_DIR}.tar.gz"
echo "3. 进入目录: cd $DEPLOY_DIR"
echo "4. 修改配置: vim configs/config.yaml"
echo "5. 启动服务: ./start.sh"
