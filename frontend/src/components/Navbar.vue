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

const components: { title: string; href: string; description: string }[] = [
  {
    title: 'Alert Dialog',
    href: '/docs/primitives/alert-dialog',
    description:
      'A modal dialog that interrupts the user with important content and expects a response.',
  },
  {
    title: 'Hover Card',
    href: '/docs/primitives/hover-card',
    description:
      'For sighted users to preview content available behind a link.',
  },
  {
    title: 'Progress',
    href: '/docs/primitives/progress',
    description:
      'Displays an indicator showing the completion progress of a task, typically displayed as a progress bar.',
  },
  {
    title: 'Scroll-area',
    href: '/docs/primitives/scroll-area',
    description: 'Visually or semantically separates content.',
  },
  {
    title: 'Tabs',
    href: '/docs/primitives/tabs',
    description:
      'A set of layered sections of content—known as tab panels—that are displayed one at a time.',
  },
  {
    title: 'Tooltip',
    href: '/docs/primitives/tooltip',
    description:
      'A popup that displays information related to an element when the element receives keyboard focus or the mouse hovers over it.',
  },
]
import { ref, computed, onMounted, inject, Ref } from 'vue'
import Cookies from 'js-cookie'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'

const variant = ref('default');
const navbarClasses = computed(() => {
    return variant.value === 'default' ? 'backdrop-blur-xl' : '';
});
const username = inject('username') as Ref<string | null>;
const isLoading = ref(true);

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
        username.value = data.username;
    }
    isLoading.value = false;
});

const logout = () => {
    username.value = null;
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
                            <div class="mb-2 mt-4 text-lg font-medium">
                            shadcn/ui
                            </div>
                            <p class="text-sm leading-tight text-muted-foreground">
                            Beautifully designed components built with Radix UI and
                            Tailwind CSS.
                            </p>
                        </a>
                        </NavigationMenuLink>
                    </li>
                    <NavigationMenuLink>
                        <a
                        class="flex select-none flex-col justify-end rounded-md bg-gradient-to-b from-muted/50 to-muted p-6 no-underline outline-none focus:shadow-md"
                        href="/"
                        >
                        <!-- <img src="https://www.radix-vue.com/logo.svg" class="h-6 w-6"> -->
                        <!-- <div class="mb-2 mt-4 text-lg font-medium">
                            shadcn/ui
                        </div> -->
                        <p class="text-sm leading-tight text-muted-foreground">
                            Beautifully designed components built with Radix UI and
                            Tailwind CSS.
                        </p>
                        </a>
                    </NavigationMenuLink>
                    <NavigationMenuLink class="flex select-none flex-col justify-end rounded-md no-underline outline-none focus:shadow-md cursor-pointer">
                        Fleet Tracking
                    </NavigationMenuLink>
                    <ListItem href="/docs" title="Introduction">
                        
                    </ListItem>
                    <!-- <ListItem href="/docs/installation" title="Installation">
                        How to install dependencies and structure your app.
                    </ListItem>
                    <ListItem href="/docs/primitives/typography" title="Typography">
                        Styles for headings, paragraphs, lists...etc
                    </ListItem> -->
                    </ul>
                </NavigationMenuContent>
                </NavigationMenuItem>
                <NavigationMenuItem>
                <NavigationMenuTrigger class="bg-foreground/0">Products</NavigationMenuTrigger>
                <NavigationMenuContent>
                    <ul class="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[500px] ">
                    <ListItem
                        v-for="component in components"
                        :key="component.title"
                        :title="component.title"
                        :href="component.href"
                    >
                        {{ component.description }}
                    </ListItem>
                    </ul>
                </NavigationMenuContent>
                </NavigationMenuItem>
                <NavigationMenuItem>
                    <router-link to="/dashboard" :class="navigationMenuTriggerStyle()">
                        Dashboard
                    </router-link>
                </NavigationMenuItem>
                <NavigationMenuItem :key="username ? username : ''" v-if="username === null && !isLoading">
                    <router-link
                        to="/login"
                        :class="navigationMenuTriggerStyle()"
                    >
                        Login
                    </router-link>
                </NavigationMenuItem>
                <DropdownMenu :key="(username ? username : '') + '1'" v-if="username && !isLoading">
                    <DropdownMenuTrigger as-child>
                        <NavigationMenuItem>
                            <Avatar
                                class="flex items-center cursor-pointer"
                                
                            >
                                <AvatarImage :src="`https://avatars.jakerunzer.com/${username}.png`" alt="@radix-vue" />
                                <AvatarFallback>{{username ? username[0] : ''}}</AvatarFallback>
                            </Avatar>
                        </NavigationMenuItem>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent class="w-56">
                            <DropdownMenuLabel>My Account: {{username}}</DropdownMenuLabel>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem class="cursor-pointer" @click="logout">
                                <LogOut class="mr-2 h-4 w-4" />
                                <span>Log out</span>
                            </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
                <NavigationMenuItem v-if="isLoading">
                    <Avatar
                        class="flex items-center cursor-pointer"
                    >
                        <AvatarImage src="" alt="@radix-vue" />
                    </Avatar>
                </NavigationMenuItem>
            </NavigationMenuList>
        </NavigationMenu>
    </div>
  </div>
</nav>
</template>