#!/bin/bash

# 数据库恢复脚本

# 配置信息
DB_HOST="localhost"
DB_PORT="3306"
DB_USER="root"
DB_NAME="robot_manage"
BACKUP_DIR="./backups"

# 检查参数
if [ $# -eq 0 ]; then
    echo "使用方法: $0 <备份文件名>"
    echo "可用的备份文件:"
    ls -la $BACKUP_DIR/*.sql.gz 2>/dev/null || echo "没有找到备份文件"
    exit 1
fi

BACKUP_FILE="$1"

# 检查备份文件是否存在
if [ ! -f "$BACKUP_FILE" ]; then
    # 尝试在备份目录中查找
    if [ -f "$BACKUP_DIR/$BACKUP_FILE" ]; then
        BACKUP_FILE="$BACKUP_DIR/$BACKUP_FILE"
    else
        echo "❌ 备份文件不存在: $BACKUP_FILE"
        exit 1
    fi
fi

echo "🔄 开始恢复数据库..."
echo "数据库: $DB_NAME"
echo "备份文件: $BACKUP_FILE"

# 确认操作
read -p "⚠️  这将覆盖现有数据库，确定要继续吗? (y/N): " confirm
if [[ $confirm != [yY] ]]; then
    echo "操作已取消"
    exit 0
fi

# 解压备份文件（如果是压缩的）
if [[ $BACKUP_FILE == *.gz ]]; then
    echo "📦 解压备份文件..."
    TEMP_FILE="${BACKUP_FILE%.gz}"
    gunzip -c "$BACKUP_FILE" > "$TEMP_FILE"
    RESTORE_FILE="$TEMP_FILE"
else
    RESTORE_FILE="$BACKUP_FILE"
fi

# 执行恢复
mysql -h $DB_HOST -P $DB_PORT -u $DB_USER -p $DB_NAME < "$RESTORE_FILE"

if [ $? -eq 0 ]; then
    echo "✅ 数据库恢复成功!"
    
    # 清理临时文件
    if [[ $BACKUP_FILE == *.gz ]] && [ -f "$TEMP_FILE" ]; then
        rm "$TEMP_FILE"
        echo "🧹 已清理临时文件"
    fi
else
    echo "❌ 数据库恢复失败!"
    exit 1
fi