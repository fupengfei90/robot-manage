-- Dashboard服务次数统计SQL
-- 这是后端代码中实际使用的查询语句

-- 日统计（今天）
SELECT COUNT(*) 
FROM wecom_message_records 
WHERE msg_id NOT LIKE 'test-%'
  AND DATE(created_at) = CURDATE() 
  AND (is_valid = 1 OR is_valid IS NULL);

-- 周统计（最近7天）
SELECT COUNT(*) 
FROM wecom_message_records 
WHERE msg_id NOT LIKE 'test-%'
  AND DATE(created_at) >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
  AND (is_valid = 1 OR is_valid IS NULL);

-- 月统计（最近30天）
SELECT COUNT(*) 
FROM wecom_message_records 
WHERE msg_id NOT LIKE 'test-%'
  AND DATE(created_at) >= DATE_SUB(CURDATE(), INTERVAL 1 MONTH)
  AND (is_valid = 1 OR is_valid IS NULL);

-- 年统计（最近365天）
SELECT COUNT(*) 
FROM wecom_message_records 
WHERE msg_id NOT LIKE 'test-%'
  AND DATE(created_at) >= DATE_SUB(CURDATE(), INTERVAL 1 YEAR)
  AND (is_valid = 1 OR is_valid IS NULL);

-- 验证查询：查看今天的数据
SELECT 
    DATE(created_at) as date,
    is_valid,
    type,
    COUNT(*) as count
FROM wecom_message_records 
WHERE msg_id NOT LIKE 'test-%'
  AND DATE(created_at) = CURDATE()
GROUP BY DATE(created_at), is_valid, type;

-- 验证查询：查看最近7天的数据分布
SELECT 
    DATE(created_at) as date,
    COUNT(*) as total,
    SUM(CASE WHEN is_valid = 1 OR is_valid IS NULL THEN 1 ELSE 0 END) as valid_count,
    SUM(CASE WHEN is_valid = 0 THEN 1 ELSE 0 END) as invalid_count
FROM wecom_message_records 
WHERE msg_id NOT LIKE 'test-%'
  AND DATE(created_at) >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)
GROUP BY DATE(created_at)
ORDER BY date DESC;
