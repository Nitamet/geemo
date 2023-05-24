<template>
    <div class="champion row q-gutter-sm items-center">
        <q-avatar size="56px" rounded
            ><img :src="props.championIconUrl" :alt="props.championName"
        /></q-avatar>
        <span class="text-weight-bolder text-h6">{{ props.championName }}</span>
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
