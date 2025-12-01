# 导入旧版项目规划数据到V2

## 执行步骤

### 方式1：使用mysql命令行

```bash
mysql -h 9.134.95.57 -P 3306 -u xedv -p xedv < robot-manage/backend/migrations/004_insert_initial_data_v2.sql
```

### 方式2：使用数据库管理工具

1. 打开Navicat、DBeaver或其他数据库管理工具
2. 连接到数据库：
   - 主机：9.134.95.57
   - 端口：3306
   - 用户：xedv
   - 数据库：xedv
3. 打开SQL文件：`robot-manage/backend/migrations/004_insert_initial_data_v2.sql`
4. 执行SQL脚本

## 数据说明

脚本会插入以下数据：

### 模块（6个）
1. Dashboard模块 📊
2. 信息管理系统模块 🗄️
3. 智能交付模块 🚀
4. 数字员工模块 🤖
5. 系统管理模块 ⚙️
6. 数字人形象 👤

### 任务（共46个）
- Dashboard模块：12个任务
- 信息管理系统模块：9个任务
- 智能交付模块：8个任务
- 数字员工模块：2个任务
- 系统管理模块：7个任务
- 数字人形象：2个任务

### 工时转换规则
- "3天" → 24小时（3 × 8）
- "1天" → 8小时
- 所有工时已转换为小时数

## 验证

执行完成后，访问：http://localhost:8088/help/project-plan-v2

应该能看到所有模块和任务数据。
