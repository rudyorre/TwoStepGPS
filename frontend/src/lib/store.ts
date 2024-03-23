import { defineStore } from 'pinia';
import { Device } from '@/lib/types'

export const useUserStore = defineStore({
    id: 'user',
    state: () => ({
        username: '',
    }),
    actions: {
        setUsername(username: string) {
            this.username = username;
        },
    },
});

export const useDeviceStore = defineStore({
    id: 'devices',
    state: () => ({
        devices: [] as Device[],
        selectedDevice: null as Device | null,
    }),
    actions: {
        setSelectedDevice(device: Device) {
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
        setColor(deviceId: string, color: string) {
            const device = this.devices.find(device => device.device_id === deviceId);
            if (device) {
                device.color = color;
            }
        },
        setHidden(device: Device, hidden: boolean) {
            const deviceId = device.device_id
            const targetDevice = this.devices.find(device => device.device_id === deviceId);
            if (targetDevice) {
                targetDevice.is_hidden = hidden;
            }
        }
    }
});
