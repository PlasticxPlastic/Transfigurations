<template>
  <div class="min-h-screen bg-gray-100">
    <Navbar v-if="isLoggedIn" />
    <main :class="{'container mx-auto py-6': isLoggedIn}">
      <router-view @login-success="handleLoginSuccess"></router-view>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import Navbar from './components/Navbar.vue'

const isLoggedIn = ref(false)
const route = useRoute()

const handleLoginSuccess = () => {
  isLoggedIn.value = true
}

const checkLoginStatus = () => {
  isLoggedIn.value = localStorage.getItem('isLoggedIn') === 'true'
}

onMounted(() => {
  checkLoginStatus()
})

// Watch for route changes to update login status
watch(() => route.path, () => {
  checkLoginStatus()
})
</script>

<style>
@import 'tailwindcss/base';
@import 'tailwindcss/components';
@import 'tailwindcss/utilities';

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style> 