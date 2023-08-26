<template>
    <div class="text-center q-pa-md flex flex-center">
        <div>
            <div class="text-h2">An error has occurred. Restart the app.</div>

            <div class="q-mt-md row q-gutter-lg flex-center">
                <q-btn
                    class="button q-mt-lg"
                    :label="$t('showLog')"
                    size="xl"
                    color="primary"
                    padding="xs lg"
                    push
                    @click="OpenLogFolder()"
                />
                <q-btn
                    class="button q-mt-lg"
                    :label="$t('restartApp')"
                    size="xl"
                    color="secondary"
                    padding="xs lg"
                    push
                    @click="Restart()"
                />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { OpenLogFolder, Restart } from 'app/wailsjs/go/main/App';
</script>
