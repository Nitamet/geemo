<template>
    <q-img
        class="summoner-spell"
        :src="props.iconUrl"
        :width="size"
        :height="size"
    >
        <q-tooltip class="secondary text-body2" :offset="[10, 10]">
            {{ props.name }}
        </q-tooltip>
    </q-img>
</template>

<script setup lang="ts">
const props = defineProps({
    name: {
        type: String,
        required: true,
    },
    iconUrl: {
        type: String,
        required: true,
    },
});

const size = '48px';
</script>

<style lang="scss">
.summoner-spell {
    border-radius: 4px;
}
</style>
