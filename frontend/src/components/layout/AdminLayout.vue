<template>
  <div class="flex h-screen bg-gray-100">
    <!-- Sidebar -->
    <aside class="w-64 bg-blue-900 text-white flex flex-col">
      <!-- Logo -->
      <div class="p-6 border-b border-blue-800">
        <h1 class="text-2xl font-bold text-white">Staycore</h1>
        <p class="text-blue-300 text-xs mt-1">Panel de administración</p>
      </div>

      <!-- Nav -->
      <nav class="flex-1 p-4 space-y-1">
        <RouterLink
          v-for="item in navItems"
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-3 px-4 py-3 rounded-lg text-blue-200 hover:bg-blue-800 hover:text-white transition-colors"
          active-class="bg-blue-700 text-white"
        >
          <i :class="item.icon" />
          <span>{{ item.label }}</span>
        </RouterLink>
      </nav>

      <!-- Footer -->
      <div class="p-4 border-t border-blue-800">
        <p class="text-blue-400 text-xs text-center">v1.0.0</p>
      </div>
    </aside>

    <!-- Main -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- Topbar -->
      <header class="bg-white shadow-sm px-8 py-4 flex items-center justify-between">
        <h2 class="text-lg font-semibold text-gray-700">{{ pageTitle }}</h2>
        <div class="flex items-center gap-3">
          <span class="text-sm text-gray-500">Admin</span>
          <div class="w-8 h-8 bg-blue-600 rounded-full flex items-center justify-center">
            <span class="text-white text-sm font-bold">A</span>
          </div>
        </div>
      </header>

      <!-- Content -->
      <main class="flex-1 overflow-y-auto p-8">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'

defineProps<{
  pageTitle?: string
}>()

const navItems = [
  { path: '/', label: 'Dashboard', icon: 'pi pi-home' },
  { path: '/guests', label: 'Huéspedes', icon: 'pi pi-users' },
  { path: '/properties', label: 'Propiedades', icon: 'pi pi-building' },
  { path: '/rooms', label: 'Habitaciones', icon: 'pi pi-th-large' },
  { path: '/bookings', label: 'Reservas', icon: 'pi pi-calendar' },
]
</script>