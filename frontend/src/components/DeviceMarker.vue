<template>
  <div class="arrow-marker" :style="{ transform: `rotate(${rotation}deg)` }">
    <!-- SVG for an arrow icon -->
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
    </svg>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

const props = defineProps({
  rotation: {
    type: Number,
    default: 0
  }
})
</script>

<style scoped>
.arrow-marker {
  /* Add any styles you need for the arrow marker here */
  width: 24px;
  height: 24px;
}
</style>