# 数字员工模块实现指南

## 已完成的后端代码

### 1. 模型层 (Model)
✅ `/backend/internal/model/digital_employee.go`
- MessageRecord - 历史会话记录模型
- ExportRecord - 服务回传记录模型
- 查询参数、请求体、统计数据等相关模型

### 2. 数据访问层 (Repository)
✅ `/backend/internal/repository/digital_employee_repository.go`
- 会话记录的CRUD操作
- 导出记录的CRUD操作
- 统计数据查询
- 批量操作支持

### 3. 业务逻辑层 (Service)
✅ `/backend/internal/service/digital_employee_service.go`
- 会话记录管理业务逻辑
- 导出记录管理业务逻辑
- 数据统计和分析

### 4. 控制器层 (Handler)
✅ `/backend/internal/handler/digital_employee_handler.go`
- RESTful API端点实现
- 请求参数验证
- 响应格式化

### 5. 前端API
✅ `/frontend/src/api/digital-employee.ts`
- TypeScript类型定义
- API接口封装

## 需要完成的集成工作

### 1. 后端路由注册

在 `/backend/internal/server/router.go` 中添加路由：

```go
// 在 setupRoutes 函数中添加
func (s *Server) setupRoutes() {
    // ... 现有代码 ...
    
    // 数字员工模块
    digitalEmployee := v1.Group("/digital-employee")
    {
        // 历史会话记录
        digitalEmployee.GET("/message-records", s.digitalEmployeeHandler.GetMessageRecords)
        digitalEmployee.GET("/message-records/:id", s.digitalEmployeeHandler.GetMessageRecordByID)
        digitalEmployee.POST("/message-records", s.digitalEmployeeHandler.CreateMessageRecord)
        digitalEmployee.PUT("/message-records/:id", s.digitalEmployeeHandler.UpdateMessageRecord)
        digitalEmployee.DELETE("/message-records/:id", s.digitalEmployeeHandler.DeleteMessageRecord)
        digitalEmployee.POST("/message-records/batch", s.digitalEmployeeHandler.BatchOperateMessageRecords)
        digitalEmployee.GET("/conversations/:conversation_id/messages", s.digitalEmployeeHandler.GetMessagesByConversationID)
        digitalEmployee.GET("/conversation-groups", s.digitalEmployeeHandler.GetConversationGroups)
        digitalEmployee.GET("/message-statistics", s.digitalEmployeeHandler.GetMessageStatistics)
        
        // 服务回传记录
        digitalEmployee.GET("/export-records", s.digitalEmployeeHandler.GetExportRecords)
        digitalEmployee.GET("/export-records/:id", s.digitalEmployeeHandler.GetExportRecordByID)
        digitalEmployee.POST("/export-records", s.digitalEmployeeHandler.CreateExportRecord)
        digitalEmployee.PUT("/export-records/:id", s.digitalEmployeeHandler.UpdateExportRecord)
        digitalEmployee.DELETE("/export-records/:id", s.digitalEmployeeHandler.DeleteExportRecord)
        digitalEmployee.POST("/export-records/batch", s.digitalEmployeeHandler.BatchOperateExportRecords)
        digitalEmployee.POST("/export-records/:id/retry", s.digitalEmployeeHandler.RetryExportEmail)
        digitalEmployee.GET("/export-statistics", s.digitalEmployeeHandler.GetExportStatistics)
    }
}
```

### 2. 依赖注入

在 `/backend/internal/bootstrap/wire.go` 中添加：

```go
// 添加到 InitializeServer 函数
func InitializeServer() (*server.Server, error) {
    // ... 现有代码 ...
    
    // 数字员工模块
    digitalEmployeeRepo := repository.NewDigitalEmployeeRepository(db)
    digitalEmployeeService := service.NewDigitalEmployeeService(digitalEmployeeRepo)
    digitalEmployeeHandler := handler.NewDigitalEmployeeHandler(digitalEmployeeService)
    
    // 传递给 Server
    srv := server.NewServer(
        cfg,
        // ... 其他handler ...
        digitalEmployeeHandler,
    )
    
    return srv, nil
}
```

### 3. 前端路由配置

在 `/frontend/src/router/index.ts` 中添加：

```typescript
{
  path: '/digital-employee',
  name: 'digitalEmployee',
  redirect: '/digital-employee/message-records',
  children: [
    {
      path: 'message-records',
      name: 'messageRecords',
      component: () => import('@/views/digital-employee/MessageRecordsView.vue')
    },
    {
      path: 'export-records',
      name: 'exportRecords',
      component: () => import('@/views/digital-employee/ExportRecordsView.vue')
    }
  ]
}
```

### 4. 侧边栏菜单

在 `/frontend/src/layouts/MainLayout.vue` 中添加菜单项：

```vue
<el-sub-menu index="digitalEmployee" class="menu-item">
  <template #title>
    <span class="menu-item__icon">🤖</span>
    <span>数字员工</span>
  </template>
  <el-menu-item index="messageRecords" @click="$router.push('/digital-employee/message-records')">
    历史会话记录
  </el-menu-item>
  <el-menu-item index="exportRecords" @click="$router.push('/digital-employee/export-records')">
    服务回传记录
  </el-menu-item>
</el-sub-menu>
```

## 前端页面实现

### 页面文件结构
```
frontend/src/views/digital-employee/
├── MessageRecordsView.vue      # 历史会话记录主页面
├── ExportRecordsView.vue        # 服务回传记录主页面
└── components/
    ├── MessageDetailDialog.vue  # 会话详情对话框
    ├── ConversationFlow.vue     # 对话流展示
    ├── ExportDetailDialog.vue   # 导出详情对话框
    └── StatisticsPanel.vue      # 统计面板组件
```

### 核心功能实现要点

#### 1. 历史会话记录页面
- 使用 `el-table` 展示数据
- 实现高级筛选（时间范围、用户、会话ID等）
- 支持全文搜索（input/output字段）
- 批量操作（删除、标记调试）
- 会话详情弹窗（完整内容展示）
- 对话流可视化（按会话分组）

#### 2. 服务回传记录页面
- 状态标签（成功/失败/待发送）
- 邮件信息展示
- 重试发送功能
- 文件下载链接
- 批量操作支持

#### 3. 数据可视化
- 使用 ECharts 展示趋势图
- 用户活跃度统计
- 导出成功率饼图
- 实时数据刷新

### 样式要求
- 使用现有的 CSS 变量系统
- 保持 glass-effect 磨砂玻璃效果
- 渐变色彩和动画效果
- 响应式布局

## API 端点列表

### 历史会话记录
- `GET /api/v1/digital-employee/message-records` - 获取列表
- `GET /api/v1/digital-employee/message-records/:id` - 获取详情
- `POST /api/v1/digital-employee/message-records` - 创建记录
- `PUT /api/v1/digital-employee/message-records/:id` - 更新记录
- `DELETE /api/v1/digital-employee/message-records/:id` - 删除记录
- `POST /api/v1/digital-employee/message-records/batch` - 批量操作
- `GET /api/v1/digital-employee/conversations/:id/messages` - 获取会话消息
- `GET /api/v1/digital-employee/conversation-groups` - 获取会话分组
- `GET /api/v1/digital-employee/message-statistics` - 获取统计数据

### 服务回传记录
- `GET /api/v1/digital-employee/export-records` - 获取列表
- `GET /api/v1/digital-employee/export-records/:id` - 获取详情
- `POST /api/v1/digital-employee/export-records` - 创建记录
- `PUT /api/v1/digital-employee/export-records/:id` - 更新记录
- `DELETE /api/v1/digital-employee/export-records/:id` - 删除记录
- `POST /api/v1/digital-employee/export-records/batch` - 批量操作
- `POST /api/v1/digital-employee/export-records/:id/retry` - 重试发送
- `GET /api/v1/digital-employee/export-statistics` - 获取统计数据

## 权限配置

在RBAC系统中添加以下权限：

```sql
-- 历史会话记录权限
INSERT INTO wecom_permissions (intent, intent2, description) VALUES
('digital-employee', 'message-records', '历史会话记录管理'),
('digital-employee', 'message-records.view', '查看会话记录'),
('digital-employee', 'message-records.create', '创建会话记录'),
('digital-employee', 'message-records.update', '更新会话记录'),
('digital-employee', 'message-records.delete', '删除会话记录');

-- 服务回传记录权限
INSERT INTO wecom_permissions (intent, intent2, description) VALUES
('digital-employee', 'export-records', '服务回传记录管理'),
('digital-employee', 'export-records.view', '查看导出记录'),
('digital-employee', 'export-records.create', '创建导出记录'),
('digital-employee', 'export-records.update', '更新导出记录'),
('digital-employee', 'export-records.delete', '删除导出记录'),
('digital-employee', 'export-records.retry', '重试发送邮件');
```

## 测试建议

### 后端测试
1. 单元测试 Repository 层
2. 集成测试 Service 层
3. API 端点测试

### 前端测试
1. 组件单元测试
2. E2E 测试关键流程
3. 性能测试（大数据量）

## 部署注意事项

1. 确保数据库表已创建
2. 配置文件路径权限（导出文件存储）
3. 邮件服务配置（SMTP设置）
4. 日志级别配置
5. 性能监控和告警

## 后续优化方向

1. 实现实时消息推送（WebSocket）
2. 添加数据导出功能（Excel/CSV）
3. 实现高级搜索（Elasticsearch集成）
4. 添加数据归档功能
5. 实现邮件模板管理
6. 添加数据可视化大屏
7. 实现智能推荐和分析

## 开发进度

- [x] 后端Model层
- [x] 后端Repository层
- [x] 后端Service层
- [x] 后端Handler层
- [x] 前端API定义
- [ ] 后端路由注册
- [ ] 依赖注入配置
- [ ] 前端页面实现
- [ ] 前端路由配置
- [ ] 权限集成
- [ ] 测试
- [ ] 文档完善
