#!/bin/bash

set -e

echo "======================================"
echo "项目规划 V2 部署脚本"
echo "======================================"

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 数据库配置
DB_HOST="9.134.95.57"
DB_PORT="3306"
DB_USER="xedv"
DB_NAME="xedv"

echo -e "${YELLOW}步骤 1/5: 执行数据库迁移${NC}"
echo "请手动执行以下SQL脚本："
echo "文件路径: robot-manage/backend/migrations/002_create_project_plan_v2_tables.sql"
echo ""
echo "可以使用以下命令："
echo "mysql -h ${DB_HOST} -P ${DB_PORT} -u ${DB_USER} -p ${DB_NAME} < robot-manage/backend/migrations/002_create_project_plan_v2_tables.sql"
echo ""
read -p "数据库迁移完成后，按回车继续..."

echo -e "${YELLOW}步骤 2/5: 编译后端${NC}"
cd robot-manage/backend
echo "正在编译后端..."
go build -o robot-manage cmd/server/main.go
if [ $? -eq 0 ]; then
    echo -e "${GREEN}后端编译成功${NC}"
else
    echo -e "${RED}后端编译失败${NC}"
    exit 1
fi
cd ../..

echo -e "${YELLOW}步骤 3/5: 安装前端依赖${NC}"
cd robot-manage/frontend
if [ ! -d "node_modules/vuedraggable" ]; then
    echo "正在安装 vuedraggable..."
    npm install vuedraggable@next
fi

echo -e "${YELLOW}步骤 4/5: 编译前端${NC}"
echo "正在编译前端..."
npm run build
if [ $? -eq 0 ]; then
    echo -e "${GREEN}前端编译成功${NC}"
else
    echo -e "${RED}前端编译失败${NC}"
    exit 1
fi
cd ../..

echo -e "${YELLOW}步骤 5/5: 部署${NC}"
echo "正在复制文件到后端dist目录..."
rm -rf robot-manage/backend/dist
cp -r robot-manage/frontend/dist robot-manage/backend/

echo -e "${GREEN}======================================"
echo "部署完成！"
echo "======================================${NC}"
echo ""
echo "启动服务："
echo "  cd robot-manage/backend"
echo "  ./robot-manage"
echo ""
echo "访问地址："
echo "  旧版项目规划: http://localhost:8088/help/project-plan"
echo "  新版项目规划V2: http://localhost:8088/help/project-plan-v2"
echo ""
