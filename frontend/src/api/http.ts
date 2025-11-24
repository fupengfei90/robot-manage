import axios from 'axios'

const http = axios.create({
  baseURL: '/api/v1',
  timeout: 10000
})

http.interceptors.request.use((config) => {
  const token = localStorage.getItem('robot_manage_token')
  if (token) {
    config.headers = config.headers || {}
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

http.interceptors.response.use(
  (res) => {
    // 后端返回格式: {code: 0, message: "success", data: {...}}
    const response = res.data
    if (response && typeof response === 'object' && 'code' in response) {
      // 如果code不为0，说明有错误，抛出异常让错误回调处理
      if (response.code !== 0) {
        const error: any = new Error(response.message || '请求失败')
        error.response = {
          status: response.code >= 400 ? response.code : 400,
          data: response
        }
        throw error
      }
      // 返回data字段
      return response.data
    }
    // 如果不是标准格式，直接返回
    return response
  },
  (error) => {
    // 处理错误响应
    if (error.response) {
      const { data, status } = error.response
      // 如果后端返回了标准格式的错误 {code, message, data}
      if (data && typeof data === 'object' && 'message' in data) {
        const errorMessage = data.message || '请求失败'
        // 处理HTTP状态码错误
        if (status === 401 || (data.code && data.code === 401)) {
          // 清除token并跳转到登录页
          localStorage.removeItem('robot_manage_token')
          if (window.location.pathname !== '/login') {
            window.location.href = '/login'
          }
          return Promise.reject(new Error('登录已过期，请重新登录'))
        }
        return Promise.reject(new Error(errorMessage))
      }
      // 处理HTTP状态码错误（非标准格式）
      if (status === 401) {
        localStorage.removeItem('robot_manage_token')
        if (window.location.pathname !== '/login') {
          window.location.href = '/login'
        }
        return Promise.reject(new Error('登录已过期，请重新登录'))
      }
      return Promise.reject(new Error(data?.message || `请求失败: ${status}`))
    }
    return Promise.reject(error)
  }
)

export default http
