import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'dashboard',
    component: () => import('../views/dashboard/DashboardView.vue')
  },
  {
    path: '/cmdb/duty',
    name: 'dutySchedule',
    component: () => import('../views/cmdb/DutyScheduleView.vue')
  },
  {
    path: '/cmdb/milestone',
    name: 'milestone',
    component: () => import('../views/cmdb/MilestoneView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
