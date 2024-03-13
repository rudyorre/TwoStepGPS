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
        <form class="w-full max-w-md mx-auto">   
            <!-- <label for="default-email" class="mb-2 text-sm font-medium sr-onlytext-white">Email sign-up</label> -->
            <div class="relative">
                <div class="absolute inset-y-0 rtl:inset-x-0 start-0 flex items-center ps-3.5 pointer-events-none">
                    <svg class="w-4 h-4 text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 16">
                        <path d="m10.036 8.278 9.258-7.79A1.979 1.979 0 0 0 18 0H2A1.987 1.987 0 0 0 .641.541l9.395 7.737Z"/>
                        <path d="M11.241 9.817c-.36.275-.801.425-1.255.427-.428 0-.845-.138-1.187-.395L0 2.6V14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V2.5l-8.759 7.317Z"/>
                    </svg>
                </div>
                <input type="email" id="default-email" class="block w-full p-4 ps-10 text-sm  border rounded-lg bg-gray-800 border-gray-700 placeholder-gray-400 text-white focus:ring-blue-500 focus:border-blue-500" placeholder="Enter your email here..." required />
                <button type="submit" class="text-white absolute end-2.5 bottom-2.5 focus:ring-4 focus:outline-none font-medium rounded-lg text-sm px-4 py-2 bg-blue-600 hover:bg-blue-700 focus:ring-blue-800">Sign up</button>
            </div>
        </form>
    </div>
    <div class="grid md:grid-cols-2 gap-8 px-40">
        <div class="bg-gray-800 border border-gray-700 rounded-lg p-8 md:p-12 z-10">
            <a href="#" class="text-xs font-medium inline-flex items-center px-2.5 py-0.5 rounded-md bg-gray-700 text-green-400 mb-2">
                <svg class="w-2.5 h-2.5 me-1.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 18">
                    <path d="M17 11h-2.722L8 17.278a5.512 5.512 0 0 1-.9.722H17a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1ZM6 0H1a1 1 0 0 0-1 1v13.5a3.5 3.5 0 1 0 7 0V1a1 1 0 0 0-1-1ZM3.5 15.5a1 1 0 1 1 0-2 1 1 0 0 1 0 2ZM16.132 4.9 12.6 1.368a1 1 0 0 0-1.414 0L9 3.55v9.9l7.132-7.132a1 1 0 0 0 0-1.418Z"/>
                </svg>
                Design
            </a>
            <h2 class="text-white text-3xl font-extrabold mb-2">Start with Flowbite Design System</h2>
            <p class="text-lg font-normal text-gray-400 mb-4">Static websites are now used to bootstrap lots of websites and are becoming the basis for a variety of tools that even influence both web designers and developers.</p>
            <a href="#" class="text-blue-500 hover:underline font-medium text-lg inline-flex items-center">Read more
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
                Code
            </a>
            <h2 class="text-white text-3xl font-extrabold mb-2">Best react libraries around the web</h2>
            <p class="text-lg font-normal text-gray-400 mb-4">Static websites are now used to bootstrap lots of websites and are becoming the basis for a variety of tools that even influence both web designers and developers.</p>
            <a href="#" class="text-blue-500 hover:underline font-medium text-lg inline-flex items-center">Read more
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
