<template>
  <AdminLayout pageTitle="Reservas">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <Select
          v-model="statusFilter"
          :options="statusOptions"
          optionLabel="label"
          optionValue="value"
          placeholder="Todos los estados"
          class="w-52"
          @change="fetchBookings"
        />
      </div>
      <Button
        label="Nueva reserva"
        icon="pi pi-plus"
        @click="openDialog"
      />
    </div>

    <!-- Tabla -->
    <div class="bg-white rounded-xl shadow-sm">
      <DataTable
        :value="bookings"
        :loading="loading"
        paginator
        :rows="10"
        stripedRows
      >
        <Column header="Huésped">
          <template #body="{ data }">
            <span class="font-medium">{{ getGuestName(data.guest_id) }}</span>
          </template>
        </Column>
        <Column header="Habitación">
          <template #body="{ data }">
            {{ data.room_id }}
          </template>
        </Column>
        <Column header="Check-in">
          <template #body="{ data }">
            {{ formatDate(data.check_in_date) }}
          </template>
        </Column>
        <Column header="Check-out">
          <template #body="{ data }">
            {{ formatDate(data.check_out_date) }}
          </template>
        </Column>
        <Column header="Estado">
          <template #body="{ data }">
            <Tag :value="data.status" :severity="getStatusSeverity(data.status)" />
          </template>
        </Column>
        <Column field="source" header="Origen" />
        <Column header="Acciones">
          <template #body="{ data }">
            <div class="flex gap-2">
              <Button
                v-if="data.status === 'pending' || data.status === 'confirmed'"
                label="Check-in"
                icon="pi pi-sign-in"
                size="small"
                severity="success"
                @click="doCheckIn(data.id)"
              />
              <Button
                v-if="data.status === 'checked_in'"
                label="Check-out"
                icon="pi pi-sign-out"
                size="small"
                severity="warning"
                @click="doCheckOut(data.id)"
              />
              <Button
                v-if="data.status === 'pending'"
                icon="pi pi-times"
                size="small"
                severity="danger"
                text
                @click="cancelBooking(data.id)"
              />
            </div>
          </template>
        </Column>
      </DataTable>
    </div>

    <!-- Dialog nueva reserva -->
    <Dialog
      v-model:visible="dialogVisible"
      header="Nueva reserva"
      :style="{ width: '560px' }"
      modal
    >
      <div class="grid grid-cols-2 gap-4 mt-2">
        <div class="flex flex-col gap-1 col-span-2">
          <label class="text-sm font-medium text-gray-600">Huésped *</label>
          <Select
            v-model="form.guest_id"
            :options="guests"
            optionLabel="full_name"
            optionValue="id"
            placeholder="Seleccionar huésped"
            filter
          />
        </div>
        <div class="flex flex-col gap-1 col-span-2">
          <label class="text-sm font-medium text-gray-600">Habitación *</label>
          <Select
            v-model="form.room_id"
            :options="rooms"
            optionLabel="label"
            optionValue="id"
            placeholder="Seleccionar habitación"
          />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Fecha check-in *</label>
          <DatePicker v-model="form.check_in_date" dateFormat="dd/mm/yy" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Fecha check-out *</label>
          <DatePicker v-model="form.check_out_date" dateFormat="dd/mm/yy" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Huéspedes</label>
          <InputNumber v-model="form.guest_count" :min="1" :max="10" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Origen</label>
          <Select
            v-model="form.source"
            :options="sourceOptions"
            optionLabel="label"
            optionValue="value"
            placeholder="Seleccionar"
          />
        </div>
        <div class="flex flex-col gap-1 col-span-2">
          <label class="text-sm font-medium text-gray-600">Notas</label>
          <Textarea v-model="form.notes" rows="2" placeholder="Notas adicionales..." />
        </div>
      </div>

      <template #footer>
        <Button label="Cancelar" severity="secondary" text @click="closeDialog" />
        <Button label="Guardar" icon="pi pi-check" :loading="saving" @click="saveBooking" />
      </template>
    </Dialog>
  </AdminLayout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import AdminLayout from '@/components/layout/AdminLayout.vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import Select from 'primevue/select'
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'
import DatePicker from 'primevue/datepicker'
import Textarea from 'primevue/textarea'
import Tag from 'primevue/tag'
import api from '@/services/api'

interface Guest {
  id: string
  first_name: string
  last_name: string
  full_name?: string
}

interface Room {
  id: string
  label: string
}

interface Booking {
  id: string
  guest_id: string
  room_id: string
  status: string
  check_in_date: string
  check_out_date: string
  source: string
}

const bookings = ref<Booking[]>([])
const guests = ref<Guest[]>([])
const rooms = ref<Room[]>([])
const loading = ref(false)
const saving = ref(false)
const dialogVisible = ref(false)
const statusFilter = ref(null)

const form = ref({
  guest_id: '',
  room_id: '',
  check_in_date: null,
  check_out_date: null,
  guest_count: 1,
  source: 'direct',
  notes: ''
})

const statusOptions = [
  { label: 'Todos', value: null },
  { label: 'Pendiente', value: 'pending' },
  { label: 'Confirmado', value: 'confirmed' },
  { label: 'Check-in', value: 'checked_in' },
  { label: 'Check-out', value: 'checked_out' },
  { label: 'Cancelado', value: 'cancelled' },
]

const sourceOptions = [
  { label: 'Directo', value: 'direct' },
  { label: 'Booking.com', value: 'booking' },
  { label: 'Airbnb', value: 'airbnb' },
  { label: 'Expedia', value: 'expedia' },
]

const getStatusSeverity = (status: string) => {
  const map: Record<string, string> = {
    pending: 'warn',
    confirmed: 'info',
    checked_in: 'success',
    checked_out: 'secondary',
    cancelled: 'danger',
  }
  return map[status] || 'info'
}

const formatDate = (date: string) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('es-ES')
}

const getGuestName = (guestId: string) => {
  const guest = guests.value.find((g: any) => g.id === guestId)
  return guest ? `${guest.first_name} ${guest.last_name}` : guestId
}

const fetchBookings = async () => {
  loading.value = true
  const res = await api.get('/bookings')
  bookings.value = res.data.data
  loading.value = false
}

const fetchGuests = async () => {
  const res = await api.get('/guests')
  guests.value = res.data.data.map((g: any) => ({
    ...g,
    full_name: `${g.first_name} ${g.last_name}`
  }))
}

const fetchRooms = async () => {
  const propRes = await api.get('/properties')
  const properties = propRes.data.data
  const allRooms: any[] = []
  for (const prop of properties) {
    const roomRes = await api.get(`/properties/${prop.id}/rooms`)
    roomRes.data.data.forEach((r: any) => {
      allRooms.push({
        ...r,
        label: `${prop.name} — Hab. ${r.number} (${r.type})`
      })
    })
  }
  rooms.value = allRooms
}

const openDialog = async () => {
  form.value = {
    guest_id: '',
    room_id: '',
    check_in_date: null,
    check_out_date: null,
    guest_count: 1,
    source: 'direct',
    notes: ''
  }
  await Promise.all([fetchGuests(), fetchRooms()])
  dialogVisible.value = true
}

const closeDialog = () => {
  dialogVisible.value = false
}

const saveBooking = async () => {
  if (!form.value.guest_id || !form.value.room_id || !form.value.check_in_date || !form.value.check_out_date) return
  saving.value = true
  await api.post('/bookings', {
    ...form.value,
    check_in_date: new Date(form.value.check_in_date as Date).toISOString(),
    check_out_date: new Date(form.value.check_out_date as Date).toISOString(),
  })
  await fetchBookings()
  saving.value = false
  closeDialog()
}

const doCheckIn = async (id: string) => {
  await api.post(`/bookings/${id}/checkin`)
  await fetchBookings()
}

const doCheckOut = async (id: string) => {
  await api.post(`/bookings/${id}/checkout`)
  await fetchBookings()
}

const cancelBooking = async (id: string) => {
  await api.delete(`/bookings/${id}`)
  await fetchBookings()
}

onMounted(async () => {
  await Promise.all([fetchBookings(), fetchGuests()])
})
</script>