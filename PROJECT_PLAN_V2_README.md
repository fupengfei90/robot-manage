# 项目规划 V2 使用指南

## 功能特性

### 1. 层级化管理
- 支持模块和子模块的多层级结构
- 每个模块可以包含多个子模块和任务项
- 清晰的树形结构展示

### 2. 拖拽调整
- 支持拖拽调整模块顺序（开发中）
- 可视化的层级关系调整

### 3. 丰富的字段设计

#### 模块字段
- 名称：模块标识名称
- 图标：emoji图标
- 父模块：支持层级关系

#### 任务字段
- 名称：任务标识名称（必填）
- 状态：未开始/进行中/已完成/阻塞/已取消
- 优先级：高/中/低/紧急
- 负责人：任务负责人
- 计划工时：计划投入的工时（小时）
- 实际耗时：实际消耗的工时（小时）
- 详细描述：任务的详细说明

### 4. 进度可视化
- 总体进度百分比
- 已完成/总任务数统计
- 计划工时/实际工时统计
- 模块级别的进度展示

### 5. 快速操作
- 一键添加模块/子模块
- 快速添加任务
- 展开/收起全部模块
- 快速编辑和删除

## 部署步骤

### 1. 数据库迁移

连接到数据库并执行迁移脚本：

```bash
# 方式1：使用mysql命令行
mysql -h 9.134.95.57 -u xedv -p xedv < backend/migrations/002_create_project_plan_v2_tables.sql

# 方式2：使用数据库管理工具（如Navicat、DBeaver等）
# 打开 backend/migrations/002_create_project_plan_v2_tables.sql 文件
# 在数据库管理工具中执行SQL脚本
```

### 2. 重新编译后端

```bash
cd backend
go build -o robot-manage cmd/server/main.go
```

### 3. 重新编译前端

```bash
cd frontend
npm install
npm run build
```

### 4. 重启服务

```bash
# 停止现有服务
pkill robot-manage

# 启动新服务
cd backend
./robot-manage
```

## 访问路径

- 旧版项目规划：`/help/project-plan`
- 新版项目规划V2：`/help/project-plan-v2`

## 使用说明

### 添加模块
1. 点击"添加模块"按钮
2. 输入模块名称和图标
3. 点击保存

### 添加子模块
1. 在模块卡片上点击"子模块"按钮
2. 输入子模块信息
3. 点击保存

### 添加任务
1. 在模块卡片上点击"任务"按钮
2. 填写任务详细信息
3. 点击保存

### 编辑和删除
- 点击编辑图标可以修改模块或任务
- 点击删除图标可以删除模块或任务（删除模块会同时删除所有子模块和任务）

### 快速更新状态
- 在任务列表中直接选择状态下拉框即可快速更新任务状态

## API接口

### 模块管理
- `GET /api/v1/project-plan-v2/tree` - 获取模块树
- `POST /api/v1/project-plan-v2/modules` - 创建模块
- `PUT /api/v1/project-plan-v2/modules/:id` - 更新模块
- `DELETE /api/v1/project-plan-v2/modules/:id` - 删除模块
- `POST /api/v1/project-plan-v2/modules/batch` - 批量更新模块

### 任务管理
- `POST /api/v1/project-plan-v2/items` - 创建任务
- `PUT /api/v1/project-plan-v2/items/:id` - 更新任务
- `DELETE /api/v1/project-plan-v2/items/:id` - 删除任务
- `POST /api/v1/project-plan-v2/items/batch` - 批量更新任务

## 数据库表结构

### project_modules_v2
- id: 主键
- parent_id: 父模块ID（NULL表示顶级模块）
- name: 模块名称
- icon: 图标
- sort_order: 排序
- created_at: 创建时间
- updated_at: 更新时间

### project_items_v2
- id: 主键
- module_id: 所属模块ID
- name: 任务名称
- status: 状态
- priority: 优先级
- owner: 负责人
- planned_hours: 计划工时（小时）
- actual_hours: 实际工时（小时）
- description: 详细描述
- sort_order: 排序
- created_at: 创建时间
- updated_at: 更新时间

## 后续优化建议

1. 完善拖拽排序功能
2. 添加甘特图视图
3. 支持任务依赖关系
4. 添加任务评论功能
5. 支持文件附件
6. 添加任务提醒功能
7. 支持导出Excel报表
8. 添加任务时间线视图
