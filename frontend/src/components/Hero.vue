<script setup lang="ts">
</script>

<template>
<section class="bg-background">
    <div class="py-8 px-4 mx-auto max-w-screen-xl text-center lg:py-16 z-10 relative">
        <a href="https://github.com/rudyorre/trackerGPS" target="_blank" class="inline-flex justify-between items-center py-1 px-1 pe-4 mb-7 text-sm rounded-full bg-blue-900 text-blue-300 hover:bg-blue-800 hover:text-foreground">
            <span class="text-xs bg-blue-600 rounded-full text-white px-4 py-1.5 me-3">New</span> <span class="text-sm font-medium">Source code on Github! See what's new</span> 
            <svg class="w-2.5 h-2.5 ms-2 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
            </svg>
        </a>
        <h1 class="mb-4 text-4xl font-extrabold tracking-tight leading-none md:text-5xl lg:text-6xl text-white">
            We keep your business moving.
        </h1>
        <p class="mb-8 text-lg font-normal lg:text-xl sm:px-16 lg:px-48 text-gray-200">
            Our powerful GPS tracking solution gives you real-time insights into your fleet's activity, enabling you to optimize routes, improve driver behavior, and ensure the safety of your vehicles.
        </p>

    </div>
    <div class="grid md:grid-cols-3 gap-8 px-40">
        <div class="bg-gray-800 border border-gray-700 rounded-lg p-8 md:p-12 z-10">
            <a href="#" class="text-xs font-medium inline-flex items-center px-2.5 py-0.5 rounded-md bg-gray-700 text-green-400 mb-2">
                <svg class="w-2.5 h-2.5 me-1.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 16">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4 1 8l4 4m10-8 4 4-4 4M11 1 9 15"/>
                </svg>
                Frontend
            </a>
            <h2 class="text-white text-3xl font-extrabold mb-2">Vue3 + Vite</h2>
            <p class="text-lg font-normal text-gray-400 mb-4">A single page application (SPA) that utilizes the Google Maps API for the dashboard view. Allows signed in users to customize several properties about the dashboard.</p>
            <a href="https://github.com/rudyorre/TwoStepGPS/tree/main/frontend" target="_blank" class="text-blue-500 hover:underline font-medium text-lg inline-flex items-center">Frontend Source Code
                <svg class="w-3.5 h-3.5 ms-2 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/>
                </svg>
            </a>
        </div>
        <div class="bg-gray-800 border border-gray-700 rounded-lg p-8 md:p-12 z-10">
            <a href="#" class="text-xs font-medium inline-flex items-center px-2.5 py-0.5 rounded-md bg-gray-700 text-purple-400 mb-2">
                <svg class="w-2.5 h-2.5 me-1.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 16">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4 1 8l4 4m10-8 4 4-4 4M11 1 9 15"/>
                </svg>
                Backend
            </a>
            <h2 class="text-white text-3xl font-extrabold mb-2">Go + PostgresDB</h2>
            <p class="text-lg font-normal text-gray-400 mb-4">Backend server handles JWT authentication, integrates with the OneStepGPS API, and interfaces with the Postgres database.</p>
            <a href="https://github.com/rudyorre/TwoStepGPS/tree/main/backend" target="_blank" class="text-blue-500 hover:underline font-medium text-lg inline-flex items-center">Backend Source Code
                <svg class="w-3.5 h-3.5 ms-2 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/>
                </svg>
            </a>
        </div>
        <div class="bg-gray-800 border border-gray-700 rounded-lg p-8 md:p-12 z-10">
            <a href="#" class="text-xs font-medium inline-flex items-center px-2.5 py-0.5 rounded-md bg-gray-700 text-yellow-400 mb-2">
                <svg class="w-2.5 h-2.5 me-1.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 16">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4 1 8l4 4m10-8 4 4-4 4M11 1 9 15"/>
                </svg>
                Infrastructure
            </a>
            <h2 class="text-white text-3xl font-extrabold mb-2">Vercel + Fly.io</h2>
            <p class="text-lg font-normal text-gray-400 mb-4">The frontend is hosted as Vue/Vite app on Vercel, while the backend Go server is on a serverless fly.io VM which mounts the Postgres volume on startup.</p>
            <a href="https://github.com/rudyorre/TwoStepGPS" target="_blank" class="text-blue-500 hover:underline font-medium text-lg inline-flex items-center">Github Repo
                <svg class="w-3.5 h-3.5 ms-2 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5h12m0 0L9 1m4 4L9 9"/>
                </svg>
            </a>
        </div>
    </div>
    <!-- <div class="bg-gradient-to-b to-transparent from-muted w-full h-full absolute top-0 left-0"></div> -->
</section>
</template>

<style scoped>
/* h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
} */
</style>
