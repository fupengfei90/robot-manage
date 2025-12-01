-- ITSM出包记录表
CREATE TABLE IF NOT EXISTS `itsm_package_records` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `subsystem` VARCHAR(100) NOT NULL COMMENT '子系统',
  `package_name` VARCHAR(255) NOT NULL COMMENT '物料包',
  `requirement_id` VARCHAR(50) NOT NULL COMMENT '需求ID',
  `itsm_ticket` VARCHAR(50) NOT NULL COMMENT 'ITSM单',
  `status` VARCHAR(50) NOT NULL COMMENT '状态',
  `owner` VARCHAR(100) NOT NULL COMMENT '负责人',
  `environment` VARCHAR(50) NOT NULL COMMENT '发布环境',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_subsystem` (`subsystem`),
  INDEX `idx_status` (`status`),
  INDEX `idx_environment` (`environment`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ITSM出包记录表';
