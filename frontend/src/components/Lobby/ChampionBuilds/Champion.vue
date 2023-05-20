<template>
    <div class="champion row q-gutter-sm items-center">
        <q-avatar size="48px" rounded
            ><img :src="props.championIconUrl" :alt="props.championName"
        /></q-avatar>
        <span class="text-weight-bold">{{ props.championName }}</span>
        <q-separator class="q-mt-md separator" />
    </div>
</template>

<script setup>
const props = defineProps({
    championName: {
        type: String,
        required: true,
    },
    championIconUrl: {
        type: String,
        required: true,
    },
});
</script>

<style lang="scss">
@import '../../../css/variables.scss';

.separator {
    width: 100%;
    height: 1px;
    background-color: $divider-color;
}
</style>
