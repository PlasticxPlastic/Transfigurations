<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 via-white to-orange-50 flex items-center justify-center p-4">
    <div class="w-full max-w-md">
      <!-- Logo or Title Section -->
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold bg-gradient-to-r from-blue-600 to-orange-500 bg-clip-text text-transparent">
          System Integration
        </h1>
        <p class="text-gray-600 mt-2">Welcome back! Please login to continue.</p>
      </div>

      <!-- Login Form Card -->
      <div class="bg-white rounded-2xl shadow-xl p-8 backdrop-blur-sm bg-opacity-90">
        <form @submit.prevent="handleLogin" class="space-y-6">
          <!-- Username Input -->
          <div class="space-y-2">
            <label for="username" class="block text-sm font-medium text-gray-700">
              Username
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd" />
                </svg>
              </div>
              <input
                type="text"
                id="username"
                v-model="username"
                class="block w-full pl-10 pr-3 py-3 border border-gray-300 rounded-xl shadow-sm placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-150 ease-in-out"
                placeholder="Enter your username"
                required
              />
            </div>
          </div>

          <!-- Password Input -->
          <div class="space-y-2">
            <label for="password" class="block text-sm font-medium text-gray-700">
              Password
            </label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
                </svg>
              </div>
              <input
                type="password"
                id="password"
                v-model="password"
                class="block w-full pl-10 pr-3 py-3 border border-gray-300 rounded-xl shadow-sm placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition duration-150 ease-in-out"
                placeholder="Enter your password"
                required
              />
            </div>
          </div>

          <!-- Login Button -->
          <button
            type="submit"
            class="w-full flex justify-center py-3 px-4 border border-transparent rounded-xl shadow-sm text-sm font-medium text-white bg-gradient-to-r from-blue-500 to-orange-500 hover:from-blue-600 hover:to-orange-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transform transition duration-150 ease-in-out hover:scale-[1.02]"
          >
            Sign In
          </button>
        </form>

        <!-- Error Message -->
        <div v-if="error" class="mt-4 text-center text-red-500 text-sm">
          {{ error }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const username = ref('')
const password = ref('')
const error = ref('')

const API_URL = import.meta.env.PROD 
  ? 'https://backend-4mm5dzcqn-ohms-projects-4b3e1e96.vercel.app/api'
  : 'http://localhost:8080/api'

const handleLogin = async () => {
  try {
    error.value = ''
    const response = await axios.post(`${API_URL}/login`, {
      username: username.value,
      password: password.value
    })
    
    if (response.data.message === 'Login successful') {
      localStorage.setItem('isLoggedIn', 'true')
      // Emit login success event
      emit('login-success')
      router.push('/')
    }
  } catch (err) {
    error.value = err.response?.data?.error || 'Login failed. Please try again.'
  }
}

const emit = defineEmits(['login-success'])
</script> 