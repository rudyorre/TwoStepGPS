<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import Sidebar from '@/components/Sidebar.vue'
import { fetchDevices } from '@/lib/utils'

let intervalId: number | null = null;
const pollingRateSeconds = 30;

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
    <div class="">
        <div class="bg-background">
            <div class="grid lg:grid-cols-5">
                <Sidebar class="sidebar lg:block h-[100vh]" />
                <router-view class="col-span-3 lg:col-span-4 lg:border-l" />
            </div>
        </div>
    </div>
</template>
