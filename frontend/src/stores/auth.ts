import { defineStore } from 'pinia'
import { authApi } from '../api/auth'
import type { LoginRequest, UserInfo } from '../types/auth'

interface AuthState {
  token: string | null
  user: UserInfo | null
  isAuthenticated: boolean
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    token: localStorage.getItem('robot_manage_token'),
    user: null,
    isAuthenticated: false
  }),

  getters: {
    isLoggedIn: (state) => !!state.token && state.isAuthenticated
  },

  actions: {
    async login(credentials: LoginRequest) {
      try {
        const response = await authApi.login(credentials)
        
        this.token = response.token
        this.user = response.user
        this.isAuthenticated = true

        // 保存token到localStorage
        localStorage.setItem('robot_manage_token', response.token)
        
        return response
      } catch (error: any) {
        this.logout()
        throw error
      }
    },

    async fetchUserInfo() {
      try {
        const user = await authApi.getCurrentUser()
        this.user = user
        this.isAuthenticated = true
        return user
      } catch (error) {
        this.logout()
        throw error
      }
    },

    async logout() {
      try {
        await authApi.logout()
      } catch (error) {
        // 即使API调用失败，也要清除本地状态
        console.error('Logout API error:', error)
      } finally {
        this.token = null
        this.user = null
        this.isAuthenticated = false
        localStorage.removeItem('robot_manage_token')
      }
    },

    checkAuth() {
      const token = localStorage.getItem('robot_manage_token')
      if (token) {
        this.token = token
        this.isAuthenticated = true
        // 可以在这里验证token是否有效
        return true
      }
      return false
    }
  }
})

