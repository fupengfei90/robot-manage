-- 创建调度任务表 (基于wecom_tasks表结构)
CREATE TABLE IF NOT EXISTS wecom_tasks (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) COLLATE utf8mb4_bin NOT NULL COMMENT '任务名称',
    active TINYINT(4) DEFAULT '0' COMMENT '是否可用',
    cron VARCHAR(100) COLLATE utf8mb4_bin NOT NULL COMMENT '定时任务表达式',
    workflow VARCHAR(100) COLLATE utf8mb4_bin NOT NULL COMMENT '工作流ID',
    exec_token VARCHAR(100) COLLATE utf8mb4_bin NOT NULL COMMENT '执行token',
    created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
    updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
    deleted_at DATETIME(3) DEFAULT NULL,
    cron_entry_id BIGINT(20) NOT NULL DEFAULT '0',
    category VARCHAR(32) COLLATE utf8mb4_bin NOT NULL COMMENT '任务分类',
    dcn VARCHAR(100) COLLATE utf8mb4_bin NOT NULL COMMENT '数据中心',
    PRIMARY KEY (id),
    KEY idx_wecom_tasks_deleted_at (deleted_at),
    KEY idx_wecom_tasks_active (active),
    KEY idx_wecom_tasks_category (category),
    KEY idx_wecom_tasks_dcn (dcn),
    UNIQUE KEY uk_wecom_tasks_name_dcn (name, dcn, deleted_at)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- 插入示例数据
INSERT INTO wecom_tasks (id, name, active, cron, workflow, exec_token, created_at, updated_at, deleted_at, cron_entry_id, category, dcn) VALUES
(1, '值班人提醒', 1, '0 0 10 * * *', 'bd35a471-a068-4cb1-8215-2c0003ba0382', 'app-WGfAztV6fmzz60cZZ6p9aG74', '2024-01-01 00:00:00.000', '2025-11-19 17:41:01.847', NULL, 1, 'rota', '6A0'),
(2, '每周周报', 1, '0 0 0 * * 5', 'bd35a471-a068-4cb1-8215-2c0003ba0382', 'app-WGfAztV6fmzz60cZZ6p9aG74', '2024-01-01 00:00:00.000', '2025-11-19 17:41:01.849', NULL, 2, 'weekly-report', '6A0'),
(3, '版本拉齐通知', 1, '0 30 10 * * 1,2,3,4,5', 'bd35a471-a068-4cb1-8215-2c0003ba0382', 'app-WGfAztV6fmzz60cZZ6p9aG74', '2024-01-01 00:00:00.000', '2025-11-19 17:41:01.851', NULL, 3, 'version-align', '6A0');