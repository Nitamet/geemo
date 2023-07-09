import { defineStore } from 'pinia';

export const useSettingsStore = defineStore('settings', {
    state: () => ({
        autoImport: false,
        showNativeTitleBar: true,
    }),
    getters: {},
    actions: {},
});
