<template>
  <div class="login-container">
    <!-- 背景装饰 -->
    <div class="login__bg-decoration"></div>
    
    <!-- 登录卡片 -->
    <div class="login-card glass-effect animate-scale-in">
      <div class="login-card__header">
        <div class="login-logo">
          <div class="login-logo__icon">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M12 2L2 7L12 12L22 7L12 2Z" stroke="currentColor" stroke-width="2" fill="none"/>
              <path d="M2 17L12 22L22 17" stroke="currentColor" stroke-width="2" fill="none"/>
              <path d="M2 12L12 17L22 12" stroke="currentColor" stroke-width="2" fill="none"/>
            </svg>
          </div>
          <div class="login-logo__text">
            <h1 class="login-logo__name">Robot Manage</h1>
            <p class="login-logo__subtitle">小娇智能运维平台</p>
          </div>
        </div>
      </div>

      <el-form
        ref="loginFormRef"
        :model="loginForm"
        :rules="loginRules"
        class="login-form"
        @submit.prevent="handleLogin"
      >
        <el-form-item prop="username">
          <el-input
            v-model="loginForm.username"
            :placeholder="t('auth.username')"
            size="large"
            :prefix-icon="User"
            class="login-input"
            clearable
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            type="password"
            :placeholder="t('auth.password')"
            size="large"
            :prefix-icon="Lock"
            class="login-input"
            show-password
            @keyup.enter="handleLogin"
          />
        </el-form-item>

        <el-form-item>
          <el-checkbox v-model="loginForm.remember" class="remember-checkbox">
            {{ t('auth.rememberMe') }}
          </el-checkbox>
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            size="large"
            class="login-button"
            :loading="loading"
            @click="handleLogin"
          >
            {{ loading ? '...' : t('auth.login') }}
          </el-button>
        </el-form-item>
      </el-form>

      <div v-if="errorMessage" class="error-message">
        <el-alert
          :title="errorMessage"
          type="error"
          :closable="false"
          show-icon
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { User, Lock } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'
import { useAuthStore } from '../../stores/auth'
import { ElMessage } from 'element-plus'
import { useCommon } from '../../composables/useCommon'

const { t } = useCommon()

const router = useRouter()
const authStore = useAuthStore()

const loginFormRef = ref<FormInstance>()
const loading = ref(false)
const errorMessage = ref('')

const loginForm = reactive({
  username: '',
  password: '',
  remember: false
})

const loginRules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
  ]
}

const handleLogin = async () => {
  if (!loginFormRef.value) return

  await loginFormRef.value.validate(async (valid) => {
    if (!valid) return

    loading.value = true
    errorMessage.value = ''

    try {
      await authStore.login({
        username: loginForm.username,
        password: loginForm.password
      })

      if (loginForm.remember) {
        localStorage.setItem('remember_username', loginForm.username)
      } else {
        localStorage.removeItem('remember_username')
      }

      ElMessage.success(t('auth.loginSuccess'))
      router.push('/')
    } catch (error: any) {
      errorMessage.value = error.message || t('auth.loginFailed')
    } finally {
      loading.value = false
    }
  })
}

// 检查是否有记住的用户名
const rememberedUsername = localStorage.getItem('remember_username')
if (rememberedUsername) {
  loginForm.username = rememberedUsername
  loginForm.remember = true
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  background: var(--gradient-bg);
  padding: var(--spacing-xl);
  overflow: hidden;
}

/* 背景装饰 */
.login__bg-decoration {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: 
    radial-gradient(circle at 20% 50%, rgba(124, 58, 237, 0.15) 0%, transparent 50%),
    radial-gradient(circle at 80% 80%, rgba(59, 130, 246, 0.15) 0%, transparent 50%);
  pointer-events: none;
  z-index: 0;
}

/* 登录卡片 */
.login-card {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 440px;
  padding: var(--spacing-2xl);
  background: var(--gradient-card);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-card);
  transition: all var(--transition-base);
}

.login-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-card-hover);
  border-color: var(--border-color-active);
}

.login-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: var(--gradient-accent);
  border-radius: var(--radius-xl) var(--radius-xl) 0 0;
}

/* Logo区域 */
.login-card__header {
  margin-bottom: var(--spacing-2xl);
  text-align: center;
}

.login-logo {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-md);
}

.login-logo__icon {
  width: 64px;
  height: 64px;
  color: var(--accent-color);
  display: flex;
  align-items: center;
  justify-content: center;
  filter: drop-shadow(0 0 12px rgba(124, 58, 237, 0.5));
}

.login-logo__icon svg {
  width: 100%;
  height: 100%;
}

.login-logo__text {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--spacing-xs);
}

.login-logo__name {
  font-size: 1.75rem;
  font-weight: 700;
  margin: 0;
  background: var(--gradient-accent);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1.2;
}

.login-logo__subtitle {
  font-size: 0.875rem;
  color: var(--text-muted);
  font-weight: 500;
  margin: 0;
}

/* 表单样式 */
.login-form {
  margin-top: var(--spacing-xl);
}

.login-form :deep(.el-form-item) {
  margin-bottom: var(--spacing-lg);
}

.login-form :deep(.el-form-item__label) {
  color: var(--text-secondary);
  font-weight: 500;
}

.login-input :deep(.el-input__wrapper) {
  background: var(--bg-glass);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  box-shadow: none;
  transition: all var(--transition-base);
}

.login-input :deep(.el-input__wrapper:hover) {
  border-color: var(--border-color-hover);
  background: var(--bg-glass-hover);
}

.login-input :deep(.el-input__wrapper.is-focus) {
  border-color: var(--accent-color);
  box-shadow: 0 0 0 2px var(--accent-color-light);
}

.login-input :deep(.el-input__inner) {
  color: var(--text-primary);
}

.login-input :deep(.el-input__inner::placeholder) {
  color: var(--text-muted);
}

.remember-checkbox {
  margin: 0;
}

.remember-checkbox :deep(.el-checkbox__label) {
  color: var(--text-secondary);
  font-size: 0.875rem;
}

.remember-checkbox :deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
  background-color: var(--accent-color);
  border-color: var(--accent-color);
}

/* 登录按钮 */
.login-button {
  width: 100%;
  height: 48px;
  font-size: 1rem;
  font-weight: 600;
  background: var(--gradient-accent);
  border: none;
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-md);
  transition: all var(--transition-base);
}

.login-button:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-glow);
  background: var(--gradient-accent);
}

.login-button:active {
  transform: translateY(0);
}

/* 错误提示 */
.error-message {
  margin-top: var(--spacing-lg);
  animation: fadeIn 0.3s ease;
}

.error-message :deep(.el-alert) {
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: var(--radius-md);
}

.error-message :deep(.el-alert__title) {
  color: var(--color-danger);
  font-size: 0.875rem;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .login-card {
    padding: var(--spacing-xl);
    max-width: 100%;
  }

  .login-logo__icon {
    width: 48px;
    height: 48px;
  }

  .login-logo__name {
    font-size: 1.5rem;
  }
}
</style>

