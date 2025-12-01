# 剩余视图国际化更新指南

## 当前状态
✅ 已完成：MainLayout, DutyScheduleView, RBACView, LoginView, Dashboard组件
⏳ 待更新：4个视图文件

## 快速更新步骤

### 1. 在 `<script setup>` 中添加导入
```typescript
import { useCommon } from '../../composables/useCommon'
const { t } = useCommon()
```

### 2. 替换常用文本（查找替换）

#### 通用按钮和操作
- `新增` → `{{ t('common.add') }}`
- `编辑` → `{{ t('common.edit') }}`
- `删除` → `{{ t('common.delete') }}`
- `查询` → `{{ t('common.search') }}`
- `重置` → `{{ t('common.reset') }}`
- `确定` → `{{ t('common.confirm') }}`
- `取消` → `{{ t('common.cancel') }}`
- `操作` → `:label="t('common.actions')"`
- `状态` → `:label="t('common.status')"`

#### 待更新文件及其特定翻译键

**cmdb/MilestoneView.vue**
- `大事记管理` → `{{ t('cmdb.milestone.title') }}`
- `事件日期` → `:label="t('cmdb.milestone.eventDate')"`
- `事件内容` → `:label="t('cmdb.milestone.eventContent')"`
- `事件类型` → `:label="t('cmdb.milestone.eventType')"`
- `相关系统` → `:label="t('cmdb.milestone.relatedSystem')"`
- `处理人` → `:label="t('cmdb.milestone.handler')"`

**digital-employee/MessageRecordsView.vue**
- `历史会话记录` → `{{ t('digital.message.title') }}`
- `会话ID` → `:label="t('digital.message.conversationId')"`
- `用户消息` → `:label="t('digital.message.userMessage')"`
- `助手回复` → `:label="t('digital.message.assistantReply')"`

**digital-employee/ExportRecordsView.vue**
- `服务回传记录` → `{{ t('digital.export.title') }}`
- `服务名称` → `:label="t('digital.export.serviceName')"`
- `回传时间` → `:label="t('digital.export.exportTime')"`
- `接收人` → `:label="t('digital.export.recipient')"`

**system/ScheduleTaskView.vue**
- `调度任务管理` → `{{ t('system.task.title') }}`
- `任务名称` → `:label="t('system.task.taskName')"`
- `任务类型` → `:label="t('system.task.taskType')"`
- `Cron表达式` → `:label="t('system.task.cronExpression')"`
- `上次执行时间` → `:label="t('system.task.lastExecuteTime')"`
- `下次执行时间` → `:label="t('system.task.nextExecuteTime')"`
- `立即执行` → `{{ t('system.task.executeNow') }}`

## 注意事项
1. 表格列的 label 属性需要使用 `:label` 绑定
2. 按钮文本直接使用 `{{ t('key') }}`
3. computed 中的文本也要使用 `t('key')`
4. ElMessage 提示信息使用 `t('common.xxxSuccess')`

## 验证
更新后切换语言按钮（EN/中），检查文本是否正确切换。
