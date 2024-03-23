<script setup lang="ts">
import { MoreHorizontal } from 'lucide-vue-next'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { toast } from 'vue-sonner'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from '@/components/ui/dialog'
import { Device } from '@/lib/types'
import { ref, defineEmits } from 'vue';
import Cookies from 'js-cookie'
import { isUserLoggedIn } from '@/lib/utils'
import { useRouter } from 'vue-router'
import { useDeviceStore } from '@/lib/store'

const deviceStore = useDeviceStore();

const props = defineProps<{
  device: Device
}>();
const emit = defineEmits(['nickname-changed']);

const color = props.device.color ? props.device.color : '#AA4A44';

function copy(id: string) {
  navigator.clipboard.writeText(id);
  const now = new Date();
  const options: Intl.DateTimeFormatOptions = { weekday: 'long', year: 'numeric',  month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };
  const formattedDate = now.toLocaleDateString('en-US', options);
  toast('Device ID copied to clipboard', {
    description: formattedDate,
  });
}

const selectedColor = ref(color);
const router = useRouter();


const saveChanges = async () => {
  if (isUserLoggedIn()) {
    const token = Cookies.get('token');
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/change-color`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        device_id: props.device.device_id,
        color: selectedColor.value,
      }),
    });
  
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
  
    location.reload();
  
    toast('Updated color', {
      description: "",
    });
  } else {
    toast('Your preference will not be saved.', {
      description: 'You must be logged in to save your preferences.',
      action: {
        label: 'Login',
        onClick: () => router.push('/login'),
      },
    });
  }
};

let nickname = ref(props.device.nickname);

const submitNickname = async () => {
  // Update the nickname locally
  deviceStore.setNickname(props.device.device_id, nickname.value);

  // Update the nickname on the backend
  if (isUserLoggedIn()) {
    const token = Cookies.get('token');
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/change-nickname`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        device_id: props.device.device_id,
        nickname: nickname.value,
      }),
    });
  
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
  
    toast(nickname.value === "" ? 'Removed nickname' : `Updated nickname to ${nickname.value}`, {
      description: "",
    });

    // Clear the input field
    nickname.value = '';
  } else {
    toast('Your preference will not be saved.', {
      description: 'You must be logged in to save your preferences.',
      action: {
        label: 'Login',
        onClick: () => router.push('/login'),
      },
    });
  }
};
</script>

<template>
  <DropdownMenu>
    <DropdownMenuTrigger as-child>
      <Button variant="ghost" class="h-8 w-8 p-0">
        <span class="sr-only">Open menu</span>
        <MoreHorizontal class="h-4 w-4" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent align="end">
      <DropdownMenuLabel>Device</DropdownMenuLabel>
      <DropdownMenuItem @click="copy(device.device_id)" class="cursor-pointer">
        <div class="font-bold">ID:</div>&nbsp;{{ device.device_id }}
      </DropdownMenuItem>
      <DropdownMenuSeparator />
      <Dialog>
        <DialogTrigger as-child>
          <div
            class="relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 hover:bg-muted"
          >
            Change Nickname
          </div>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Change Nickname of {{ device.nickname ? device.nickname : device.display_name }}'s device</DialogTitle>
            <DialogDescription>
              Input a new nickname for the device.
            </DialogDescription>
          </DialogHeader>
    
          <form @submit.prevent="submitNickname">
            <input v-model="nickname" type="text" placeholder=" Enter new nickname" class="rounded-md"/>
            <DialogFooter>
              <DialogClose as-child>
                <Button variant="secondary">Cancel</Button>
              </DialogClose>
              <DialogClose as-child>
                <Button variant="default" type="submit">Save changes</Button>
              </DialogClose>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>
      <Dialog>
        <!-- <DialogTrigger as-child @click.stop>
          <DropdownMenuItem class="cursor-pointer" @click.stop>Change Color</DropdownMenuItem>
          Change Color
        </DialogTrigger> -->
        <DialogTrigger as-child>
          <div
            class="relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 hover:bg-muted"
          >
            Change Color
          </div>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Change Color of {{ device.nickname ? device.nickname : device.display_name }}'s device</DialogTitle>
            <DialogDescription>
              Use the color picker to change the color.
            </DialogDescription>
          </DialogHeader>
          
          <input type="color" v-model="selectedColor"/>

          <DialogFooter>
            <DialogClose as-child>
              <Button variant="secondary">Cancel</Button>
            </DialogClose>
            <DialogClose as-child>
              <Button variant="default" @click="saveChanges">Save changes</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </DropdownMenuContent>
  </DropdownMenu>
</template>