import { defineStore } from 'pinia';
import { Device } from '@/lib/types'

export const useDeviceStore = defineStore({
    id: 'devices',
    state: () => ({
        devices: [] as Device[],
    }),
    actions: {
        setDevices(updatedDevices: Device[]) {
            this.devices = updatedDevices;
        },
        setNickname(deviceId: string, nickname: string) {
            const device = this.devices.find(device => device.device_id === deviceId);
            if (device) {
                device.nickname = nickname;
            }
            console.log(nickname);
        }
    }
});
