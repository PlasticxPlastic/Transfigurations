<template>
  <nav class="bg-gradient-to-r from-blue-600 to-blue-800 shadow-lg">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <!-- Logo/Brand -->
        <div class="flex-shrink-0">
          <h1 class="text-white text-xl font-bold">Transfiguration</h1>
        </div>

        <!-- Navigation Links -->
        <div class="hidden md:block">
          <div class="ml-10 flex items-baseline space-x-4">
            <router-link 
              to="/" 
              class="text-white hover:bg-blue-700 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200"
              active-class="bg-blue-700"
            >
              Home
            </router-link>
            <router-link 
              to="/wand-power" 
              class="text-white hover:bg-blue-700 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200"
              active-class="bg-blue-700"
            >
              Wand Power
            </router-link>
            <router-link 
              to="/transfiguration-year1" 
              class="text-white hover:bg-blue-700 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200"
              active-class="bg-blue-700"
            >
              Transfiguration Year 1
            </router-link>
            <router-link 
              to="/transfiguration-year2" 
              class="text-white hover:bg-blue-700 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200"
              active-class="bg-blue-700"
            >
              Transfiguration Year 2
            </router-link>
            <button
              @click="handleLogout"
              class="bg-orange-500 hover:bg-orange-600 text-white px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200"
            >
              Logout
            </button>
          </div>
        </div>

        <!-- Mobile menu button -->
        <div class="md:hidden">
          <button 
            @click="isOpen = !isOpen"
            class="inline-flex items-center justify-center p-2 rounded-md text-white hover:bg-blue-700 focus:outline-none"
          >
            <svg 
              class="h-6 w-6" 
              :class="{'hidden': isOpen, 'block': !isOpen}" 
              xmlns="http://www.w3.org/2000/svg" 
              fill="none" 
              viewBox="0 0 24 24" 
              stroke="currentColor"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
            <svg 
              class="h-6 w-6" 
              :class="{'block': isOpen, 'hidden': !isOpen}" 
              xmlns="http://www.w3.org/2000/svg" 
              fill="none" 
              viewBox="0 0 24 24" 
              stroke="currentColor"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile menu -->
    <div 
      class="md:hidden" 
      :class="{'block': isOpen, 'hidden': !isOpen}"
    >
      <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
        <router-link 
          to="/" 
          class="text-white hover:bg-blue-700 block px-3 py-2 rounded-md text-base font-medium transition-colors duration-200"
          active-class="bg-blue-700"
        >
          Home
        </router-link>
        <router-link 
          to="/wand-power" 
          class="text-white hover:bg-blue-700 block px-3 py-2 rounded-md text-base font-medium transition-colors duration-200"
          active-class="bg-blue-700"
        >
          Wand Power
        </router-link>
        <router-link 
          to="/transfiguration-year1" 
          class="text-white hover:bg-blue-700 block px-3 py-2 rounded-md text-base font-medium transition-colors duration-200"
          active-class="bg-blue-700"
        >
          Transfiguration Year 1
        </router-link>
        <router-link 
          to="/transfiguration-year2" 
          class="text-white hover:bg-blue-700 block px-3 py-2 rounded-md text-base font-medium transition-colors duration-200"
          active-class="bg-blue-700"
        >
          Transfiguration Year 2
        </router-link>
        <button
          @click="handleLogout"
          class="bg-orange-500 hover:bg-orange-600 text-white block w-full text-left px-3 py-2 rounded-md text-base font-medium transition-colors duration-200"
        >
          Logout
        </button>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isOpen = ref(false)

const handleLogout = () => {
  localStorage.removeItem('isLoggedIn')
  router.push('/login')
}
</script> 