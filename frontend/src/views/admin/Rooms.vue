<template>
  <AdminLayout pageTitle="Habitaciones">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <Button
          icon="pi pi-arrow-left"
          text
          @click="router.back()"
        />
        <h2 class="text-gray-600 font-medium">{{ propertyName }}</h2>
      </div>
      <Button label="Nueva habitación" icon="pi pi-plus" @click="openDialog" />
    </div>

    <!-- Grid habitaciones -->
    <div class="grid grid-cols-4 gap-4">
      <div
        v-for="room in rooms"
        :key="room.id"
        class="bg-white rounded-xl shadow-sm p-5 flex flex-col gap-3"
        :class="{
          'border-2 border-green-400': room.status === 'available',
          'border-2 border-red-400': room.status === 'occupied',
        }"
      >
        <div class="flex items-center justify-between">
          <span class="text-2xl font-bold text-gray-800">{{ room.number }}</span>
          <Tag
            :value="room.status === 'available' ? 'Disponible' : 'Ocupada'"
            :severity="room.status === 'available' ? 'success' : 'danger'"
          />
        </div>
        <div class="flex flex-col gap-1">
          <p class="text-sm text-gray-500 capitalize">{{ room.type }}</p>
          <p class="text-sm text-gray-500">Capacidad: {{ room.capacity }}</p>
          <p class="text-sm font-semibold text-blue-600">${{ room.price }}/noche</p>
        </div>
        <Button
          icon="pi pi-pencil"
          size="small"
          text
          severity="secondary"
          @click="editRoom(room)"
        />
      </div>
    </div>

    <!-- Dialog -->
    <Dialog
      v-model:visible="dialogVisible"
      :header="editMode ? 'Editar habitación' : 'Nueva habitación'"
      :style="{ width: '480px' }"
      modal
    >
      <div class="grid grid-cols-2 gap-4 mt-2">
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Número *</label>
          <InputText v-model="form.number" placeholder="101" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Tipo *</label>
          <Select
            v-model="form.type"
            :options="roomTypes"
            optionLabel="label"
            optionValue="value"
            placeholder="Seleccionar"
          />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Capacidad</label>
          <InputNumber v-model="form.capacity" :min="1" :max="10" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Piso</label>
          <InputNumber v-model="form.floor" :min="1" />
        </div>
        <div class="flex flex-col gap-1 col-span-2">
          <label class="text-sm font-medium text-gray-600">Precio por noche</label>
          <InputNumber v-model="form.price" mode="currency" currency="USD" />
        </div>
      </div>

      <template #footer>
        <Button label="Cancelar" severity="secondary" text @click="closeDialog" />
        <Button label="Guardar" icon="pi pi-check" :loading="saving" @click="saveRoom" />
      </template>
    </Dialog>
  </AdminLayout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import AdminLayout from '@/components/layout/AdminLayout.vue'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'
import Select from 'primevue/select'
import Tag from 'primevue/tag'
import api from '@/services/api'

const router = useRouter()
const route = useRoute()
const propertyId = route.params.id as string

const rooms = ref([])
const propertyName = ref('')
const loading = ref(false)
const saving = ref(false)
const dialogVisible = ref(false)
const editMode = ref(false)
const editId = ref('')

const form = ref({
  number: '',
  type: '',
  capacity: 1,
  floor: 1,
  price: 0,
})

const roomTypes = [
  { label: 'Simple', value: 'simple' },
  { label: 'Doble', value: 'double' },
  { label: 'Triple', value: 'triple' },
  { label: 'Suite', value: 'suite' },
]

const fetchRooms = async () => {
  loading.value = true
  const res = await api.get(`/properties/${propertyId}/rooms`)
  rooms.value = res.data.data
  loading.value = false
}

const fetchProperty = async () => {
  const res = await api.get(`/properties/${propertyId}`)
  propertyName.value = res.data.data.name
}

const openDialog = () => {
  editMode.value = false
  form.value = { number: '', type: '', capacity: 1, floor: 1, price: 0 }
  dialogVisible.value = true
}

const editRoom = (room: any) => {
  editMode.value = true
  editId.value = room.id
  form.value = {
    number: room.number,
    type: room.type,
    capacity: room.capacity,
    floor: room.floor,
    price: room.price,
  }
  dialogVisible.value = true
}

const closeDialog = () => {
  dialogVisible.value = false
}

const saveRoom = async () => {
  if (!form.value.number || !form.value.type) return
  saving.value = true
  if (editMode.value) {
    await api.put(`/properties/${propertyId}/rooms/${editId.value}`, form.value)
  } else {
    await api.post(`/properties/${propertyId}/rooms`, form.value)
  }
  await fetchRooms()
  saving.value = false
  closeDialog()
}

onMounted(async () => {
  await Promise.all([fetchRooms(), fetchProperty()])
})
</script>