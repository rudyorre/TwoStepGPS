<script setup lang="ts">
// import Navbar from '@/components/Navbar.vue'
import Sidebar from '@/components/Sidebar.vue'
import { useRouter } from 'vue-router'
const router = useRouter()

const redirectReload = async () => {
  await router.push({ path: '/dashboard' })
  router.go(0)
}
</script>

<template>
    <!-- <Navbar variant="dashboard" /> -->
    <div class="">
        <div class="bg-background">
            <div class="grid lg:grid-cols-5">
                <Sidebar class="sidebar hidden lg:block h-[100vh]" />
                <router-view class="col-span-3 lg:col-span-4 lg:border-l" />
            </div>
        </div>
    </div>
</template>
