-- ============================================
-- 更新用户密码脚本
-- 使用方法: 先运行 generate_password.go 生成新的密码哈希值
-- ============================================

-- 更新admin用户密码为 admin123456
-- bcrypt哈希值: $2a$10$0jNu6INi6uUp6yJb0sgBSu.Nan4ah.PRoiTjSYd9eOw.PIqg8xrxq
UPDATE `manage_users` 
SET `password` = '$2a$10$0jNu6INi6uUp6yJb0sgBSu.Nan4ah.PRoiTjSYd9eOw.PIqg8xrxq',
    `updated_at` = CURRENT_TIMESTAMP
WHERE `username` = 'admin';

-- 更新testuser用户密码为 test123456
-- bcrypt哈希值: $2a$10$Op.FucFYsfdhywIDC5162e6o4kjt1yjWv5JnpmRlmbHTp7azum9J6
UPDATE `manage_users` 
SET `password` = '$2a$10$Op.FucFYsfdhywIDC5162e6o4kjt1yjWv5JnpmRlmbHTp7azum9J6',
    `updated_at` = CURRENT_TIMESTAMP
WHERE `username` = 'testuser';

