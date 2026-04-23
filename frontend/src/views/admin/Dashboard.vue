<template>
  <AdminLayout pageTitle="Dashboard">
    <!-- Stats -->
    <div class="grid grid-cols-3 gap-6 mb-8">
      <div class="bg-white rounded-xl shadow-sm p-6">
        <div class="flex items-center justify-between mb-4">
          <span class="text-gray-500 text-sm font-medium uppercase">Huéspedes</span>
          <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
            <i class="pi pi-users text-blue-600" />
          </div>
        </div>
        <p class="text-3xl font-bold text-gray-800">{{ stats.guests }}</p>
        <p class="text-green-500 text-sm mt-1">Registrados</p>
      </div>

      <div class="bg-white rounded-xl shadow-sm p-6">
        <div class="flex items-center justify-between mb-4">
          <span class="text-gray-500 text-sm font-medium uppercase">Reservas</span>
          <div class="w-10 h-10 bg-green-100 rounded-lg flex items-center justify-center">
            <i class="pi pi-calendar text-green-600" />
          </div>
        </div>
        <p class="text-3xl font-bold text-gray-800">{{ stats.bookings }}</p>
        <p class="text-green-500 text-sm mt-1">Activas</p>
      </div>

      <div class="bg-white rounded-xl shadow-sm p-6">
        <div class="flex items-center justify-between mb-4">
          <span class="text-gray-500 text-sm font-medium uppercase">Habitaciones</span>
          <div class="w-10 h-10 bg-purple-100 rounded-lg flex items-center justify-center">
            <i class="pi pi-th-large text-purple-600" />
          </div>
        </div>
        <p class="text-3xl font-bold text-gray-800">{{ stats.rooms }}</p>
        <p class="text-gray-400 text-sm mt-1">Disponibles</p>
      </div>
    </div>

    <!-- Quick actions -->
    <div class="bg-white rounded-xl shadow-sm p-6">
      <h3 class="text-gray-700 font-semibold mb-4">Acciones rápidas</h3>
      <div class="flex gap-4">
        <RouterLink
          to="/guests/new"
          class="flex items-center gap-2 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors text-sm"
        >
          <i class="pi pi-plus" />
          Nuevo huésped
        </RouterLink>
        <RouterLink
          to="/bookings/new"
          class="flex items-center gap-2 bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors text-sm"
        >
          <i class="pi pi-plus" />
          Nueva reserva
        </RouterLink>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import AdminLayout from '@/components/layout/AdminLayout.vue'
import api from '@/services/api'

const stats = ref({
  guests: 0,
  bookings: 0,
  rooms: 0
})

onMounted(async () => {
  const [guestsRes, bookingsRes] = await Promise.all([
    api.get('/guests'),
    api.get('/bookings'),
  ])

  stats.value.guests = guestsRes.data.data.length
  stats.value.bookings = bookingsRes.data.data.length
})
</script>   