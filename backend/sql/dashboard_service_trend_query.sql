-- Dashboard 服务次数趋势统计SQL
-- 统计最近7天的服务次数，包括总数、用户服务、定时任务

-- 1. 服务次数趋势（按天统计）
SELECT 
    DATE(created_at) as date,
    -- 用户服务次数（排除定时任务，只统计有效数据）
    SUM(CASE 
        WHEN (type IS NULL OR type != 'scheduled_task') 
        AND (is_valid = 1 OR is_valid IS NULL) 
        THEN 1 
        ELSE 0 
    END) as user_service_count,
    -- 定时任务次数（只统计有效数据）
    SUM(CASE 
        WHEN type = 'scheduled_task' 
        AND (is_valid = 1 OR is_valid IS NULL) 
        THEN 1 
        ELSE 0 
    END) as scheduled_task_count,
    -- 总服务次数 = 用户服务 + 定时任务
    SUM(CASE 
        WHEN (type IS NULL OR type != 'scheduled_task') 
        AND (is_valid = 1 OR is_valid IS NULL) 
        THEN 1 
        ELSE 0 
    END) + 
    SUM(CASE 
        WHEN type = 'scheduled_task' 
        AND (is_valid = 1 OR is_valid IS NULL) 
        THEN 1 
        ELSE 0 
    END) as total_service_count
FROM wecom_message_records
WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
GROUP BY DATE(created_at)
ORDER BY date ASC;

-- 2. 汇总统计（最近7天总计）
SELECT 
    '用户服务' as category,
    SUM(CASE 
        WHEN (type IS NULL OR type != 'scheduled_task') 
        AND (is_valid = 1 OR is_valid IS NULL) 
        THEN 1 
        ELSE 0 
    END) as count
FROM wecom_message_records
WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
UNION ALL
SELECT 
    '定时任务' as category,
    SUM(CASE 
        WHEN type = 'scheduled_task' 
        AND (is_valid = 1 OR is_valid IS NULL) 
        THEN 1 
        ELSE 0 
    END) as count
FROM wecom_message_records
WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
UNION ALL
SELECT 
    '总服务次数' as category,
    COUNT(*) as count
FROM wecom_message_records
WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
AND (is_valid = 1 OR is_valid IS NULL);

-- 3. 验证数据一致性（总数应该等于用户服务+定时任务）
SELECT 
    '验证' as check_type,
    (SELECT COUNT(*) 
     FROM wecom_message_records 
     WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
     AND (is_valid = 1 OR is_valid IS NULL)) as total,
    (SELECT SUM(CASE 
            WHEN (type IS NULL OR type != 'scheduled_task') 
            AND (is_valid = 1 OR is_valid IS NULL) 
            THEN 1 ELSE 0 END)
     FROM wecom_message_records 
     WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)) as user_service,
    (SELECT SUM(CASE 
            WHEN type = 'scheduled_task' 
            AND (is_valid = 1 OR is_valid IS NULL) 
            THEN 1 ELSE 0 END)
     FROM wecom_message_records 
     WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)) as scheduled_task,
    -- 验证：total 应该等于 user_service + scheduled_task
    (SELECT COUNT(*) 
     FROM wecom_message_records 
     WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
     AND (is_valid = 1 OR is_valid IS NULL)) = 
    (SELECT SUM(CASE 
            WHEN (type IS NULL OR type != 'scheduled_task') 
            AND (is_valid = 1 OR is_valid IS NULL) 
            THEN 1 ELSE 0 END) +
            SUM(CASE 
            WHEN type = 'scheduled_task' 
            AND (is_valid = 1 OR is_valid IS NULL) 
            THEN 1 ELSE 0 END)
     FROM wecom_message_records 
     WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)) as is_valid_check;

-- 4. 查看type字段分布（用于调试）
SELECT 
    CASE 
        WHEN type IS NULL THEN 'NULL'
        WHEN type = '' THEN 'EMPTY'
        ELSE type 
    END as type_value,
    COUNT(*) as count,
    ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM wecom_message_records WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)), 2) as percentage
FROM wecom_message_records
WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
GROUP BY type
ORDER BY count DESC;

-- 5. 查看is_valid字段分布（用于调试）
SELECT 
    CASE 
        WHEN is_valid IS NULL THEN 'NULL'
        WHEN is_valid = 1 THEN '有效'
        WHEN is_valid = 0 THEN '无效'
        ELSE CAST(is_valid AS CHAR)
    END as validity,
    COUNT(*) as count
FROM wecom_message_records
WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
GROUP BY is_valid;
