<template>
    <div class="champion row q-gutter-sm items-center">
        <q-avatar size="56px" rounded
            ><img :src="props.championIconUrl" :alt="props.championName"
        /></q-avatar>
        <div class="text-weight-bolder champion-name">
            {{ props.championName }}
        </div>
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
.champion-name {
    max-width: 100px;
    font-size: 18px;
}
</style>
