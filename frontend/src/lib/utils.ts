import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'
import Cookies from 'js-cookie'
import { type Ref } from 'vue'
import type { Updater } from '@tanstack/vue-table'
import { useDeviceStore } from '@/lib/store'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const isUserLoggedIn = () => {
  return !!Cookies.get('token');
};

export function valueUpdater<T extends Updater<any>>(updaterOrValue: T, ref: Ref) {
  ref.value
    = typeof updaterOrValue === 'function'
      ? updaterOrValue(ref.value)
      : updaterOrValue
}

export const fetchDevices = async () => {
  const deviceStore = useDeviceStore();
  if (isUserLoggedIn()) {
      const token = Cookies.get('token');
      const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/get-device-settings`, {
          method: 'GET',
          headers: {
              Authorization: `Bearer ${token}`,
          },
      });
      const json = await response.json();
      deviceStore.setDevices(json);

  } else {
      const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/device-locations`);
      const json = await response.json();
      deviceStore.setDevices(json);
  }
};
