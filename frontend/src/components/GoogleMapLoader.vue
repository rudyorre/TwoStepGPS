<template>
  <div id="map" class="h-full"></div>
</template>

<script setup lang="ts">
import { onMounted, nextTick } from 'vue'
import { Loader } from '@googlemaps/js-api-loader'

const loader = new Loader({
  apiKey: import.meta.env.VITE_MAPS_API_KEY,
  version: 'weekly',
  libraries: ['places']
})

// let map: any;

onMounted(async () => {
  await nextTick();
  if (document.getElementById('map')) {
    loader.load().then(async () => {
      const { Map } = await google.maps.importLibrary("maps") as google.maps.MapsLibrary;

      new Map(document.getElementById("map") as HTMLElement, {
        center: { lat: 34.0549, lng: -118.2426 },
        zoom: 8,
      });
    });
  }
});
</script>