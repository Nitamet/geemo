<template>
    <q-img
        class="summoner-spell"
        :src="props.iconUrl"
        :width="size"
        :height="size"
    />
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
