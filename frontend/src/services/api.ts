import axios from 'axios'

export interface User {
  id: number
  name: string
  email: string
  created_at: string
  updated_at: string
}

export interface ApiResponse<T> {
  success: boolean
  data: T
}

export interface ApiError {
  success: false
  error: {
    error: string
    message: string
    code: number
  }
}

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json'
  }
})

class UserService {
  async getUsers(): Promise<User[]> {
    try {
      const response = await api.get<ApiResponse<User[]>>('/users')
      if (!response.data.success) {
        throw new Error('Failed to fetch users')
      }
      return response.data.data
    } catch (error: any) {
      console.error('Error fetching users:', error)
      if (error.response?.data?.error?.message) {
        throw new Error(error.response.data.error.message)
      }
      throw error
    }
  }

  async getUserById(id: number): Promise<User> {
    try {
      const response = await api.get<ApiResponse<User>>(`/users/${id}`)
      if (!response.data.success) {
        throw new Error('Failed to fetch user')
      }
      return response.data.data
    } catch (error: any) {
      console.error(`Error fetching user ${id}:`, error)
      if (error.response?.data?.error?.message) {
        throw new Error(error.response.data.error.message)
      }
      throw error
    }
  }

  async createUser(user: Omit<User, 'id' | 'created_at' | 'updated_at'>): Promise<User> {
    try {
      const response = await api.post<ApiResponse<User>>('/users', user)
      if (!response.data.success) {
        throw new Error('Failed to create user')
      }
      return response.data.data
    } catch (error: any) {
      console.error('Error creating user:', error)
      if (error.response?.data?.error?.message) {
        throw new Error(error.response.data.error.message)
      }
      if (error.response?.status === 409) {
        throw new Error('A user with this email already exists')
      }
      throw error
    }
  }

  async updateUser(id: number, user: Omit<User, 'id' | 'created_at' | 'updated_at'>): Promise<User> {
    try {
      const response = await api.put<ApiResponse<User>>(`/users/${id}`, user)
      if (!response.data.success) {
        throw new Error('Failed to update user')
      }
      return response.data.data
    } catch (error: any) {
      console.error(`Error updating user ${id}:`, error)
      if (error.response?.data?.error?.message) {
        throw new Error(error.response.data.error.message)
      }
      if (error.response?.status === 409) {
        throw new Error('A user with this email already exists')
      } else if (error.response?.status === 404) {
        throw new Error('User not found')
      }
      throw error
    }
  }

  async deleteUser(id: number): Promise<void> {
    try {
      await api.delete(`/users/${id}`)
    } catch (error: any) {
      console.error(`Error deleting user ${id}:`, error)
      if (error.response?.data?.error?.message) {
        throw new Error(error.response.data.error.message)
      }
      if (error.response?.status === 404) {
        throw new Error('User not found')
      }
      throw error
    }
  }
}

export default new UserService()
