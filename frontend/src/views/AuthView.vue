<script setup lang="ts">
import UserAuthForm from '@/components/UserAuthForm.vue'
</script>

<template>
  <div class="container relative h-[800px] flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-2 lg:px-0">
    <div class="relative hidden h-full flex-col bg-background p-10 text-white dark:border-r lg:flex">
      <!-- <div class="absolute inset-0 bg-zinc-900" /> -->
      <div class="absolute inset-0 flex justify-center items-center">
          <img src="/touristmap.svg" class="transform scale-75" />
      </div>
      <div class="relative z-20 mt-auto">
        <!-- <blockquote class="space-y-2">
          <p class="text-lg">
            &ldquo;This library has saved me countless hours of work and
            helped me deliver stunning designs to my clients faster than
            ever before.&rdquo;
          </p>
          <footer class="text-sm">
            Sofia Davis
          </footer>
        </blockquote> -->
      </div>
    </div>
    <div class="lg:p-8">
      <div class="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
        <div class="flex flex-col space-y-2 text-center">
          <h1 class="text-2xl font-semibold tracking-tight">
            {{ $route.path === '/signup' ? 'Create an account' : 'Log in to your account' }}
          </h1>
          <p class="text-sm text-muted-foreground">
            Enter your email below to {{ $route.path === '/signup' ? 'create an account' : 'log in to your account' }}
          </p>
        </div>
        <UserAuthForm />
        <p class="px-8 text-center text-sm text-muted-foreground" v-if="$route.path === '/signup'" >
          By clicking continue, you agree to our
          <a
            href="/terms"
            class="underline underline-offset-4 hover:text-primary"
          >
            Terms of Service
          </a>
          and
          <a
            href="/privacy"
            class="underline underline-offset-4 hover:text-primary"
          >
            Privacy Policy
          </a>
          .
        </p>
      </div>
    </div>
  </div>
</template>