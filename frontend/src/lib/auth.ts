import { useUserStore, useDeviceStore } from "@/lib/store"
import Cookies from "js-cookie";


export const fetchUser = async () => {
    const userStore = useUserStore();
    const deviceStore = useDeviceStore();

    const token = Cookies.get('token');
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/profile`, {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    const data = await response.json();
    if (response.ok) {
        userStore.setUsername(data.username);
    }

    await deviceStore.fetchDevices();
};

export const logOut = async () => {
    const userStore = useUserStore();
    const deviceStore = useDeviceStore();

    userStore.resetState();
    deviceStore.resetState();
    Cookies.remove('token');
    deviceStore.fetchDevices();
};
