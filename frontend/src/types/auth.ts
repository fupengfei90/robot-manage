export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  user: {
    id: string
    username: string
    role?: string
  }
}

export interface UserInfo {
  id: string
  username: string
  role?: string
  email?: string
  avatar?: string
}

