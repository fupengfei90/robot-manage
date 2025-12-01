# CMDB数据导入说明

## 快速导入

### 方式1：使用合并脚本（推荐）
```bash
cd backend/sql
mysql -u root -p your_database < import_all_cmdb.sql
```

### 方式2：分步执行
```bash
cd backend/sql

# 1. 创建表
mysql -u root -p your_database < create_cmdb_tables.sql

# 2. 导入WB CMDB数据
mysql -u root -p your_database < insert_wb_cmdb.sql

# 3. 导入VB CMDB数据
mysql -u root -p your_database < insert_vb_cmdb.sql
```

## 验证数据

```sql
-- 查看记录数
SELECT 'WB CMDB' as table_name, COUNT(*) as count FROM wb_cmdb_info
UNION ALL
SELECT 'VB CMDB' as table_name, COUNT(*) as count FROM vb_cmdb_info;

-- 查看示例数据
SELECT * FROM wb_cmdb_info LIMIT 5;
SELECT * FROM vb_cmdb_info LIMIT 5;
```

## 重新导入

如需重新导入，先清空数据：

```sql
TRUNCATE TABLE wb_cmdb_info;
TRUNCATE TABLE vb_cmdb_info;
```

然后重新执行导入脚本。

## 文件说明

- `create_cmdb_tables.sql` - 建表SQL
- `insert_wb_cmdb.sql` - WB CMDB数据（约500条）
- `insert_vb_cmdb.sql` - VB CMDB数据（约7900条）
- `import_all_cmdb.sql` - 完整导入脚本
