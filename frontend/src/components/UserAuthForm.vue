<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Cookies from 'js-cookie'
// import LucideSpinner from '~icons/lucide/loader-2'
// import GitHubLogo from '~icons/radix-icons/github-logo'

import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

const router = useRouter();
const username = ref('');
const password = ref('');
const errorMessage = ref('');

const isLoading = ref(false)
async function onSubmit(event: Event) {
  event.preventDefault()
  isLoading.value = true;
  const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      username: username.value,
      password: password.value
    }),
  });
  if (response.ok) {
    const token = await response.text();
    Cookies.set('token', token, { expires: 14 });
    if (window.history.length > 1) {
      router.go(-1);
    } else {
      router.push('/');
    }
  } else {
    username.value = '';
    password.value = '';
    errorMessage.value = 'Invalid username or password';
    isLoading.value = false;
  }
}
</script>

<template>
  <div :class="cn('grid gap-6', $attrs.class ?? '')">
    <form @submit="onSubmit">
      <div class="grid gap-2">
        <div class="grid gap-1">
          <Label class="sr-only" for="username">
            Username
          </Label>
          <Input
            id="username"
            v-model="username"
            placeholder="username"
            type="username"
            auto-capitalize="none"
            auto-complete="username"
            auto-correct="off"
            :disabled="isLoading"
          />
        </div>
        <div class="grid gap-1">
          <Label class="sr-only" for="password">
            Password
          </Label>
          <Input
            id="password"
            v-model="password"
            placeholder="password"
            type="password"
            auto-capitalize="none"
            auto-complete="password"
            auto-correct="off"
            :disabled="isLoading"
          />
        </div>
        <p v-if="errorMessage" class="text-red-500">{{ errorMessage }}</p>
        <Button :disabled="isLoading">
          <LucideSpinner v-if="isLoading" class="mr-2 h-4 w-4 animate-spin" />
          Sign In
        </Button>
      </div>
    </form>
    <div class="relative">
      <div class="absolute inset-0 flex items-center">
        <span class="w-full border-t" />
      </div>
      <div class="relative flex justify-center text-xs uppercase">
        <span class="bg-background px-2 text-muted-foreground">
          Or continue with
        </span>
      </div>
    </div>
    <Button variant="outline" type="button" :disabled="isLoading">
      <LucideSpinner v-if="isLoading" class="mr-2 h-4 w-4 animate-spin" />
      <GitHubLogo v-else class="mr-2 h-4 w-4" />
      GitHub
    </Button>
  </div>
</template>