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
            <span>{{ $t('appVersion') }}: {{ version }}</span>
            <div>
                <q-checkbox
                    v-model="autoUpdate"
                    :label="$t('autoUpdateOption')"
                    color="teal"
                />
                <q-checkbox
                    v-model="showNativeTitleBar"
                    :label="$t('showNativeTitleBarOption')"
                    color="teal"
                />
                <q-checkbox
                    v-model="autoImport"
                    :label="$t('autoImportSelectedBuildOption')"
                    color="teal"
                />
                <q-select
                    v-model="locale"
                    :options="localeOptions"
                    :label="$t('appLanguage')"
                    outlined
                    emit-value
                    map-options
                    dark
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
    GetLanguage,
    GetShowNativeTitleBarSetting,
    SetAutoImportSetting,
    SetAutoUpdateSetting,
    SetLanguage,
    SetShowNativeTitleBarSetting,
} from 'app/wailsjs/go/main/App';
import { storeToRefs } from 'pinia';
import { useSettingsStore } from 'stores/settings-store';
import { useI18n } from 'vue-i18n';

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

const locale = useI18n({ useScope: 'global' }).locale;
const localeOptions = [
    { label: 'English', value: 'en-US' },
    { label: 'Русский', value: 'ru-RU' },
];

watch(locale, async (value) => {
    await SetLanguage(value);
});

onBeforeMount(async () => {
    autoImport.value = await GetAutoImportSetting();
    showNativeTitleBar.value = await GetShowNativeTitleBarSetting();
    autoUpdate.value = await GetAutoUpdateSetting();
    version.value = await GetCurrentVersion();
    locale.value = await GetLanguage();
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
