<script setup lang="ts">
import { ref, Ref, inject } from 'vue'
import { useRouter } from 'vue-router'
import Cookies from 'js-cookie'
import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

const router = useRouter();
const username = inject('username') as Ref<string | null>;
const formUsername = ref('');
const formPassword = ref('');
const errorMessage = ref('');

const isLoading = ref(false)

async function onSubmit(event: Event) {
  event.preventDefault()
  isLoading.value = true;
  const endpoint = router.currentRoute.value.path === '/login' ? 'login' : 'signup';
  const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/${endpoint}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      username: formUsername.value,
      password: formPassword.value
    }),
  });
  if (response.ok) {
    username.value = formUsername.value;
    const json = await response.json();
    const token = json.token;
    Cookies.set('token', token, { expires: 14 });
    if (window.history.length > 1) {
      router.go(-1);
    } else {
      router.push('/');
    }
  } else {
    formUsername.value = '';
    formPassword.value = '';
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
            v-model="formUsername"
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
            v-model="formPassword"
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
          {{$route.path === '/login' ? "Sign In" : "Sign Up"}}
        </Button>
      </div>
    </form>
    <div class="relative">
      <div class="absolute inset-0 flex items-center">
        <span class="w-full border-t" />
      </div>
      <div class="relative flex justify-center text-xs uppercase">
        <span class="bg-background px-2 text-muted-foreground">
          {{$route.path === '/login' ? "Don't have an account?" : "Already have an account?"}}
        </span>
      </div>
    </div>
    <router-link
      :to="$route.path === '/signup' ? '/login' : '/signup'"
      v-slot="{ navigate }"
      class="inline-block"
    >
    <Button variant="outline" type="button" :disabled="isLoading" @click="navigate" class="w-full">
        <LucideSpinner v-if="isLoading" class="mr-2 h-4 w-4 animate-spin" />
        {{$route.path === '/login' ? "Sign Up" : "Sign In"}}
      </Button>
    </router-link>
  </div>
</template>