import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/auth/LoginView.vue'),
    meta: { requiresAuth: false, layout: false }
  },
  {
    path: '/',
    name: 'dashboard',
    component: () => import('../views/dashboard/DashboardView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/cmdb/duty',
    name: 'dutySchedule',
    component: () => import('../views/cmdb/DutyScheduleView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/cmdb/milestone',
    name: 'milestone',
    component: () => import('../views/cmdb/MilestoneView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/cmdb/batch-time',
    name: 'batchTime',
    component: () => import('../views/cmdb/BatchTimeView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/cmdb/wb-cmdb',
    name: 'wbCMDB',
    component: () => import('../views/cmdb/WBCMDBView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/cmdb/vb-cmdb',
    name: 'vbCMDB',
    component: () => import('../views/cmdb/VBCMDBView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/cmdb/itsm-package',
    name: 'itsmPackage',
    component: () => import('../views/cmdb/ITSMPackageView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/system/schedule-tasks',
    name: 'scheduleTask',
    component: () => import('../views/system/ScheduleTaskView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/system/rbac',
    name: 'rbac',
    component: () => import('../views/system/RBACView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/digital/message-records',
    name: 'messageRecords',
    component: () => import('../views/digital-employee/MessageRecordsView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/digital/export-records',
    name: 'exportRecords',
    component: () => import('../views/digital-employee/ExportRecordsView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/help/user-guide',
    name: 'userGuide',
    component: () => import('../views/help/UserGuideView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/help/project-plan',
    name: 'projectPlan',
    component: () => import('../views/system/ProjectPlanView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/system/weekly-report',
    name: 'weeklyReport',
    component: () => import('../views/system/WeeklyReportView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/system/digital-human',
    name: 'digitalHuman',
    component: () => import('../views/system/DigitalHumanView.vue'),
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  const requiresAuth = to.meta.requiresAuth !== false

  // 检查是否有token
  if (requiresAuth) {
    if (!authStore.checkAuth()) {
      next({ name: 'login', query: { redirect: to.fullPath } })
      return
    }
  } else {
    // 如果已登录，访问登录页时重定向到首页
    if (to.name === 'login' && authStore.checkAuth()) {
      next('/')
      return
    }
  }

  next()
})

export default router
