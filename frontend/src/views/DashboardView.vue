<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue';
import Sidebar from '@/components/Sidebar.vue'
import { Device } from '@/lib/types'
import { isUserLoggedIn } from '@/lib/utils'
import Cookies from 'js-cookie'

const selectedDevice = ref<Device | null>(null);
const devices = ref<Device[]>([]);

let intervalId: number | null = null;
const pollingRateSeconds = 30;

const fetchDevices = async () => {
    if (isUserLoggedIn()) {
        const token = Cookies.get('token');
        const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/get-device-settings`, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });
        devices.value = await response.json();
    } else {
        const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/device-locations`);
        devices.value = await response.json();
    }
};

onMounted(async () => {
    fetchDevices();

    // Start polling when the component is mounted
    intervalId = window.setInterval(fetchDevices, pollingRateSeconds * 1000);

    // Pause polling when the user is not looking at the website
    document.addEventListener('visibilitychange', () => {
    if (document.visibilityState === 'hidden' && intervalId !== null) {
        window.clearInterval(intervalId);
        intervalId = null;
    } else {
        intervalId = window.setInterval(fetchDevices, pollingRateSeconds * 1000);
    }
    });
});

onUnmounted(() => {
  // Stop polling when the component is unmounted
  if (intervalId !== null) {
    window.clearInterval(intervalId);
  }
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
                    @update:devices="devices = $event"
                    class="sidebar lg:block h-[100vh]"
                />
                <router-view
                    :selectedDevice="selectedDevice"
                    :devices="devices"
                    @update:selectedDevice="selectedDevice = $event"
                    @update:devices="devices = $event"
                    class="col-span-3 lg:col-span-4 lg:border-l"
                />
            </div>
        </div>
    </div>
</template>
