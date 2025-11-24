-- ============================================
-- Robot Manage 用户表初始化脚本
-- ============================================

-- 创建用户表
CREATE TABLE IF NOT EXISTS `manage_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码（bcrypt加密）',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `role` varchar(20) DEFAULT 'user' COMMENT '角色：admin-管理员, user-普通用户',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像URL',
  `status` tinyint(1) DEFAULT 1 COMMENT '状态：1-启用, 0-禁用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_status` (`status`),
  KEY `idx_role` (`role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 插入默认管理员用户
-- 用户名: admin
-- 密码: admin123456
-- bcrypt哈希值: $2a$10$0jNu6INi6uUp6yJb0sgBSu.Nan4ah.PRoiTjSYd9eOw.PIqg8xrxq
INSERT INTO `manage_users` (`username`, `password`, `email`, `role`, `status`) 
VALUES 
  ('admin', '$2a$10$0jNu6INi6uUp6yJb0sgBSu.Nan4ah.PRoiTjSYd9eOw.PIqg8xrxq', 'admin@robot-manage.com', 'admin', 1)
ON DUPLICATE KEY UPDATE `password` = '$2a$10$0jNu6INi6uUp6yJb0sgBSu.Nan4ah.PRoiTjSYd9eOw.PIqg8xrxq', `updated_at` = CURRENT_TIMESTAMP;

-- 插入测试普通用户（可选）
-- 用户名: testuser
-- 密码: test123456
-- bcrypt哈希值: $2a$10$Op.FucFYsfdhywIDC5162e6o4kjt1yjWv5JnpmRlmbHTp7azum9J6
INSERT INTO `manage_users` (`username`, `password`, `email`, `role`, `status`) 
VALUES 
  ('testuser', '$2a$10$Op.FucFYsfdhywIDC5162e6o4kjt1yjWv5JnpmRlmbHTp7azum9J6', 'test@robot-manage.com', 'user', 1)
ON DUPLICATE KEY UPDATE `password` = '$2a$10$Op.FucFYsfdhywIDC5162e6o4kjt1yjWv5JnpmRlmbHTp7azum9J6', `updated_at` = CURRENT_TIMESTAMP;
