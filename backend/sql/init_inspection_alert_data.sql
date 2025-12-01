-- 创建巡检趋势表
CREATE TABLE IF NOT EXISTS `ops_inspection_trend` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `label` varchar(50) NOT NULL COMMENT '标签（日期或分类）',
  `value` int(11) NOT NULL DEFAULT '0' COMMENT '数值',
  `segment` varchar(50) DEFAULT NULL COMMENT '分段标识',
  `sequence` int(11) NOT NULL DEFAULT '0' COMMENT '排序序号',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_sequence` (`sequence`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='巡检趋势数据';

-- 创建告警指标表
CREATE TABLE IF NOT EXISTS `ops_alert_metrics` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `label` varchar(50) NOT NULL COMMENT '标签（级别或分类）',
  `value` int(11) NOT NULL DEFAULT '0' COMMENT '数值',
  `segment` varchar(50) DEFAULT NULL COMMENT '分段标识',
  `severity` int(11) NOT NULL DEFAULT '0' COMMENT '严重程度（用于排序）',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_severity` (`severity`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警指标数据';

-- 插入巡检趋势模拟数据（最近7天）
INSERT INTO `ops_inspection_trend` (`label`, `value`, `segment`, `sequence`) VALUES
('Mon', 12, '自动', 1),
('Tue', 18, '自动', 2),
('Wed', 22, '自动', 3),
('Thu', 25, '自动', 4),
('Fri', 21, '自动', 5),
('Sat', 30, '自动', 6),
('Sun', 35, '自动', 7);

-- 插入告警指标模拟数据（按级别）
INSERT INTO `ops_alert_metrics` (`label`, `value`, `segment`, `severity`) VALUES
('P1', 3, 'P1', 1),
('P2', 11, 'P2', 2),
('P3', 20, 'P3', 3),
('P4', 8, 'P4', 4);
