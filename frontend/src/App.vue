<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
import { Button } from '@/components/ui/button'
import Spotlight from '@/components/Spotlight.vue'
import Header from '@/components/Header.vue'
import Hero from '@/components/Hero.vue'
</script>

<template>
  <div class="min-h-screen bg-background">
    <Header />
    <Hero />
    <div class="h-[1000px]"></div>
  </div>
  <!-- <header>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />
      <div className="font-bold text-5xl">test2 text</div>
    </div>
  </header>

  <main className="bg-red-500 min-h-screen min-w-screen">
    <TheWelcome />
    <Spotlight />
    <Button>Click me</Button>
    
  </main> -->
</template>

<style scoped>
  /* header {
    line-height: 1.5;
  }

  .logo {
    display: block;
    margin: 0 auto 2rem;
  }

  @media (min-width: 1024px) {
    header {
      display: flex;
      place-items: center;
      padding-right: calc(var(--section-gap) / 2);
    }

    .logo {
      margin: 0 2rem 0 0;
    }

    header .wrapper {
      display: flex;
      place-items: flex-start;
      flex-wrap: wrap;
    }
  } */
</style>
