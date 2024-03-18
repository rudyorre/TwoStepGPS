<script setup lang="ts">
import { h, ref, defineEmits, watchEffect } from 'vue'
import type {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
} from '@tanstack/vue-table'
import {
  FlexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getSortedRowModel,
  useVueTable,
  Table as TableType,
} from '@tanstack/vue-table'
import { ArrowUpDown, ChevronDown } from 'lucide-vue-next'

import DropdownAction from '@/components/DataTableColumn.vue'
import { Button } from '@/components/ui/button'
import { Checkbox } from '@/components/ui/checkbox'
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Input } from '@/components/ui/input'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { valueUpdater } from '@/lib/utils'
import { Device } from '@/lib/types'
import { isUserLoggedIn } from '@/lib/utils'
import Cookies from 'js-cookie'
import { toast } from 'vue-sonner'
import { useRouter } from 'vue-router'

const router = useRouter();
const props = defineProps<{
  selectedDevice: Device | null,
  devices: Device[],
}>();

const emit = defineEmits(['update:selectedDevice', 'update:devices']);
const selectDevice = (device: Device) => {
    emit('update:selectedDevice', device);
};

const hideDevice = async (device: Device) => {
  const hide = !device.is_hidden;
  device.is_hidden = hide;

  if (isUserLoggedIn()) {
    const token = Cookies.get('token');
    await fetch(`${import.meta.env.VITE_BACKEND_URL}/hide-device`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        device_id: device.device_id,
        hide: hide,
      }),
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

const columns: ColumnDef<Device>[] = [
  {
    id: 'hide',
    header: "Hide",
    cell: ({ row }) => h(Checkbox, {
      'checked': row.original.is_hidden,
        'onClick': (event: Event) => {
            event.stopPropagation();
            hideDevice(row.original);
        },
        'ariaLabel': 'Select row',
    }),
    enableSorting: false,
    enableHiding: false,
  },

  {
    accessorKey: 'display_name',
    header: ({ column }) => {
      return h(Button, {
        variant: 'ghost',
        onClick: () => column.toggleSorting(column.getIsSorted() === 'asc'),
      }, () => ['Name', h(ArrowUpDown, { class: 'ml-2 h-4 w-4' })])
    },
    cell: ({ row }) => h(
      'div',
      { class: '' },
      row.original.nickname ? `${row.original.nickname} (${row.original.display_name})` : row.original.display_name 
    ),
  },
  {
    id: 'actions',
    enableHiding: false,
    cell: ({ row }) => {
        const device = row.original

        return h('div', {
            class: 'relative',
            'onClick': (event: Event) => {
                event.stopPropagation();
            },
        }, h(DropdownAction, {
        device,
      }))
    },
  },
]

const sorting = ref<SortingState>([])
const columnFilters = ref<ColumnFiltersState>([])
const columnVisibility = ref<VisibilityState>({})
const rowSelection = ref({})

let table = ref<TableType<Device> | null>(null);

watchEffect(() => {
    if (props.devices) {
        table.value = useVueTable({
          data: props.devices,
          columns,
          getCoreRowModel: getCoreRowModel(),
          getSortedRowModel: getSortedRowModel(),
          getFilteredRowModel: getFilteredRowModel(),
          onSortingChange: updaterOrValue => valueUpdater(updaterOrValue, sorting),
          onColumnFiltersChange: updaterOrValue => valueUpdater(updaterOrValue, columnFilters),
          onColumnVisibilityChange: updaterOrValue => valueUpdater(updaterOrValue, columnVisibility),
          onRowSelectionChange: updaterOrValue => valueUpdater(updaterOrValue, rowSelection),
          state: {
            get sorting() { return sorting.value },
            get columnFilters() { return columnFilters.value },
            get columnVisibility() { return columnVisibility.value },
            get rowSelection() { return rowSelection.value },
          },
        });
    }
})
</script>

<template>
  <div class="w-full" v-if="table !== null">
    <div class="flex gap-2 items-center py-4">
      <Input
        class="max-w-sm"
        placeholder="Filter names..."
        :model-value="table.getColumn('display_name')?.getFilterValue() as string"
        @update:model-value=" table.getColumn('display_name')?.setFilterValue($event)"
      />
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="outline" class="ml-auto">
            Columns <ChevronDown class="ml-2 h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuCheckboxItem
            v-for="column in table.getAllColumns().filter((column) => column.getCanHide())"
            :key="column.id"
            class="capitalize"
            :checked="column.getIsVisible()"
            @update:checked="(value) => {
              column.toggleVisibility(!!value)
            }"
          >
            {{ column.id }}
          </DropdownMenuCheckboxItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
            <TableHead v-for="header in headerGroup.headers" :key="header.id">
              <FlexRender v-if="!header.isPlaceholder" :render="header.column.columnDef.header" :props="header.getContext()" />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <TableRow
                class="cursor-pointer"
                :class="{ 'bg-primary-foreground': row.original.device_id === selectedDevice?.device_id }"
                v-for="row in table.getRowModel().rows"
                :key="row.id"
                :data-state="row.getIsSelected() && 'selected'"
                @click="selectDevice(row.original)"
            >
              <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
                <FlexRender :render="cell.column.columnDef.cell" :props="cell.getContext()" />
              </TableCell>
            </TableRow>
          </template>

          <TableRow v-else>
            <TableCell
              col-span="{columns.length}"
              class="h-24 text-center"
            >
              No results.
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
  </div>
</template>