-- CMDB数据完整导入脚本
-- 执行方式: mysql -u root -p your_database < import_all_cmdb.sql

-- 1. 创建表
SOURCE create_cmdb_tables.sql;

-- 2. 导入WB CMDB数据
SOURCE insert_wb_cmdb.sql;

-- 3. 导入VB CMDB数据
SOURCE insert_vb_cmdb.sql;

-- 4. 验证数据
SELECT 'WB CMDB记录数:' as info, COUNT(*) as count FROM wb_cmdb_info
UNION ALL
SELECT 'VB CMDB记录数:' as info, COUNT(*) as count FROM vb_cmdb_info;
