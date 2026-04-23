import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '@/views/admin/Dashboard.vue'
import Guests from '@/views/admin/Guests.vue'
import Bookings from '@/views/admin/Bookings.vue'
import Properties from '@/views/admin/Properties.vue'
import Rooms from '@/views/admin/Rooms.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', name: 'dashboard', component: Dashboard },
    { path: '/guests', name: 'guests', component: Guests },
    { path: '/bookings', name: 'bookings', component: Bookings },
    { path: '/properties', name: 'properties', component: Properties },
    { path: '/properties/:id/rooms', name: 'rooms', component: Rooms },
  ]
})

export default router



