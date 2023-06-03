<template>
    <div>
        <span class="text-subtitle2">{{ props.groupName }}</span>
        <div class="row q-gutter-x-md q-mt-sm">
            <Item
                v-for="item in props.items"
                :key="item.id"
                :name="item.name"
                :icon-url="item.iconUrl"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { lolbuild } from 'app/wailsjs/go/models';
import ItemData = lolbuild.Item;
import Item from 'components/Lobby/Item.vue';

interface Props {
    groupName: string;
    items: ItemData[];
}

const props = defineProps<Props>();
</script>
