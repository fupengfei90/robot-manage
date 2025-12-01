# 国际化使用指南

## 已完成
- ✅ i18n store 已添加所有模块的中英文翻译
- ✅ MainLayout 已支持国际化
- ✅ DutyScheduleView 已支持国际化（示例）

## 在视图中使用国际化

### 1. 导入 composable
```typescript
import { useCommon } from '../../composables/useCommon'

const { t } = useCommon()
```

### 2. 替换硬编码文本
```vue
<!-- 之前 -->
<h1>值班管理</h1>
<el-button>新增</el-button>

<!-- 之后 -->
<h1>{{ t('cmdb.duty.title') }}</h1>
<el-button>{{ t('common.add') }}</el-button>
```

### 3. 动态属性使用 :label
```vue
<!-- 之前 -->
<el-table-column label="操作" />

<!-- 之后 -->
<el-table-column :label="t('common.actions')" />
```

## 可用的翻译键

### 通用 (common)
- search, add, edit, delete, save, cancel, confirm, submit, reset
- export, import, refresh, operation, status, actions
- enabled, disabled, batchDelete, pleaseSelect
- confirmDelete, deleteSuccess, saveSuccess, operationSuccess

### CMDB (cmdb.duty / cmdb.milestone)
- duty: title, dutyDate, wbStaff, fbStaff, addSchedule, editSchedule
- milestone: title, eventDate, dayOfWeek, eventContent, eventType, relatedSystem, handler, addEvent, editEvent

### 数字员工 (digital.message / digital.export)
- message: title, conversationId, userMessage, assistantReply, timestamp, conversationGroup
- export: title, serviceName, exportTime, status, recipient, content, retry

### 系统管理 (system.task / system.rbac)
- task: title, taskName, taskType, cronExpression, status, lastExecuteTime, nextExecuteTime, executeNow, addTask, editTask
- rbac: title, user, role, permission, username, realName, email, phone, roleName, roleCode, permissionName, permissionCode, assignRoles, assignPermissions

### 认证 (auth)
- login, username, password, rememberMe, loginSuccess, loginFailed

## 待更新文件
- [ ] cmdb/MilestoneView.vue
- [ ] digital-employee/MessageRecordsView.vue
- [ ] digital-employee/ExportRecordsView.vue
- [ ] system/ScheduleTaskView.vue
- [ ] system/RBACView.vue
- [ ] auth/LoginView.vue
- [ ] help/UserGuideView.vue
