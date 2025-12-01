<template>
  <div class="layout">
    <!-- 背景装饰 -->
    <div class="layout__bg-decoration"></div>
    
    <!-- 侧边栏 -->
    <el-aside :width="isCollapsed ? '80px' : '260px'" class="layout__aside glass-effect animate-slide-in-left" :class="{ 'is-collapsed': isCollapsed }">
      <div class="logo-container">
        <div class="logo">
          <div class="logo__icon">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M12 2L2 7L12 12L22 7L12 2Z" stroke="currentColor" stroke-width="2" fill="none"/>
              <path d="M2 17L12 22L22 17" stroke="currentColor" stroke-width="2" fill="none"/>
              <path d="M2 12L12 17L22 12" stroke="currentColor" stroke-width="2" fill="none"/>
            </svg>
          </div>
          <div class="logo__text">
            <span class="logo__name">Robot</span>
            <span class="logo__subtitle">Manage</span>
          </div>
        </div>
      </div>
      <el-button 
        class="collapse-btn" 
        @click="isCollapsed = !isCollapsed"
        circle
      >
        {{ isCollapsed ? '→' : '←' }}
      </el-button>
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapsed"
        class="layout__menu"
        :background-color="themeStore.mode === 'dark' ? 'transparent' : 'transparent'"
        :text-color="themeStore.mode === 'dark' ? 'var(--text-tertiary)' : 'var(- -text-secondary)'"
        :active-text-color="'var(--accent-color)'"
        @select="handleMenuSelect"
      >
        <el-menu-item index="dashboard" class="menu-item" @click="$router.push('/')">
          <template #title>
            <span class="menu-item__icon">📊</span>
            <span>{{ t('layout.dashboard') }}</span>
          </template>
        </el-menu-item>
        <el-sub-menu index="cmdb" class="menu-item">
          <template #title>
            <span class="menu-item__icon">🗄️</span>
            <span>{{ t('layout.cmdb') }}</span>
          </template>
          <el-menu-item index="dutySchedule" @click="$router.push('/cmdb/duty')">
            <template #title>
              <span class="menu-item__icon">📅</span>
              <span>{{ t('layout.dutySchedule') }}</span>
            </template>
          </el-menu-item>
          <el-menu-item index="milestone" @click="$router.push('/cmdb/milestone')">
            <template #title>
              <span class="menu-item__icon">🎯</span>
              <span>{{ t('layout.milestone') }}</span>
            </template>
          </el-menu-item>
          <el-menu-item index="batchTime" @click="$router.push('/cmdb/batch-time')">
            <template #title>
              <span class="menu-item__icon">⏱️</span>
              <span>{{ t('layout.batchTime') }}</span>
            </template>
          </el-menu-item>
          <el-menu-item index="wbCMDB" @click="$router.push('/cmdb/wb-cmdb')">
            <template #title>
              <span class="menu-item__icon">🖥️</span>
              <span>WB CMDB</span>
            </template>
          </el-menu-item>
          <el-menu-item index="vbCMDB" @click="$router.push('/cmdb/vb-cmdb')">
            <template #title>
              <span class="menu-item__icon">💻</span>
              <span>VB CMDB</span>
            </template>
          </el-menu-item>
          <el-menu-item index="itsmPackage" @click="$router.push('/cmdb/itsm-package')">
            <template #title>
              <span class="menu-item__icon">📦</span>
              <span>ITSM出包记录</span>
            </template>
          </el-menu-item>
        </el-sub-menu>
        <el-menu-item index="delivery" class="menu-item">
          <template #title>
            <span class="menu-item__icon">🚀</span>
            <span>{{ t('layout.delivery') }}</span>
          </template>
        </el-menu-item>
        <el-sub-menu index="digital" class="menu-item">
          <template #title>
            <span class="menu-item__icon">🤖</span>
            <span>{{ t('layout.digital') }}</span>
          </template>
          <el-menu-item index="messageRecords" @click="$router.push('/digital/message-records')">
            <template #title>
              <span class="menu-item__icon">💬</span>
              <span>{{ t('layout.messageRecords') }}</span>
            </template>
          </el-menu-item>
          <el-menu-item index="exportRecords" @click="$router.push('/digital/export-records')">
            <template #title>
              <span class="menu-item__icon">📤</span>
              <span>{{ t('layout.exportRecords') }}</span>
            </template>
          </el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="system" class="menu-item">
          <template #title>
            <span class="menu-item__icon">⚙️</span>
            <span>{{ t('layout.system') }}</span>
          </template>
          <el-menu-item index="scheduleTask" @click="$router.push('/system/schedule-tasks')">
            <template #title>
              <span class="menu-item__icon">⏰</span>
              <span>{{ t('layout.scheduleTask') }}</span>
            </template>
          </el-menu-item>
          <el-menu-item index="rbac" @click="$router.push('/system/rbac')">
            <template #title>
              <span class="menu-item__icon">👥</span>
              <span>{{ t('layout.rbac') }}</span>
            </template>
          </el-menu-item>
          <el-menu-item index="weeklyReport" @click="$router.push('/system/weekly-report')">
            <template #title>
              <span class="menu-item__icon">📄</span>
              <span>{{ t('layout.weeklyReport') }}</span>
            </template>
          </el-menu-item>
        </el-sub-menu>
        <el-menu-item index="digitalHuman" class="menu-item" @click="$router.push('/system/digital-human')">
          <template #title>
            <span class="menu-item__icon">👤</span>
            <span>{{ t('layout.digitalHuman') }}</span>
          </template>
        </el-menu-item>
        <el-menu-item index="userGuide" class="menu-item" @click="router.push('/help/user-guide')">
          <template #title>
            <span class="menu-item__icon">❓</span>
            <span>{{ t('layout.userGuide') }}</span>
          </template>
        </el-menu-item>
        <el-menu-item index="projectPlan" class="menu-item" @click="router.push('/help/project-plan')">
          <template #title>
            <span class="menu-item__icon">📊</span>
            <span>{{ t('layout.projectPlan') }}</span>
          </template>
        </el-menu-item>
      </el-menu>
    </el-aside>
    
    <!-- 主内容区 -->
    <div class="layout__content">
      <!-- 顶部导航栏 -->
      <header class="layout__header glass-effect animate-fade-in-down">
        <div class="header__left">
          <div class="header__title-group">
            <h1 class="header__title">{{ t('layout.title') }}</h1>
            <p class="header__subtitle">{{ t('layout.subtitle') }}</p>
          </div>
        </div>
        <div class="layout__header-actions">
          <el-button
            :icon="themeIcon"
            circle
            class="action-btn hover-lift"
            @click="themeStore.toggle()"
            :title="t('layout.toggleTheme')"
          />
          <el-button
            circle
            class="action-btn hover-lift"
            @click="i18n.toggle()"
            :title="t('layout.toggleLanguage')"
          >
            {{ i18n.locale === 'zh' ? 'EN' : '中' }}
          </el-button>
          <el-tag type="success" class="status-tag">Alpha</el-tag>
          
          <!-- 用户信息 -->
          <el-dropdown v-if="authStore.user" @command="handleLogout" class="user-dropdown">
            <div class="user-info hover-lift">
              <el-avatar :size="32" class="user-avatar">
                <el-icon><User /></el-icon>
              </el-avatar>
              <span class="username">{{ authStore.user.username }}</span>
              <el-icon class="dropdown-icon"><ArrowDown /></el-icon>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="logout" divided>
                  <el-icon><User /></el-icon>
                  {{ t('layout.logout') }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </header>
      
      <!-- 主内容 -->
      <main class="layout__main animate-fade-in">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Sunny, Moon, User, ArrowDown } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useThemeStore } from '../stores/theme'
import { useAuthStore } from '../stores/auth'
import { useI18n } from '../composables/useI18n'

const route = useRoute()
const router = useRouter()
const themeStore = useThemeStore()
const authStore = useAuthStore()
const i18n = useI18n()
const { t } = i18n
const activeMenu = computed(() => route.name?.toString() ?? 'dashboard')
const isCollapsed = ref(false)

const themeIcon = computed(() => themeStore.mode === 'dark' ? Sunny : Moon)

const handleMenuSelect = (index: string) => {
  const routes: Record<string, string> = {
    dashboard: '/',
    dutySchedule: '/cmdb/duty',
    milestone: '/cmdb/milestone',
    batchTime: '/cmdb/batch-time',
    wbCMDB: '/cmdb/wb-cmdb',
    vbCMDB: '/cmdb/vb-cmdb',
    itsmPackage: '/cmdb/itsm-package',
    messageRecords: '/digital/message-records',
    exportRecords: '/digital/export-records',
    scheduleTask: '/system/schedule-tasks',
    rbac: '/system/rbac',
    projectPlan: '/help/project-plan',
    weeklyReport: '/system/weekly-report',
    userGuide: '/help/user-guide',
    digitalHuman: '/system/digital-human'
  }
  if (routes[index]) {
    router.push(routes[index])
  }
}

const handleLogout = async () => {
  try {
    await ElMessageBox.confirm(
      t('layout.confirmLogout'),
      t('layout.logout'),
      {
        confirmButtonText: t('common.confirm'),
        cancelButtonText: t('common.cancel'),
        type: 'warning'
      }
    )
    await authStore.logout()
    ElMessage.success(t('common.operationSuccess'))
    router.push('/login')
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(t('auth.loginFailed'))
    }
  }
}

onMounted(async () => {
  themeStore.applyTheme()
  if (authStore.token && !authStore.user) {
    try {
      await authStore.fetchUserInfo()
    } catch (error) {
      console.error('Failed to fetch user info:', error)
    }
  }
})
</script>

<style scoped>
.layout {
  display: flex;
  min-height: 100vh;
  position: relative;
  background: var(--gradient-bg);
  color: var(--text-primary);
  overflow: hidden;
}

/* 背景装饰 */
.layout__bg-decoration {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: 
    radial-gradient(circle at 20% 50%, rgba(124, 58, 237, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 80% 80%, rgba(59, 130, 246, 0.1) 0%, transparent 50%);
  pointer-events: none;
  z-index: 0;
}

/* 侧边栏 */
.layout__aside {
  position: relative;
  z-index: 10;
  padding: var(--spacing-xl) var(--spacing-lg);
  border-right: 1px solid var(--border-color);
  box-shadow: var(--shadow-lg);
  min-height: 100vh;
  transition: width 0.3s ease;
}

.layout__aside.is-collapsed {
  padding: var(--spacing-xl) var(--spacing-sm);
}

.layout__aside.is-collapsed .logo__text {
  display: none;
}



.collapse-btn {
  position: absolute;
  right: -15px;
  top: 20px;
  z-index: 20;
  background: var(--bg-glass);
  border: 1px solid var(--border-color);
  color: var(--text-primary);
  width: 30px;
  height: 30px;
  padding: 0;
  font-size: 14px;
  transition: all var(--transition-base);
}

.collapse-btn:hover {
  background: var(--bg-glass-hover);
  border-color: var(--accent-color);
  color: var(--accent-color);
  transform: scale(1.1);
}

.logo-container {
  margin-bottom: var(--spacing-2xl);
}

.logo {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-md);
  border-radius: var(--radius-lg);
  background: var(--gradient-card);
  transition: all var(--transition-base);
}

.logo:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-glow);
}

.logo__icon {
  width: 40px;
  height: 40px;
  color: var(--accent-color);
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo__icon svg {
  width: 100%;
  height: 100%;
  filter: drop-shadow(0 0 8px rgba(124, 58, 237, 0.5));
}

.logo__text {
  display: flex;
  flex-direction: column;
}

.logo__name {
  font-size: 1.25rem;
  font-weight: 700;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1.2;
}

.logo__subtitle {
  font-size: 0.875rem;
  color: var(--text-muted);
  font-weight: 500;
}

.layout__menu {
  border: none;
  background: transparent !important;
}

.menu-item {
  margin-bottom: var(--spacing-xs);
  border-radius: var(--radius-md);
  transition: all var(--transition-base);
  position: relative;
  overflow: hidden;
}

.menu-item::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  width: 3px;
  background: var(--gradient-accent);
  transform: scaleY(0);
  transition: transform var(--transition-base);
}

.menu-item.is-active::before {
  transform: scaleY(1);
}

.menu-item:hover {
  background: var(--bg-glass-hover) !important;
  transform: translateX(4px);
}

.menu-item__icon {
  font-size: 1.25rem;
  margin-right: var(--spacing-sm);
}

:deep(.el-menu--collapse) .menu-item__icon {
  margin-right: 0;
}



.layout__aside.is-collapsed .logo {
  justify-content: center;
}

/* 主内容区 */
.layout__content {
  flex: 1;
  position: relative;
  z-index: 1;
  padding: var(--spacing-xl);
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

/* 顶部导航栏 */
.layout__header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md) var(--spacing-lg);
  margin-bottom: var(--spacing-md);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-md);
  position: sticky;
  top: var(--spacing-lg);
  z-index: 100;
}

.header__left {
  flex: 1;
}

.header__title-group {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.header__title {
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.header__subtitle {
  font-size: 0.875rem;
  color: var(--text-muted);
  font-weight: 400;
}

.layout__header-actions {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.action-btn {
  background: var(--bg-glass);
  border: 1px solid var(--border-color);
  color: var(--text-primary);
  transition: all var(--transition-base);
}

.action-btn:hover {
  background: var(--bg-glass-hover);
  border-color: var(--accent-color);
  color: var(--accent-color);
  transform: translateY(-2px);
  box-shadow: var(--shadow-glow);
}

.status-tag {
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-full);
  font-weight: 500;
  box-shadow: var(--shadow-sm);
}

/* 用户信息样式 */
.user-dropdown {
  cursor: pointer;
}

.user-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-md);
  background: var(--bg-glass);
  border: 1px solid var(--border-color);
  transition: all var(--transition-base);
}

.user-info:hover {
  background: var(--bg-glass-hover);
  border-color: var(--accent-color);
  transform: translateY(-2px);
  box-shadow: var(--shadow-glow);
}

.user-avatar {
  background: var(--gradient-accent);
  color: white;
}

.username {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--text-primary);
  max-width: 100px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.dropdown-icon {
  font-size: 12px;
  color: var(--text-muted);
  transition: transform var(--transition-base);
}

.user-dropdown:hover .dropdown-icon {
  transform: rotate(180deg);
}

/* 主内容 */
.layout__main {
  flex: 1;
  position: relative;
  z-index: 1;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .layout__aside {
    width: 80px !important;
    padding: var(--spacing-md);
  }
  
  .logo__text {
    display: none;
  }
  
  .menu-item span:not(.menu-item__icon) {
    display: none;
  }
  
  .layout__content {
    padding: var(--spacing-md);
  }
  
  .header__title {
    font-size: 1.25rem;
  }
  
  .header__subtitle {
    display: none;
  }
  
  .username {
    display: none;
  }
  
  .user-info {
    padding: var(--spacing-sm);
  }
}
</style>
