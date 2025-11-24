import http from './http'
import type { LoginRequest, LoginResponse } from '../types/auth'

export const authApi = {
  /**
   * 用户登录
   */
  login(data: LoginRequest): Promise<LoginResponse> {
    return http.post('/auth/login', data)
  },

  /**
   * 获取当前用户信息
   */
  getCurrentUser(): Promise<any> {
    return http.get('/auth/me')
  },

  /**
   * 退出登录
   */
  logout(): Promise<void> {
    return http.post('/auth/logout')
  }
}

