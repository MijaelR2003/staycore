<template>
  <AdminLayout pageTitle="Propiedades">
    <!-- Header -->
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-gray-500 text-sm">{{ properties.length }} propiedades registradas</h2>
      <Button label="Nueva propiedad" icon="pi pi-plus" @click="openDialog" />
    </div>

    <!-- Cards -->
    <div class="grid grid-cols-3 gap-6">
      <div
        v-for="prop in properties"
        :key="prop.id"
        class="bg-white rounded-xl shadow-sm p-6 flex flex-col gap-3"
      >
        <div class="flex items-start justify-between">
          <div>
            <h3 class="font-semibold text-gray-800 text-lg">{{ prop.name }}</h3>
            <p class="text-gray-400 text-sm mt-1">{{ prop.address }}</p>
          </div>
          <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
            <i class="pi pi-building text-blue-600" />
          </div>
        </div>
        <div class="flex flex-col gap-1">
          <p class="text-sm text-gray-500">
            <i class="pi pi-phone mr-2" />{{ prop.phone || 'Sin teléfono' }}
          </p>
          <p class="text-sm text-gray-500">
            <i class="pi pi-envelope mr-2" />{{ prop.email || 'Sin email' }}
          </p>
        </div>
        <div class="flex gap-2 mt-2">
          <Button
            label="Ver habitaciones"
            icon="pi pi-th-large"
            size="small"
            severity="secondary"
            @click="goToRooms(prop.id)"
          />
          <Button
            icon="pi pi-pencil"
            size="small"
            text
            @click="editProperty(prop)"
          />
        </div>
      </div>
    </div>

    <!-- Dialog -->
    <Dialog
      v-model:visible="dialogVisible"
      :header="editMode ? 'Editar propiedad' : 'Nueva propiedad'"
      :style="{ width: '520px' }"
      modal
    >
      <div class="flex flex-col gap-4 mt-2">
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Nombre *</label>
          <InputText v-model="form.name" placeholder="Hotel Central" />
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-sm font-medium text-gray-600">Dirección</label>
          <InputText v-model="form.address" placeholder="Av. 16 de Julio 123" />
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-gray-600">Teléfono</label>
            <InputText v-model="form.phone" placeholder="2123456" />
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-sm font-medium text-gray-600">Email</label>
            <InputText v-model="form.email" placeholder="hotel@email.com" />
          </div>
        </div>
      </div>

      <template #footer>
        <Button label="Cancelar" severity="secondary" text @click="closeDialog" />
        <Button label="Guardar" icon="pi pi-check" :loading="saving" @click="saveProperty" />
      </template>
    </Dialog>
  </AdminLayout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import AdminLayout from '@/components/layout/AdminLayout.vue'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import api from '@/services/api'

const router = useRouter()
const properties = ref([])
const loading = ref(false)
const saving = ref(false)
const dialogVisible = ref(false)
const editMode = ref(false)
const editId = ref('')

const form = ref({
  name: '',
  address: '',
  phone: '',
  email: '',
})

const fetchProperties = async () => {
  loading.value = true
  const res = await api.get('/properties')
  properties.value = res.data.data
  loading.value = false
}

const openDialog = () => {
  editMode.value = false
  form.value = { name: '', address: '', phone: '', email: '' }
  dialogVisible.value = true
}

const editProperty = (prop: any) => {
  editMode.value = true
  editId.value = prop.id
  form.value = {
    name: prop.name,
    address: prop.address,
    phone: prop.phone,
    email: prop.email,
  }
  dialogVisible.value = true
}

const closeDialog = () => {
  dialogVisible.value = false
}

const saveProperty = async () => {
  if (!form.value.name) return
  saving.value = true
  if (editMode.value) {
    await api.put(`/properties/${editId.value}`, form.value)
  } else {
    await api.post('/properties', form.value)
  }
  await fetchProperties()
  saving.value = false
  closeDialog()
}

const goToRooms = (propertyId: string) => {
  router.push(`/properties/${propertyId}/rooms`)
}

onMounted(fetchProperties)
</script>