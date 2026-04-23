<template>
  <AdminLayout pageTitle="Huéspedes">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <InputText
          v-model="search"
          placeholder="Buscar por nombre o documento..."
          class="w-80"
        />
      </div>
      <Button
        label="Nuevo huésped"
        icon="pi pi-plus"
        @click="openDialog"
      />
    </div>

    <!-- Tabla -->
    <div class="bg-white rounded-xl shadow-sm">
      <DataTable
        :value="guests"
        :loading="loading"
        paginator
        :rows="10"
        stripedRows
        class="p-datatable-sm"
      >
        <Column field="first_name" header="Nombre">
          <template #body="{ data }">
            {{ data.first_name }} {{ data.last_name }}
          </template>
        </Column>
        <Column field="document_type" header="Tipo doc." />
        <Column field="document_number" header="Documento" />
        <Column field="nationality" header="Nacionalidad" />
        <Column field="email" header="Email" />
        <Column field="phone" header="Teléfono" />
        <Column header="Acciones">
          <template #body="{ data }">
            <Button
              icon="pi pi-eye"
              severity="secondary"
              text
              @click="viewGuest(data)"
            />
          </template>
        </Column>
      </DataTable>
    </div>

    <!-- Dialog nuevo huésped -->
    <Dialog
      v-model:visible="dialogVisible"
      header="Nuevo huésped"
      :style="{ width: '600px' }"
      modal
    >
      <div class="grid grid-cols-2 gap-4 mt-2">
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Nombre *</label>
          <InputText v-model="form.first_name" placeholder="Juan" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Apellido *</label>
          <InputText v-model="form.last_name" placeholder="González" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Tipo de documento *</label>
          <Select
            v-model="form.document_type"
            :options="documentTypes"
            optionLabel="label"
            optionValue="value"
            placeholder="Seleccionar"
          />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Número de documento *</label>
          <InputText v-model="form.document_number" placeholder="12345678" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">País del documento</label>
          <InputText v-model="form.document_country" placeholder="BO" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Nacionalidad</label>
          <InputText v-model="form.nationality" placeholder="Boliviana" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Email</label>
          <InputText v-model="form.email" placeholder="juan@email.com" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Teléfono</label>
          <InputText v-model="form.phone" placeholder="77712345" />
        </div>
      </div>

      <template #footer>
        <Button label="Cancelar" severity="secondary" text @click="closeDialog" />
        <Button label="Guardar" icon="pi pi-check" :loading="saving" @click="saveGuest" />
      </template>
    </Dialog>
  </AdminLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import AdminLayout from '@/components/layout/AdminLayout.vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import Dialog from 'primevue/dialog'
import Select from 'primevue/select'
import api from '@/services/api'

const guests = ref([])
const loading = ref(false)
const saving = ref(false)
const dialogVisible = ref(false)
const search = ref('')

const form = ref({
  first_name: '',
  last_name: '',
  document_type: '',
  document_number: '',
  document_country: '',
  nationality: '',
  email: '',
  phone: ''
})

const documentTypes = [
  { label: 'DNI', value: 'dni' },
  { label: 'Pasaporte', value: 'passport' },
  { label: 'Cédula', value: 'cedula' },
  { label: 'Otro', value: 'other' },
]

const fetchGuests = async () => {
  loading.value = true
  const res = await api.get('/guests')
  guests.value = res.data.data
  loading.value = false
}

const openDialog = () => {
  form.value = {
    first_name: '',
    last_name: '',
    document_type: '',
    document_number: '',
    document_country: '',
    nationality: '',
    email: '',
    phone: ''
  }
  dialogVisible.value = true
}

const closeDialog = () => {
  dialogVisible.value = false
}

const saveGuest = async () => {
  if (!form.value.first_name || !form.value.last_name || !form.value.document_number) {
    return
  }
  saving.value = true
  await api.post('/guests', form.value)
  await fetchGuests()
  saving.value = false
  closeDialog()
}

const viewGuest = (guest: any) => {
  console.log(guest)
}

onMounted(fetchGuests)
</script>