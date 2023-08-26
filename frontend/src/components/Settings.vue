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
            <div>
                {{ $t('appVersion') }}: {{ version }}
                <q-btn
                    class="button"
                    v-if="isUpdateAvailable"
                    :label="$t('update')"
                    size="md"
                    color="primary"
                    padding="sm"
                    push
                    @click="Update()"
                />
            </div>
            <div>
                {{ $t('appSettings') }}:
                <q-checkbox
                    v-model="autoUpdate"
                    :label="$t('autoUpdateOption')"
                    color="teal"
                    dark
                />
                <q-checkbox
                    v-model="showNativeTitleBar"
                    :label="$t('showNativeTitleBarOption')"
                    color="teal"
                    dark
                />
                <q-checkbox
                    v-model="autoImport"
                    :label="$t('autoImportSelectedBuildOption')"
                    color="teal"
                    dark
                />
                <br />
                {{ $t('activeSources') }}:
                <q-option-group
                    :options="allSources"
                    type="checkbox"
                    v-model="activeSources"
                    dark
                />
                {{ $t('other') }}:
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
    GetActiveSources,
    GetAutoImportSetting,
    GetAutoUpdateSetting,
    GetCurrentVersion,
    GetLanguage,
    GetShowNativeTitleBarSetting,
    IsUpdateAvailable,
    SetActiveSources,
    SetAutoImportSetting,
    SetAutoUpdateSetting,
    SetLanguage,
    SetShowNativeTitleBarSetting,
    Update,
} from 'app/wailsjs/go/main/App';
import { storeToRefs } from 'pinia';
import { useSettingsStore } from 'stores/settings-store';
import { useI18n } from 'vue-i18n';
import { GetSources } from 'app/wailsjs/go/lolbuild/Loader';

const showSettings = ref(false);
const version = ref('');

const settingsStore = useSettingsStore();
const { autoImport, showNativeTitleBar, autoUpdate, language, activeSources } =
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

watch(activeSources, async (value) => {
    await SetActiveSources(value);
});

const locale = useI18n({ useScope: 'global' }).locale;
const localeOptions = [
    { label: 'English', value: 'en-US' },
    { label: 'Русский', value: 'ru-RU' },
];

watch(locale, async (value) => {
    language.value = value;

    await SetLanguage(value);
});

const allSources = ref<{ label: string; value: string }[]>([]);

const isUpdateAvailable = ref(false);

onBeforeMount(async () => {
    autoImport.value = await GetAutoImportSetting();
    showNativeTitleBar.value = await GetShowNativeTitleBarSetting();
    autoUpdate.value = await GetAutoUpdateSetting();
    version.value = await GetCurrentVersion();
    locale.value = await GetLanguage();
    activeSources.value = await GetActiveSources();

    const sources = await GetSources();
    for (const source of sources) {
        allSources.value.push({ label: source.name, value: source.slug });
    }

    if (activeSources.value.length === 0) {
        activeSources.value = sources.map((source) => source.slug);
    }

    isUpdateAvailable.value = await IsUpdateAvailable();
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
