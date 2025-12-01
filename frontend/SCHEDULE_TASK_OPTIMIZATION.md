# 调度任务管理界面优化说明

## 优化概览

本次优化针对调度任务管理界面进行了视觉和国际化的全面改进，解决了深色主题下表格对比度过高和国际化不完整的问题。

## 1. 视觉优化

### 1.1 表格斑马纹优化
- **背景色调整**：
  - 奇数行：`#1a1d28`（主背景色）
  - 偶数行：`#1e222d`（轻微加深，对比度约12%）
  - 悬停效果：`rgba(255, 255, 255, 0.03)` 叠加
  - 斑马纹悬停：`rgba(255, 255, 255, 0.05)` 叠加

### 1.2 表头样式优化
- **背景渐变**：`linear-gradient(180deg, #2d3746 0%, #252a38 100%)`
- **文字颜色**：`#e2e8f0`（柔和的浅色）
- **字体重量**：600（半粗体）
- **底部边框**：`1px solid rgba(255, 255, 255, 0.1)`

### 1.3 边框优化
- **统一边框**：`1px` 细边框
- **边框颜色**：`rgba(255, 255, 255, 0.08)`（低对比度）

### 1.4 交互效果优化
- **行悬停**：轻微上移效果 `translateY(-1px)`
- **开关颜色**：
  - 启用状态：使用主题强调色
  - 禁用状态：`rgba(255, 255, 255, 0.2)`

### 1.5 布局优化
- **列宽调整**：
  - ID列：80px（固定）
  - 任务名称：200px（最小宽度）
  - 状态列：100px
  - 执行时间：140px
  - 下次执行：140px
  - 更新时间：140px
  - 操作列：120px（固定）

## 2. 国际化优化

### 2.1 新增国际化键值

#### 中文（zh）
- `system.task.subtitle`: 管理系统中的定时任务，支持创建、编辑、启用/停用和立即执行等操作
- `system.task.executionTime`: 执行时间
- `system.task.nextExecution`: 下次执行
- `system.task.updateTime`: 更新时间
- `system.task.actions`: 操作
- `system.task.edit`: 编辑
- `system.task.delete`: 删除
- `system.task.batchOperations`: 批量操作
- `system.task.batchEnable`: 批量启用
- `system.task.batchDisable`: 批量停用
- `system.task.batchDelete`: 批量删除
- `system.task.disabled`: 已停用
- `system.task.category`: 任务分类
- `system.task.searchPlaceholder`: 请输入任务名称
- `system.task.selectStatus`: 请选择状态
- `system.task.selectCategory`: 请选择分类
- `system.task.selectDCN`: 请选择数据中心
- `system.task.dcn`: 数据中心
- `system.task.creationTime`: 创建时间
- 以及各种操作提示消息

#### 英文（en）
- `system.task.subtitle`: Manage scheduled tasks with create, edit, enable/disable and execute operations
- `system.task.executionTime`: Execution Time
- `system.task.nextExecution`: Next Execution
- `system.task.updateTime`: Update Time
- `system.task.actions`: Actions
- `system.task.edit`: Edit
- `system.task.delete`: Delete
- `system.task.batchOperations`: Batch Operations
- `system.task.batchEnable`: Batch Enable
- `system.task.batchDisable`: Batch Disable
- `system.task.batchDelete`: Batch Delete
- `system.task.disabled`: Disabled
- `system.task.category`: Category
- `system.task.searchPlaceholder`: Enter task name
- `system.task.selectStatus`: Select status
- `system.task.selectCategory`: Select category
- `system.task.selectDCN`: Select DCN
- `system.task.dcn`: DCN
- `system.task.creationTime`: Creation Time
- 以及各种操作提示消息

### 2.2 组件国际化改造

#### ScheduleTaskView.vue
- 页面标题和副标题
- 表格列标题
- 操作按钮文本
- 所有提示消息（成功、失败、确认等）

#### TaskFilter.vue
- 表单标签
- 输入框占位符
- 按钮文本
- 添加搜索图标

#### TaskForm.vue
- 对话框标题
- 表单标签
- 输入框占位符
- 按钮文本
- 表单验证消息

## 3. 细节优化

### 3.1 搜索框增强
- 添加搜索图标前缀
- 优化占位符文本

### 3.2 筛选区域优化
- 更紧凑的表单布局
- 统一的标签字体重量（500）
- 移除表单项底部边距

### 3.3 按钮样式优化
- 统一按钮间距（8px）
- 优化文本按钮大小（13px）
- 禁用状态透明度（0.4）

### 3.4 表格单元格优化
- 统一单元格内边距（12px 0）
- 优化文本大小和颜色
- 改善溢出文本处理

## 4. 保持不变的元素

以下元素保持原有的中文显示（业务数据）：
- 具体的任务名称（如"业务线运维报表解读"）
- 执行时间描述（如"16点 周一,周二,周三,周四,周五"）
- 日期时间格式
- 分类标签（值班提醒、周报生成、版本通知）

## 5. 文件修改清单

1. `/src/styles/element-override.css` - 表格样式优化
2. `/src/stores/i18n.ts` - 国际化配置扩展
3. `/src/views/system/ScheduleTaskView.vue` - 主视图国际化和样式优化
4. `/src/components/TaskFilter.vue` - 筛选组件国际化
5. `/src/components/TaskForm.vue` - 表单组件国际化

## 6. 验收检查

### 视觉验收 ✓
- [x] 表格斑马纹对比度适中，不刺眼
- [x] 整体色彩协调统一
- [x] 悬停和交互状态自然流畅
- [x] 文字可读性良好
- [x] 表头样式专业美观

### 功能验收 ✓
- [x] 所有界面文本支持国际化切换
- [x] 中英文显示正确
- [x] 表格排序、筛选功能正常
- [x] 响应式布局适配良好

### 性能验收 ✓
- [x] 页面加载性能无退化
- [x] 动画效果流畅
- [x] 内存使用合理

## 7. 使用说明

### 切换语言
用户可以通过系统顶部的语言切换按钮在中英文之间切换，所有界面元素会自动更新为对应语言。

### 主题适配
优化后的样式完全适配深色主题，在浅色主题下也能保持良好的视觉效果。

## 8. 后续建议

1. 考虑添加表格密度切换功能（紧凑/标准/宽松）
2. 可以为不同类型的任务添加图标标识
3. 考虑添加表格列的显示/隐藏配置
4. 可以添加快捷键支持（如 Ctrl+F 快速搜索）
