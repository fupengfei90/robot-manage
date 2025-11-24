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
