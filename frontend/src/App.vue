<script setup lang="ts">
import Navbar from '@/components/Navbar.vue'
import Footer from '@/components/Footer.vue'
</script>

<template>
  <body class="flex flex-col min-h-screen">
    <Navbar id="navbar" v-if="!($route.path.includes('/dashboard'))" />
    <main class="flex-grow">
      <router-view />
    </main>
    <Footer v-if="!($route.path.includes('/dashboard'))" />
  </body>
</template>

<style scoped>
</style>
