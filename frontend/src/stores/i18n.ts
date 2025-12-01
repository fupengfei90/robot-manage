import { defineStore } from 'pinia'

export type Locale = 'zh' | 'en'

interface Translations {
  [key: string]: string | Translations
}

const translations: Record<Locale, Translations> = {
  zh: {
    projectPlan: {
      title: '项目规划与进度',
      subtitle: '追踪项目开发进度和完成情况',
      completed: '已完成',
      progress: '完成进度'
    },
    common: {
      search: '搜索',
      add: '新增',
      edit: '编辑',
      delete: '删除',
      save: '保存',
      cancel: '取消',
      confirm: '确认',
      submit: '提交',
      reset: '重置',
      export: '导出',
      import: '导入',
      refresh: '刷新',
      operation: '操作',
      status: '状态',
      enabled: '启用',
      disabled: '禁用',
      actions: '操作',
      batchDelete: '批量删除',
      pleaseSelect: '请选择',
      confirmDelete: '确认删除',
      deleteSuccess: '删除成功',
      saveSuccess: '保存成功',
      operationSuccess: '操作成功'
    },
    layout: {
      title: '小娇智能运维平台',
      subtitle: '一站式运维指挥中枢',
      dashboard: '仪表盘',
      cmdb: 'CMDB',
      delivery: '智能交付',
      digital: '数字员工',
      system: '系统管理',
      toggleTheme: '切换主题',
      toggleLanguage: '切换语言',
      dutySchedule: '值班管理',
      milestone: '大事记',
      batchTime: '批量时间',
      messageRecords: '历史会话记录',
      exportRecords: '服务回传记录',
      scheduleTask: '调度任务管理',
      rbac: '用户角色权限',
      userGuide: '小娇使用方法',
      projectPlan: '项目规划',
      weeklyReport: '周报汇报',
      digitalHuman: '数字人形象',
      logout: '退出登录',
      confirmLogout: '确定要退出登录吗？'
    },
    dashboard: {
      metrics: {
        title: '运营指标',
        subtitle: '服务能力与巡检告警总览',
        realtime: '实时刷新',
        daily: '日服务次数',
        weekly: '周服务次数',
        monthly: '月服务次数',
        yearly: '年服务次数',
        users: '在册用户',
        serviceUsers: '服务用户数',
        serviceTrend: '服务次数趋势',
        inspectionTrend: '巡检分析次数',
        alertLevel: '告警分析次数',
        weekSummary: '本周汇总',
        monthSummary: '本月汇总',
        userService: '用户服务',
        scheduledTask: '定时任务'
      },
      knowledge: {
        title: '知识库动态',
        recentUpdates: '最近更新',
        hotTopics: '热门知识',
        author: '作者',
        views: '阅读',
        recallRate: '召回'
      },
      timeline: {
        title: '值班与大事记',
        dutyOfficer: '今日值班',
        milestones: '大事记',
        deployPlans: '投产计划',
        system: '版本',
        window: '内容',
        owner: '负责人',
        viewSchedule: '查看排班'
      },
      assistant: {
        title: '助手总结',
        range: '范围',
        highlights: '亮点',
        risks: '风险',
        nextSteps: '下一步',
        daily: '日报',
        weekly: '周报',
        monthly: '月报',
        yearly: '年报'
      },
      digitalHuman: {
        title: '数字人助手',
        speak: '语音播报'
      }
    },
    cmdb: {
      duty: {
        title: '值班排班管理',
        subtitle: '管理WB和FB值班人员信息',
        dutyDate: '值班日期',
        startDate: '开始日期',
        endDate: '结束日期',
        staffName: '值班人',
        wbStaff: 'WB值班人员',
        fbStaff: 'FB值班人员',
        addSchedule: '新增排班',
        editSchedule: '编辑排班',
        createdAt: '创建时间'
      },
      milestone: {
        title: '大事记管理',
        eventDate: '事件日期',
        dayOfWeek: '星期',
        eventContent: '事件内容',
        eventType: '事件类型',
        relatedSystem: '相关系统',
        handler: '处理人',
        addEvent: '新增大事记',
        editEvent: '编辑大事记'
      },
      batchTime: {
        title: '批量时间管理',
        subtitle: '管理系统批处理时间窗口配置',
        systemName: '系统名称',
        subsysName: '子系统名称',
        batchTime: '批量时间',
        batchWindow: '批量时间窗口',
        addConfig: '新增配置',
        editConfig: '编辑配置',
        updatedAt: '更新时间'
      }
    },
    digital: {
      message: {
        title: '历史会话记录',
        conversationId: '会话ID',
        userMessage: '用户消息',
        assistantReply: '助手回复',
        timestamp: '时间戳',
        conversationGroup: '会话分组'
      },
      export: {
        title: '服务回传记录',
        serviceName: '服务名称',
        exportTime: '回传时间',
        status: '状态',
        recipient: '接收人',
        content: '内容',
        retry: '重试'
      }
    },
    system: {
      task: {
        title: '调度任务管理',
        subtitle: '管理系统中的定时任务，支持创建、编辑、启用/停用和立即执行等操作',
        taskName: '任务名称',
        taskType: '任务类型',
        cronExpression: 'Cron表达式',
        status: '状态',
        executionTime: '执行时间',
        nextExecution: '下次执行',
        lastExecuteTime: '上次执行时间',
        nextExecuteTime: '下次执行时间',
        executeNow: '执行',
        addTask: '新增任务',
        editTask: '编辑任务',
        updateTime: '更新时间',
        actions: '操作',
        edit: '编辑',
        delete: '删除',
        batchOperations: '批量操作',
        batchEnable: '批量启用',
        batchDisable: '批量停用',
        batchDelete: '批量删除',
        disabled: '已停用',
        category: '任务分类',
        searchPlaceholder: '请输入任务名称',
        selectStatus: '请选择状态',
        selectCategory: '请选择分类',
        selectDCN: '请选择数据中心',
        startTime: '开始时间',
        endTime: '结束时间',
        dcn: '数据中心',
        creationTime: '创建时间',
        enabled: '启用',
        confirmDelete: '确认删除',
        deleteConfirmMsg: '确定要删除任务"{name}"吗？删除后无法恢复。',
        confirm: '确定',
        cancel: '取消',
        deleteSuccess: '删除成功',
        deleteFailed: '删除失败',
        statusUpdateSuccess: '任务已{status}',
        statusUpdateFailed: '状态更新失败',
        executeSuccess: '任务执行成功',
        executeFailed: '任务执行失败',
        batchOperationConfirm: '确定要{operation}选中的 {count} 个任务吗？',
        batchOperationSuccess: '批量{operation}成功',
        batchOperationFailed: '批量{operation}失败',
        loadDataFailed: '加载数据失败'
      },
      rbac: {
        title: '用户角色权限管理',
        user: '用户管理',
        role: '角色管理',
        permission: '权限管理',
        username: '用户名',
        realName: '真实姓名',
        email: '邮箱',
        phone: '手机号',
        roleName: '角色名称',
        roleCode: '角色编码',
        permissionName: '权限名称',
        permissionCode: '权限编码',
        assignRoles: '分配角色',
        assignPermissions: '分配权限'
      }
    },
    auth: {
      login: '登录',
      username: '用户名',
      password: '密码',
      rememberMe: '记住我',
      loginSuccess: '登录成功',
      loginFailed: '登录失败'
    }
  },
  en: {
    projectPlan: {
      title: 'Project Plan & Progress',
      subtitle: 'Track project development progress and completion',
      completed: 'Completed',
      progress: 'Progress'
    },
    common: {
      search: 'Search',
      add: 'Add',
      edit: 'Edit',
      delete: 'Delete',
      save: 'Save',
      cancel: 'Cancel',
      confirm: 'Confirm',
      submit: 'Submit',
      reset: 'Reset',
      export: 'Export',
      import: 'Import',
      refresh: 'Refresh',
      operation: 'Operation',
      status: 'Status',
      enabled: 'Enabled',
      disabled: 'Disabled',
      actions: 'Actions',
      batchDelete: 'Batch Delete',
      pleaseSelect: 'Please Select',
      confirmDelete: 'Confirm Delete',
      deleteSuccess: 'Delete Success',
      saveSuccess: 'Save Success',
      operationSuccess: 'Operation Success'
    },
    layout: {
      title: 'Intelligent Operations Platform',
      subtitle: 'Enterprise-Grade One-Stop Operations Command Center',
      dashboard: 'Dashboard',
      cmdb: 'CMDB',
      delivery: 'Smart Delivery',
      digital: 'Digital Employee',
      system: 'System Management',
      toggleTheme: 'Toggle Theme',
      toggleLanguage: 'Toggle Language',
      dutySchedule: 'Duty Schedule',
      milestone: 'Milestones',
      batchTime: 'Batch Time',
      messageRecords: 'Message Records',
      exportRecords: 'Export Records',
      scheduleTask: 'Schedule Tasks',
      rbac: 'RBAC',
      userGuide: 'User Guide',
      projectPlan: 'Project Plan',
      weeklyReport: 'Weekly Report',
      digitalHuman: 'Digital Human',
      logout: 'Logout',
      confirmLogout: 'Are you sure to logout?'
    },
    dashboard: {
      metrics: {
        title: 'Operations Metrics',
        subtitle: 'Service Capability & Inspection Alert Overview',
        realtime: 'Real-time Refresh',
        daily: 'Daily Services',
        weekly: 'Weekly Services',
        monthly: 'Monthly Services',
        yearly: 'Yearly Services',
        users: 'Registered Users',
        serviceUsers: 'Service Users',
        serviceTrend: 'Service Trend',
        inspectionTrend: 'Inspection Trend',
        alertLevel: 'Alert Level',
        weekSummary: 'Week Summary',
        monthSummary: 'Month Summary',
        userService: 'User Service',
        scheduledTask: 'Scheduled Task'
      },
      knowledge: {
        title: 'Knowledge Base Updates',
        recentUpdates: 'Recent Updates',
        hotTopics: 'Hot Topics',
        author: 'Author',
        views: 'Views',
        recallRate: 'Recall Rate'
      },
      timeline: {
        title: 'Duty & Milestones',
        dutyOfficer: 'Today\'s Duty Officer',
        milestones: 'Milestones',
        deployPlans: 'Deployment Plans',
        system: 'Version',
        window: 'Window',
        owner: 'Owner',
        viewSchedule: 'View Schedule'
      },
      assistant: {
        title: 'Assistant Summary',
        range: 'Range',
        highlights: 'Highlights',
        risks: 'Risks',
        nextSteps: 'Next Steps',
        daily: 'Daily',
        weekly: 'Weekly',
        monthly: 'Monthly',
        yearly: 'Yearly'
      },
      digitalHuman: {
        title: 'Digital Assistant',
        speak: 'Voice Broadcast'
      }
    },
    cmdb: {
      duty: {
        title: 'Duty Schedule Management',
        subtitle: 'Manage WB and FB duty staff information',
        dutyDate: 'Duty Date',
        startDate: 'Start Date',
        endDate: 'End Date',
        staffName: 'Staff Name',
        wbStaff: 'WB Staff',
        fbStaff: 'FB Staff',
        addSchedule: 'Add Schedule',
        editSchedule: 'Edit Schedule',
        createdAt: 'Created At'
      },
      milestone: {
        title: 'Milestone Management',
        eventDate: 'Event Date',
        dayOfWeek: 'Day of Week',
        eventContent: 'Event Content',
        eventType: 'Event Type',
        relatedSystem: 'Related System',
        handler: 'Handler',
        addEvent: 'Add Event',
        editEvent: 'Edit Event'
      },
      batchTime: {
        title: 'Batch Time Management',
        subtitle: 'Manage system batch processing time window configuration',
        systemName: 'System Name',
        subsysName: 'Subsystem Name',
        batchTime: 'Batch Time',
        batchWindow: 'Batch Time Window',
        addConfig: 'Add Configuration',
        editConfig: 'Edit Configuration',
        updatedAt: 'Updated At'
      }
    },
    digital: {
      message: {
        title: 'Message Records',
        conversationId: 'Conversation ID',
        userMessage: 'User Message',
        assistantReply: 'Assistant Reply',
        timestamp: 'Timestamp',
        conversationGroup: 'Conversation Group'
      },
      export: {
        title: 'Export Records',
        serviceName: 'Service Name',
        exportTime: 'Export Time',
        status: 'Status',
        recipient: 'Recipient',
        content: 'Content',
        retry: 'Retry'
      }
    },
    system: {
      task: {
        title: 'Schedule Task Management',
        subtitle: 'Manage scheduled tasks with create, edit, enable/disable and execute operations',
        taskName: 'Task Name',
        taskType: 'Task Type',
        cronExpression: 'Cron Expression',
        status: 'Status',
        executionTime: 'Execution Time',
        nextExecution: 'Next Execution',
        lastExecuteTime: 'Last Execute Time',
        nextExecuteTime: 'Next Execute Time',
        executeNow: 'Execute',
        addTask: 'Add Task',
        editTask: 'Edit Task',
        updateTime: 'Update Time',
        actions: 'Actions',
        edit: 'Edit',
        delete: 'Delete',
        batchOperations: 'Batch Operations',
        batchEnable: 'Batch Enable',
        batchDisable: 'Batch Disable',
        batchDelete: 'Batch Delete',
        disabled: 'Disabled',
        category: 'Category',
        searchPlaceholder: 'Enter task name',
        selectStatus: 'Select status',
        selectCategory: 'Select category',
        selectDCN: 'Select DCN',
        startTime: 'Start Time',
        endTime: 'End Time',
        dcn: 'DCN',
        creationTime: 'Creation Time',
        enabled: 'Enabled',
        confirmDelete: 'Confirm Delete',
        deleteConfirmMsg: 'Are you sure to delete task "{name}"? This action cannot be undone.',
        confirm: 'Confirm',
        cancel: 'Cancel',
        deleteSuccess: 'Delete Success',
        deleteFailed: 'Delete Failed',
        statusUpdateSuccess: 'Status updated to {status}',
        statusUpdateFailed: 'Status update failed',
        executeSuccess: 'Task executed successfully',
        executeFailed: 'Task execution failed',
        batchOperationConfirm: 'Are you sure to {operation} {count} selected tasks?',
        batchOperationSuccess: 'Batch {operation} success',
        batchOperationFailed: 'Batch {operation} failed',
        loadDataFailed: 'Failed to load data'
      },
      rbac: {
        title: 'RBAC Management',
        user: 'User Management',
        role: 'Role Management',
        permission: 'Permission Management',
        username: 'Username',
        realName: 'Real Name',
        email: 'Email',
        phone: 'Phone',
        roleName: 'Role Name',
        roleCode: 'Role Code',
        permissionName: 'Permission Name',
        permissionCode: 'Permission Code',
        assignRoles: 'Assign Roles',
        assignPermissions: 'Assign Permissions'
      }
    },
    auth: {
      login: 'Login',
      username: 'Username',
      password: 'Password',
      rememberMe: 'Remember Me',
      loginSuccess: 'Login Success',
      loginFailed: 'Login Failed'
    }
  }
}

export const useI18nStore = defineStore('i18n', {
  state: (): { locale: Locale } => ({
    locale: (localStorage.getItem('locale') as Locale) || 'zh'
  }),
  getters: {
    t: (state) => {
      return (key: string): string => {
        const keys = key.split('.')
        let value: any = translations[state.locale]
        for (const k of keys) {
          value = value?.[k]
        }
        return typeof value === 'string' ? value : key
      }
    }
  },
  actions: {
    setLocale(locale: Locale) {
      this.locale = locale
      localStorage.setItem('locale', locale)
    },
    toggle() {
      this.setLocale(this.locale === 'zh' ? 'en' : 'zh')
    }
  }
})

