<template>
    <div v-if="null !== selectedBuild" class="build-info q-pa-lg">
        <span class="text-h5 text-weight-bold">RUNES</span>
        <div class="rune-trees row no-wrap q-mt-lg q-gutter-x-sm">
            <RuneTree />
            <div class="rune-tree column">
                <RuneTree secondary />
                <StatMods />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import StatMods from 'components/Lobby/RuneCollection/StatMods.vue';
import RuneTree from 'components/Lobby/RuneCollection/RuneTree.vue';
import { useApplicationStore } from 'stores/application-store';
import { storeToRefs } from 'pinia';

const application = useApplicationStore();
const { selectedBuild } = storeToRefs(application);
</script>
