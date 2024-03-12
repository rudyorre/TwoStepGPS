<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'
import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { ScrollArea } from '@/components/ui/scroll-area'
import Logo from '@/components/Logo.vue'
import { Device } from '@/lib/types'

defineProps<{
  selectedDevice: Device | null,
  devices: Device[],
}>();

const emit = defineEmits(['update:selectedDevice']);
const selectDevice = (device: Device) => {
    emit('update:selectedDevice', device);
};
</script>

<template>
  <div :class="cn('pb-12', $attrs.class ?? '')">
    <div class="space-y-4 py-4">
        <div class="pl-4 pt-1">
            <Logo />
        </div>
        <div class="px-3 py-2">
            <h2 class="mt-4 mb-2 px-4 text-lg font-semibold tracking-tight">
                Discover
            </h2>
            <div class="space-y-1">
                <router-link to="/dashboard">
                    <Button variant="secondary" class="w-full justify-start">
                        <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        class="mr-2 h-4 w-4"
                        >
                        <circle cx="12" cy="12" r="10" />
                        <polygon points="10 8 16 12 10 16 10 8" />
                        </svg>
                        Listen Now
                    </Button>
                </router-link>
                <router-link to="/dashboard/statistics">
                    <Button variant="ghost" class="w-full justify-start">
                        <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        class="mr-2 h-4 w-4"
                        >
                        <rect width="7" height="7" x="3" y="3" rx="1" />
                        <rect width="7" height="7" x="14" y="3" rx="1" />
                        <rect width="7" height="7" x="14" y="14" rx="1" />
                        <rect width="7" height="7" x="3" y="14" rx="1" />
                        </svg>
                        Browse
                    </Button>
                </router-link>
                <Button variant="ghost" class="w-full justify-start">
                    <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    class="mr-2 h-4 w-4"
                    >
                    <path d="M4.9 19.1C1 15.2 1 8.8 4.9 4.9" />
                    <path d="M7.8 16.2c-2.3-2.3-2.3-6.1 0-8.5" />
                    <circle cx="12" cy="12" r="2" />
                    <path d="M16.2 7.8c2.3 2.3 2.3 6.1 0 8.5" />
                    <path d="M19.1 4.9C23 8.8 23 15.1 19.1 19" />
                    </svg>
                    Radio
                </Button>
            </div>
        </div>
        <div class="py-2">
            <h2 class="relative px-7 text-lg font-semibold tracking-tight">
                Devices
            </h2>
            <ScrollArea class="h-[300px] px-1">
            <div class="space-y-1 p-2">
                <Button
                v-for="(device, i) in devices"
                :key="`${device}-${i}`"
                :variant="selectedDevice?.device_id === device.device_id ? 'secondary' : 'ghost'"
                class="w-full justify-start font-normal"
                @click="selectDevice(device)"
                >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    class="mr-2 h-4 w-4"
                >
                    <path d="M21 15V6" />
                    <path d="M18.5 18a2.5 2.5 0 1 0 0-5 2.5 2.5 0 0 0 0 5Z" />
                    <path d="M12 12H3" />
                    <path d="M16 6H3" />
                    <path d="M12 18H3" />
                </svg>
                {{ device.display_name }}
                </Button>
            </div>
            </ScrollArea>
        </div>
    </div>
  </div>
</template>