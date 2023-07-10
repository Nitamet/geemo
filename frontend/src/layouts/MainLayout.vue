<template>
    <q-layout view="lHh Lpr lFf ">
        <q-bar v-if="!settingsStore.showNativeTitleBar" class="bar">
            <q-space />
            <q-btn dense flat icon="minimize" @click="Minimize()" />
            <q-btn dense flat icon="crop_square" @click="Maximize()" />
            <q-btn dense flat icon="close" @click="Close()" />
        </q-bar>
        <q-page-container class="page">
            <router-view class="router-page" />
            <Settings />
        </q-page-container>
    </q-layout>
</template>

<script setup lang="ts">
import Settings from 'components/Settings.vue';
import { Close, Maximize, Minimize } from 'app/wailsjs/go/main/App';
import { useSettingsStore } from 'stores/settings-store';

const settingsStore = useSettingsStore();
</script>

<style lang="scss">
.router-page {
    min-height: calc(100vh - 32px) !important;
}

.bar {
    --wails-draggable: drag;
    background-color: rgba(0, 0, 0, 0);
}

.page {
    position: relative;
}
</style>
