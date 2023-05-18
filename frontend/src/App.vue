<template>
    <router-view />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useApplicationStore } from './stores/application-store';

const applicationStore = useApplicationStore();
onMounted(() => {
    applicationStore.startCheckingLeagueState();
});
</script>
