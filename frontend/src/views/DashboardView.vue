<script setup lang="ts">
import { onMounted, ref } from 'vue';
import Sidebar from '@/components/Sidebar.vue'
import { Device } from '@/lib/types'

const selectedDevice = ref<Device | null>(null);
const devices = ref([]);

onMounted(async () => {
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/device-locations`);
    devices.value = await response.json();
});
</script>

<template>
    <!-- <Navbar variant="dashboard" /> -->
    <div class="">
        <div class="bg-background">
            <div class="grid lg:grid-cols-5">
                <Sidebar
                    :selectedDevice="selectedDevice"
                    :devices="devices"
                    @update:selectedDevice="selectedDevice = $event"
                    class="sidebar hidden lg:block h-[100vh]"
                />
                <router-view
                    :selectedDevice="selectedDevice"
                    :devices="devices"
                    @update:selectedDevice="selectedDevice = $event"
                    class="col-span-3 lg:col-span-4 lg:border-l"
                />
            </div>
        </div>
    </div>
</template>
