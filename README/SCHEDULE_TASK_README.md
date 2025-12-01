# 调度任务管理模块

## 功能概述

调度任务管理模块是智能运维后台管理系统的核心功能之一，基于现有的`wecom_tasks`表结构开发，提供完整的定时任务管理功能。

## 主要特性

### 🎯 核心功能
- **任务管理**: 完整的CRUD操作（创建、查看、编辑、删除）
- **状态控制**: 支持任务启用/停用切换
- **立即执行**: 手动触发任务执行
- **批量操作**: 支持批量启用、停用、删除
- **智能筛选**: 多条件筛选和搜索

### 🛠️ 高级特性
- **Cron可视化**: 集成可视化cron表达式编辑器
- **实时验证**: cron表达式合法性验证和描述解析
- **下次执行**: 自动计算并显示任务下次执行时间
- **分类管理**: 支持任务分类和数据中心管理
- **响应式设计**: 适配不同屏幕尺寸

### 🎨 界面设计
- **深色科技风**: 采用磨砂玻璃效果和渐变设计
- **高端大气**: 与系统整体风格保持一致
- **交互友好**: 丰富的动效和状态反馈

## 技术架构

### 后端技术栈
- **Go 1.22+**: 主要开发语言
- **Gin**: Web框架
- **GORM**: ORM框架
- **MySQL**: 数据库
- **Cron**: 定时任务解析

### 前端技术栈
- **Vue 3**: 前端框架
- **TypeScript**: 类型安全
- **Element Plus**: UI组件库
- **Vite**: 构建工具

## 数据库设计

### 表结构 (wecom_tasks)
```sql
CREATE TABLE wecom_tasks (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT '任务名称',
    active TINYINT(4) DEFAULT '0' COMMENT '是否可用',
    cron VARCHAR(100) NOT NULL COMMENT '定时任务表达式',
    workflow VARCHAR(100) NOT NULL COMMENT '工作流ID',
    exec_token VARCHAR(100) NOT NULL COMMENT '执行token',
    created_at DATETIME(3) DEFAULT NULL COMMENT '创建时间',
    updated_at DATETIME(3) DEFAULT NULL COMMENT '更新时间',
    deleted_at DATETIME(3) DEFAULT NULL,
    cron_entry_id BIGINT(20) NOT NULL DEFAULT '0',
    category VARCHAR(32) NOT NULL COMMENT '任务分类',
    dcn VARCHAR(100) NOT NULL COMMENT '数据中心',
    PRIMARY KEY (id),
    KEY idx_wecom_tasks_deleted_at (deleted_at)
);
```

### 字段说明
- `name`: 任务名称，用于标识任务
- `active`: 任务状态，1=启用，0=停用
- `cron`: cron表达式，定义任务执行时间
- `workflow`: 工作流ID，关联具体的执行逻辑
- `exec_token`: 执行令牌，用于权限验证
- `category`: 任务分类（rota、weekly-report、version-align等）
- `dcn`: 数据中心标识

## API接口

### RESTful API设计
```
GET    /api/v1/system/schedule-tasks          # 获取任务列表
GET    /api/v1/system/schedule-tasks/:id      # 获取任务详情
POST   /api/v1/system/schedule-tasks          # 创建任务
PUT    /api/v1/system/schedule-tasks/:id      # 更新任务
DELETE /api/v1/system/schedule-tasks/:id      # 删除任务
PUT    /api/v1/system/schedule-tasks/:id/status # 更新状态
POST   /api/v1/system/schedule-tasks/:id/execute # 立即执行
POST   /api/v1/system/schedule-tasks/batch    # 批量操作
GET    /api/v1/system/schedule-tasks/categories # 获取分类
GET    /api/v1/system/schedule-tasks/dcns      # 获取数据中心
```

### 请求示例
```bash
# 创建任务
curl -X POST http://localhost:8080/api/v1/system/schedule-tasks \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试任务",
    "active": 1,
    "cron": "0 0 10 * * *",
    "workflow": "test-workflow-id",
    "exec_token": "test-token",
    "category": "test",
    "dcn": "6A0"
  }'
```

## 项目结构

```
robot-manage/
├── backend/                          # 后端代码
│   ├── internal/
│   │   ├── model/
│   │   │   └── schedule_task.go      # 数据模型
│   │   ├── repository/
│   │   │   └── schedule_task_repository.go # 数据访问层
│   │   ├── service/
│   │   │   └── schedule_task_service.go    # 业务逻辑层
│   │   └── handler/
│   │       └── schedule_task_handler.go    # 控制器层
│   └── sql/
│       └── create_wecom_tasks.sql    # 数据库迁移
├── frontend/                         # 前端代码
│   ├── src/
│   │   ├── api/
│   │   │   └── schedule-task.ts      # API接口
│   │   ├── components/
│   │   │   ├── CronEditor.vue        # Cron编辑器
│   │   │   ├── TaskForm.vue          # 任务表单
│   │   │   └── TaskFilter.vue        # 筛选组件
│   │   └── views/
│   │       └── system/
│   │           └── ScheduleTaskView.vue # 主页面
└── start.sh                         # 启动脚本
```

## 快速开始

### 1. 环境要求
- Go 1.22+
- Node.js 18+
- MySQL 8.0+
- Redis (可选)

### 2. 安装依赖
```bash
# 后端依赖
cd backend
go mod tidy

# 前端依赖
cd frontend
npm install
```

### 3. 配置数据库
```bash
# 执行数据库迁移
mysql -u root -p < backend/sql/create_wecom_tasks.sql
```

### 4. 启动服务
```bash
# 一键启动（推荐）
./start.sh

# 或者分别启动
# 后端
cd backend && go run cmd/server/main.go

# 前端
cd frontend && npm run dev
```

### 5. 访问系统
- 前端地址: http://localhost:5173
- 后端地址: http://localhost:8080
- 调度任务管理: http://localhost:5173/system/schedule-tasks

## 使用指南

### 创建任务
1. 点击"新增任务"按钮
2. 填写任务基本信息
3. 使用可视化编辑器设置cron表达式
4. 配置工作流ID和执行token
5. 选择任务分类和数据中心
6. 点击"创建"完成

### Cron表达式
支持标准的6位cron表达式格式：`秒 分 时 日 月 周`

示例：
- `0 0 10 * * *` - 每天10点执行
- `0 0 0 * * 5` - 每周五0点执行
- `0 30 10 * * 1,2,3,4,5` - 工作日10:30执行

### 任务分类
- `rota`: 值班提醒
- `weekly-report`: 周报生成
- `version-align`: 版本通知
- 支持自定义分类

## 业务场景

### 1. 值班提醒 (rota)
- **用途**: 定时发送值班人提醒通知
- **执行时间**: 每天10:00
- **cron表达式**: `0 0 10 * * *`

### 2. 周报生成 (weekly-report)
- **用途**: 每周生成并发送周报
- **执行时间**: 每周五00:00
- **cron表达式**: `0 0 0 * * 5`

### 3. 版本通知 (version-align)
- **用途**: 工作日发送版本拉齐通知
- **执行时间**: 工作日10:30
- **cron表达式**: `0 30 10 * * 1,2,3,4,5`

## 开发指南

### 添加新的任务分类
1. 在前端更新分类映射
2. 在后端添加预设分类
3. 更新数据库约束（如需要）

### 扩展功能
- 任务执行历史记录
- 任务依赖关系
- 任务执行结果通知
- 任务性能监控

## 故障排除

### 常见问题
1. **cron表达式验证失败**: 检查表达式格式是否正确
2. **任务创建失败**: 检查任务名称是否重复
3. **任务执行失败**: 检查工作流ID和执行token是否有效

### 日志查看
```bash
# 后端日志
tail -f backend/logs/app.log

# 前端控制台
打开浏览器开发者工具查看Console
```

## 贡献指南

1. Fork项目
2. 创建功能分支
3. 提交代码
4. 创建Pull Request

## 许可证

MIT License

## 联系方式

如有问题或建议，请联系开发团队。