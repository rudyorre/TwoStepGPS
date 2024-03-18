<script setup lang="ts">
import { ref, provide } from 'vue'
import Navbar from '@/components/Navbar.vue'
import Footer from '@/components/Footer.vue'
import { Toaster } from '@/components/ui/sonner'

const username = ref<string | null>(null);
provide('username', username);
</script>

<template>
  <body class="flex flex-col min-h-screen">
    <Navbar id="navbar" v-if="!($route.path.includes('/dashboard'))" />
    <main class="flex-grow">
      <router-view />
    </main>
    <Footer v-if="!($route.path.includes('/dashboard'))" />
  </body>
  <Toaster />
</template>

<style scoped>
</style>
