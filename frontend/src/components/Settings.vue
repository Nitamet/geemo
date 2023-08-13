<template>
    <q-btn
        class="button settings-btn"
        color="primary"
        icon="settings"
        size="15px"
        push
        @click="showSettings = true"
    />

    <q-dialog v-model="showSettings">
        <q-card style="width: 500px" class="settings q-pa-md column">
            <span>App Version: {{ version }}</span>
            <div>
                <q-checkbox
                    v-model="autoUpdate"
                    label="Auto update (restart required to take effect)"
                    color="teal"
                />
                <q-checkbox
                    v-model="showNativeTitleBar"
                    label="Show native title bar (restart required to take effect)"
                    color="teal"
                />
                <q-checkbox
                    v-model="autoImport"
                    label="Autoimport selected build"
                    color="teal"
                />
            </div>
        </q-card>
    </q-dialog>
</template>

<script setup lang="ts">
import { onBeforeMount, ref, watch } from 'vue';
import {
    GetAutoImportSetting,
    GetAutoUpdateSetting,
    GetCurrentVersion,
    GetShowNativeTitleBarSetting,
    SetAutoImportSetting,
    SetAutoUpdateSetting,
    SetShowNativeTitleBarSetting,
} from 'app/wailsjs/go/main/App';
import { storeToRefs } from 'pinia';
import { useSettingsStore } from 'stores/settings-store';

const showSettings = ref(false);
const version = ref('');

const settingsStore = useSettingsStore();
const { autoImport, showNativeTitleBar, autoUpdate } =
    storeToRefs(settingsStore);

watch(autoImport, async (value) => {
    await SetAutoImportSetting(value);
});

watch(showNativeTitleBar, async (value) => {
    await SetShowNativeTitleBarSetting(value);
});

watch(autoUpdate, async (value) => {
    await SetAutoUpdateSetting(value);
});

onBeforeMount(async () => {
    autoImport.value = await GetAutoImportSetting();
    showNativeTitleBar.value = await GetShowNativeTitleBarSetting();
    autoUpdate.value = await GetAutoUpdateSetting();
    version.value = await GetCurrentVersion();
});
</script>

<style lang="scss">
.settings-btn {
    position: absolute;
    right: 18px;
    top: 14px;
    padding: 8px;
}

.settings {
    background-color: #0b1217;
}
</style>
