<script setup lang="ts">
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  navigationMenuTriggerStyle,
} from '@/components/ui/navigation-menu'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { LogOut } from 'lucide-vue-next'

import { ref, computed } from 'vue'
import Cookies from 'js-cookie'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'

import { useUserStore } from '@/lib/store'

const userStore = useUserStore();
const variant = ref('default');
const navbarClasses = computed(() => {
    return variant.value === 'default' ? 'backdrop-blur-xl' : '';
});

const logout = () => {
    userStore.resetState();
    Cookies.remove('token');
};

</script>

<template>


<nav :class="['sticky top-0 bg-background/20 z-20', navbarClasses]">
  <div class="flex flex-wrap items-center justify-between p-4">
    <router-link to="/" class="flex items-center space-x-3 rtl:space-x-reverse">
        <img src="https://flowbite.com/docs/images/logo.svg" class="h-8" alt="Flowbite Logo" />
        <div class="text-0">
            <span class="text-2xl font-extrabold whitespace-nowrap text-white">TwoStep</span><span class="text-2xl font-extrabold whitespace-nowrap text-blue-500">GPS</span>
        </div>
    </router-link>
    <button data-collapse-toggle="navbar-default" type="button" class="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600" aria-controls="navbar-default" aria-expanded="false">
        <span class="sr-only">Open main menu</span>
        <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 1h15M1 7h15M1 13h15"/>
        </svg>
    </button>
    <div class="hidden w-full md:block md:w-auto" id="navbar-default">
        <NavigationMenu class="pl-32">
            <NavigationMenuList>
                <NavigationMenuItem>
                <NavigationMenuTrigger class="bg-foreground/0">Overview</NavigationMenuTrigger>
                <NavigationMenuContent>
                    <ul class="grid gap-3 p-6 md:w-[400px] lg:w-[500px] lg:grid-cols-[minmax(0,.75fr)_minmax(0,1fr)]">
                    <li class="row-span-3">
                        <NavigationMenuLink as-child>
                        <a
                            class="flex h-full w-full select-none flex-col justify-end rounded-md bg-gradient-to-b from-muted/50 to-muted p-6 no-underline outline-none focus:shadow-md"
                            href="/"
                        >
                            <img src="https://www.radix-vue.com/logo.svg" class="h-6 w-6">
                            <a class="mb-2 mt-4 text-lg font-medium" href="https://rudyorre.com" target="_blank">
                            rudyorre.com
                            </a>
                            <p class="text-sm leading-tight text-muted-foreground">
                            Check out my other work on my personal website.
                            </p>
                        </a>
                        </NavigationMenuLink>
                    </li>
                    <li class="row-span-3">
                        <NavigationMenuLink>
                            <a
                            class="flex h-full w-full select-none flex-col justify-end rounded-md bg-gradient-to-b from-muted/50 to-muted p-6 no-underline outline-none focus:shadow-md"
                            href="/"
                        >
                            <router-link class="mb-2 mt-4 text-lg font-medium" to="/">
                            Why are we called TwoStepGPS?
                            </router-link>
                            <p class="text-sm leading-tight text-muted-foreground">
                                The first step is checking out our website. The second step is realizing this is a demo and going to <a href="https://www.onestepgps.com/" target="_blank">OneStepGPS</a>.
                            </p>
                        </a>
                        </NavigationMenuLink>
                    </li>
                    </ul>
                </NavigationMenuContent>
                </NavigationMenuItem>
                <NavigationMenuItem>
                <NavigationMenuTrigger class="bg-foreground/0">Products</NavigationMenuTrigger>
                <NavigationMenuContent>
                    <ul class="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-1 lg:w-[500px] ">
                        Just filler content since I don't have any products to display. Check out the dashboard instead, you don't even have to sign in!
                    </ul>
                </NavigationMenuContent>
                </NavigationMenuItem>
                <NavigationMenuItem>
                    <router-link to="/dashboard" :class="navigationMenuTriggerStyle()">
                        Dashboard
                    </router-link>
                </NavigationMenuItem>
                <NavigationMenuItem :key="userStore.username ? userStore.username : ''" v-if="userStore.username === null">
                    <router-link
                        to="/login"
                        :class="navigationMenuTriggerStyle()"
                    >
                        Login
                    </router-link>
                </NavigationMenuItem>
                <DropdownMenu :key="(userStore.username ? userStore.username : '') + '1'" v-if="userStore.username">
                    <DropdownMenuTrigger as-child>
                        <NavigationMenuItem>
                            <Avatar
                                class="flex items-center cursor-pointer"
                                
                            >
                                <AvatarImage :src="`https://avatars.jakerunzer.com/${userStore.username}.png`" alt="@radix-vue" />
                                <AvatarFallback>{{userStore.username ? userStore.username[0] : ''}}</AvatarFallback>
                            </Avatar>
                        </NavigationMenuItem>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent class="w-56">
                            <DropdownMenuLabel>My Account: {{userStore.username}}</DropdownMenuLabel>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem class="cursor-pointer" @click="logout">
                                <LogOut class="mr-2 h-4 w-4" />
                                <span>Log out</span>
                            </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </NavigationMenuList>
        </NavigationMenu>
    </div>
  </div>
</nav>
</template>