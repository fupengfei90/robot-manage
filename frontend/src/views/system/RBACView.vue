<template>
  <div class="rbac-view">
    <div class="view-header glass-effect animate-fade-in-down">
      <div class="header-content">
        <h1 class="view-title">{{ t('system.rbac.title') }}</h1>
        <p class="view-subtitle">管理系统用户、角色和权限，实现基于角色的访问控制</p>
      </div>
    </div>

    <!-- 标签页 -->
    <el-card shadow="never" class="tabs-card glass-effect animate-fade-in">
      <el-tabs v-model="activeTab" class="rbac-tabs">
        <el-tab-pane :label="t('system.rbac.user')" name="users">
          <UserManagement />
        </el-tab-pane>
        <el-tab-pane :label="t('system.rbac.role')" name="roles">
          <RoleManagement />
        </el-tab-pane>
        <el-tab-pane :label="t('system.rbac.permission')" name="permissions">
          <PermissionManagement />
        </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import UserManagement from '@/components/rbac/UserManagement.vue'
import RoleManagement from '@/components/rbac/RoleManagement.vue'
import PermissionManagement from '@/components/rbac/PermissionManagement.vue'
import { useCommon } from '../../composables/useCommon'

const { t } = useCommon()
const activeTab = ref('users')
</script>

<style scoped>
.rbac-view {
  padding: var(--spacing-md);
}

.view-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: var(--radius-xl);
}

.tabs-card {
  border-radius: var(--radius-xl);
}

.tabs-card :deep(.el-card__body) {
  padding: var(--spacing-md);
}

.rbac-tabs :deep(.el-tabs__header) {
  margin-bottom: var(--spacing-md);
}

.rbac-tabs :deep(.el-tabs__nav-wrap::after) {
  background: var(--border-color);
}

.rbac-tabs :deep(.el-tabs__item) {
  color: var(--text-secondary);
  font-weight: 500;
}

.rbac-tabs :deep(.el-tabs__item.is-active) {
  color: var(--accent-color);
}

.rbac-tabs :deep(.el-tabs__active-bar) {
  background: var(--accent-color);
}
</style>