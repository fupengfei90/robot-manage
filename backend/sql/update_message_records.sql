-- 数字员工会话记录数据更新脚本
-- 执行前请先备份数据库！

-- 1. 添加新字段（如果不存在）
ALTER TABLE wecom_message_records 
ADD COLUMN `type` VARCHAR(50) DEFAULT NULL COMMENT '会话类型：question, command, chat, scheduled_task' AFTER `debug`;

ALTER TABLE wecom_message_records 
ADD COLUMN `is_valid` INT DEFAULT 1 COMMENT '是否有效：1有效，0无效' AFTER `type`;

ALTER TABLE wecom_message_records 
ADD INDEX `idx_type` (`type`);

ALTER TABLE wecom_message_records 
ADD INDEX `idx_is_valid` (`is_valid`);

-- 2. 更新type字段：标记定时任务
UPDATE wecom_message_records 
SET `type` = 'scheduled_task' 
WHERE request_user_id = 'jiaofu';

-- 3. 更新is_valid字段：标记无效数据
UPDATE wecom_message_records 
SET is_valid = 0 
WHERE input LIKE '定时任务触发:%';

-- 4. 更新is_valid字段：特定的无效数据
UPDATE wecom_message_records 
SET is_valid = 0 
WHERE input = '定时任务触发: ITSM提单记录同步';

-- 5. 统计更新结果
SELECT 
    '总记录数' AS category,
    COUNT(*) AS count
FROM wecom_message_records
UNION ALL
SELECT 
    '定时任务' AS category,
    COUNT(*) AS count
FROM wecom_message_records
WHERE type = 'scheduled_task'
UNION ALL
SELECT 
    '无效数据' AS category,
    COUNT(*) AS count
FROM wecom_message_records
WHERE is_valid = 0
UNION ALL
SELECT 
    '有效数据' AS category,
    COUNT(*) AS count
FROM wecom_message_records
WHERE is_valid = 1;

-- 6. 查看类型分布
SELECT 
    COALESCE(type, '未分类') AS type,
    COUNT(*) AS count,
    ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM wecom_message_records), 2) AS percentage
FROM wecom_message_records
GROUP BY type
ORDER BY count DESC;

-- 7. 查看有效性分布
SELECT 
    CASE 
        WHEN is_valid = 1 THEN '有效'
        WHEN is_valid = 0 THEN '无效'
        ELSE '未知'
    END AS validity,
    COUNT(*) AS count
FROM wecom_message_records
GROUP BY is_valid;
