#!/bin/bash

# 数字员工会话记录数据更新脚本
# 使用方法: ./update_message_records.sh

set -e

# 数据库配置
DB_HOST="9.134.95.57"
DB_PORT="3306"
DB_USER="xedv"
DB_PASS="Xedv@2024"
DB_NAME="xedv"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}========================================${NC}"
echo -e "${YELLOW}数字员工会话记录数据更新脚本${NC}"
echo -e "${YELLOW}========================================${NC}"
echo ""

# 检查mysql命令
if ! command -v mysql &> /dev/null; then
    echo -e "${RED}错误: 未找到mysql命令，请先安装MySQL客户端${NC}"
    exit 1
fi

# 确认执行
echo -e "${YELLOW}警告: 此脚本将修改数据库数据！${NC}"
echo -e "数据库: ${DB_HOST}:${DB_PORT}/${DB_NAME}"
echo ""
read -p "是否继续？(yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    echo -e "${RED}操作已取消${NC}"
    exit 0
fi

echo ""
echo -e "${GREEN}开始执行数据更新...${NC}"
echo ""

# 执行SQL脚本
mysql -h"${DB_HOST}" -P"${DB_PORT}" -u"${DB_USER}" -p"${DB_PASS}" "${DB_NAME}" < update_message_records.sql

if [ $? -eq 0 ]; then
    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${GREEN}数据更新完成！${NC}"
    echo -e "${GREEN}========================================${NC}"
else
    echo ""
    echo -e "${RED}========================================${NC}"
    echo -e "${RED}数据更新失败！${NC}"
    echo -e "${RED}========================================${NC}"
    exit 1
fi
