#!/bin/bash

# 数据库备份脚本

# 配置信息
DB_HOST="9.134.95.57"
DB_PORT="3306"
DB_USER="xedv"
DB_PASSWORD="Xedv@2024"
DB_NAME="xedv"
BACKUP_DIR="./backups"
DATE=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="${BACKUP_DIR}/xedv_backup_${DATE}.sql"

# 创建备份目录
mkdir -p $BACKUP_DIR

echo "🗄️  开始备份数据库..."
echo "数据库: $DB_NAME"
echo "备份文件: $BACKUP_FILE"

# 执行备份
/opt/homebrew/opt/mysql-client/bin/mysqldump -h $DB_HOST -P $DB_PORT -u $DB_USER -p$DB_PASSWORD \
  --skip-lock-tables \
  --skip-add-locks \
  --no-tablespaces \
  --hex-blob \
  --default-character-set=utf8mb4 \
  $DB_NAME > $BACKUP_FILE

if [ $? -eq 0 ]; then
    echo "✅ 数据库备份成功!"
    echo "备份文件大小: $(du -h $BACKUP_FILE | cut -f1)"
    
    # 压缩备份文件
    gzip $BACKUP_FILE
    echo "✅ 备份文件已压缩: ${BACKUP_FILE}.gz"
    
    # 清理7天前的备份
    find $BACKUP_DIR -name "*.sql.gz" -mtime +7 -delete
    echo "🧹 已清理7天前的备份文件"
    
else
    echo "❌ 数据库备份失败!"
    exit 1
fi