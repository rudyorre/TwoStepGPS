import { defineStore } from 'pinia';
import { Device } from '@/lib/types'
import { isUserLoggedIn } from '@/lib/utils';
import Cookies from 'js-cookie';

export const useUserStore = defineStore({
    id: 'user',
    state: () => ({
        username: null as string | null,
    }),
    actions: {
        setUsername(username: string) {
            this.username = username;
        },
        resetState() {
            this.username = null;
        }
    },
});

export const useDeviceStore = defineStore({
    id: 'devices',
    state: () => ({
        devices: [] as Device[],
        selectedDevice: null as Device | null,
        intervalId: null as number | null,
    }),
    actions: {
        async fetchDevices() {
            let response;
            if (isUserLoggedIn()) {
                const token = Cookies.get('token');
                response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/get-device-settings`, {
                    method: 'GET',
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
            } else {
                response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/device-locations`);
            }
            const json = await response.json();
            this.setDevices(json);
        },
        async startPolling(pollingRateSeconds: number) {
            // Fetch the devices immediately
            await this.fetchDevices();

            // Start polling
            this.intervalId = window.setInterval(() => this.fetchDevices(), pollingRateSeconds * 1000);

            // Pause polling when the user is not at the website
            document.addEventListener('visibilitychange', () => {
                if (document.visibilityState === 'hidden' && this.intervalId !== null) {
                    clearInterval(this.intervalId);
                    this.intervalId = null;
                } else {
                    this.intervalId = window.setInterval(() => this.fetchDevices(), pollingRateSeconds * 1000);
                }
            });
        },
        setSelectedDevice(device: Device | null) {
            this.selectedDevice = device;
        },
        setDevices(updatedDevices: Device[]) {
            this.devices = updatedDevices;
        },
        setNickname(deviceId: string, nickname: string) {
            const device = this.devices.find(device => device.device_id === deviceId);
            if (device) {
                device.nickname = nickname;
            }
        },
        setColor(device: Device, color: string) {
            const targetDevice = this.devices.find(d => d.device_id === device.device_id);
            if (targetDevice) {
                targetDevice.color = color;
            }
        },
        setHidden(device: Device, hidden: boolean) {
            const deviceId = device.device_id
            const targetDevice = this.devices.find(device => device.device_id === deviceId);
            if (targetDevice) {
                targetDevice.is_hidden = hidden;
            }
        },
        resetState() {
            this.devices = [];
            this.selectedDevice = null;
        }
    }
});
