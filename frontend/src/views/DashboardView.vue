<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import Sidebar from '@/components/Sidebar.vue'
import { useDeviceStore } from '@/lib/store'

const deviceStore = useDeviceStore();


onMounted(async () => {
    deviceStore.startPolling(30);
});

onUnmounted(() => {
  // Stop polling when the component is unmounted
//   if (intervalId !== null) {
//     window.clearInterval(intervalId);
//   }
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
