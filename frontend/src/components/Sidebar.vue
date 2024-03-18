<script setup lang="ts">
import { defineProps, defineEmits, inject, Ref } from 'vue'
import { cn, isUserLoggedIn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar'
import Logo from '@/components/Logo.vue'
import { Device } from '@/lib/types'
import DeviceTable from '@/components/DeviceTable.vue'

defineProps<{
  selectedDevice: Device | null,
  devices: Device[],
}>();

const emit = defineEmits(['update:selectedDevice', 'update:devices']);

const username = inject('username') as Ref<string | null>;
</script>

<template>
  <div :class="cn('pb-12', $attrs.class ?? '')" >
    <div class="space-y-4 py-4 flex flex-col h-full">
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
                        Map View
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
                        Statistics View
                    </Button>
                </router-link>
            </div>
        </div>
        <h2 class="relative px-7 text-lg font-semibold tracking-tight">
            Devices
        </h2>
        <div class="py-2 overflow-auto">
            <DeviceTable
                :devices="devices"
                :selectedDevice="selectedDevice"
                @update:selectedDevice="selectedDevice => emit('update:selectedDevice', selectedDevice)"
                @update:devices="devices => emit('update:devices', devices)"
            />
        </div>
        <div class="absolute bottom-4 left-4 flex items-center gap-2">
            <Avatar class="flex items-center cursor-pointer" v-if="isUserLoggedIn()">
                <AvatarImage :src="`https://avatars.jakerunzer.com/${username}.png`" alt="@radix-vue" />
                <AvatarFallback>{{username ? username[0] : ''}}</AvatarFallback>
            </Avatar>
            <div class="font-bold cursor-pointer">{{username}}</div>
        </div>
    </div>
  </div>
</template>