import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import WandPower from '../views/WandPower.vue'
import TransfigurationYear1 from '../views/TransfigurationYear1.vue'
import TransfigurationYear2 from '../views/TransfigurationYear2.vue'
import Login from '../views/Login.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { requiresAuth: true }
  },
  {
    path: '/wand-power',
    name: 'WandPower',
    component: WandPower,
    meta: { requiresAuth: true }
  },
  {
    path: '/transfiguration-year1',
    name: 'TransfigurationYear1',
    component: TransfigurationYear1,
    meta: { requiresAuth: true }
  },
  {
    path: '/transfiguration-year2',
    name: 'TransfigurationYear2',
    component: TransfigurationYear2,
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true'
  
  if (to.meta.requiresAuth && !isLoggedIn) {
    next('/login')
  } else if (to.path === '/login' && isLoggedIn) {
    next('/')
  } else {
    next()
  }
})

export default router 