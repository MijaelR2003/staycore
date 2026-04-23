import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '@/views/admin/Dashboard.vue'
import Guests from '@/views/admin/Guests.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard
    },
    {
      path: '/guests',
      name: 'guests',
      component: Guests
    }
  ]
})

export default router