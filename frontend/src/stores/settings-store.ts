import { defineStore } from 'pinia';

export const useSettingsStore = defineStore('settings', {
    state: () => ({
        autoImport: false,
        showNativeTitleBar: true,
        autoUpdate: true,
        language: 'en-US',
        activeSources: [] as string[],
    }),
    getters: {},
    actions: {},
});
