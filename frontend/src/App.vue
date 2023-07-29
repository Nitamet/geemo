<template>
    <router-view />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useApplicationStore } from './stores/application-store';
import { EventsOn } from 'app/wailsjs/runtime';
import { useRouter } from 'vue-router';

const applicationStore = useApplicationStore();
onMounted(() => {
    applicationStore.startCheckingLeagueState();
});

const router = useRouter();
EventsOn('error', () => {
    router.push({ path: '/error' });
});
</script>
