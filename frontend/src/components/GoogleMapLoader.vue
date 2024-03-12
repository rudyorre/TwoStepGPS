<template>
  <div id="map" class="h-full"></div>
</template>

<script setup lang="ts">
import { onMounted, nextTick, defineProps, watch } from 'vue'
import { Loader } from '@googlemaps/js-api-loader'
import { Device } from '@/lib/types'
// import { renderToString } from '@vue/server-renderer'
// import DeviceMarker from '@/components/DeviceMarker.vue'

let map: google.maps.Map | null = null;
const props = defineProps<{
  selectedDevice: Device | null,
  devices: Device[],
}>();

const emit = defineEmits(['update:selectedDevice']);
const selectDevice = (device: Device) => {
    emit('update:selectedDevice', device);
};
watch(() => props.selectedDevice, (newDevice, _) => {
  if (newDevice && newDevice.latitude && newDevice.longitude) {
    map?.panTo({ lat: newDevice.latitude, lng: newDevice.longitude });
  }
});

const loader = new Loader({
  apiKey: import.meta.env.VITE_MAPS_API_KEY,
  version: 'weekly',
  libraries: ['places']
});

onMounted(async () => {
  const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/device-locations`);
  const devices: Device[] = await response.json();

  await nextTick();
  if (document.getElementById('map')) {
    // initMap();
    loader.load().then(async () => {
      const { Map } = await google.maps.importLibrary("maps") as google.maps.MapsLibrary;
      const { AdvancedMarkerElement } = await google.maps.importLibrary("marker") as google.maps.MarkerLibrary;

      map = new Map(document.getElementById("map") as HTMLElement, {
        center: { lat: 34.0549, lng: -118.2426 },
        zoom: 10,
        mapId: 'DEMO_MAP_ID',
      });

      for (const device of devices) {
        const markerElement = document.createElement('div');
        markerElement.style.width = '30px';
        markerElement.style.height = '30px';
        markerElement.style.cursor = 'pointer';
        markerElement.innerHTML = `
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="red" style="transform: rotate(${device.angle}deg);">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
          </svg>
        `;

        const marker = new AdvancedMarkerElement({
          map,
          position: { lat: device.latitude, lng: device.longitude },
          content: markerElement,
        });

        marker.addListener("click", () => {
          map?.panTo({ lat: device.latitude, lng: device.longitude });
          selectDevice(device);
          // const color = (markerElement.innerHTML.includes("red")) ? "blue" : "red";
          // markerElement.innerHTML = `
          //   <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke=${color} style="transform: rotate(${deviceLocation.angle}deg);">
          //     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
          //   </svg>
          // `;
        })
      }
    });
  }
});

</script>

<style scoped>
.price-tag {
  background-color: #4285F4;
  border-radius: 8px;
  color: #FFFFFF;
  font-size: 24px;
  padding: 10px 15px;
  position: relative;
}

.price-tag::after {
  content: "";
  position: absolute;
  left: 50%;
  top: 100%;
  transform: translate(-50%, 0);
  width: 0;
  height: 0;
  border-left: 8px solid transparent;
  border-right: 8px solid transparent;
  border-top: 8px solid #4285F4;
}
</style>