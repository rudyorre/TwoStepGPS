<script setup lang="ts">
import { onMounted } from 'vue'
import Navbar from '@/components/Navbar.vue'
import Footer from '@/components/Footer.vue'
import { Toaster } from '@/components/ui/sonner'
import Cookies from 'js-cookie'
import { useUserStore } from '@/lib/store'

const userStore = useUserStore();

onMounted(async () => {
    const token = Cookies.get('token');
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/profile`, {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    const data = await response.json();

    if (response.ok) {
      userStore.setUsername(data.username);
    }
});
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
