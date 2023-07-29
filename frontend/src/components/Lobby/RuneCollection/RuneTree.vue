<template>
    <div v-if="tree !== null" class="rune-tree column">
        <div class="self-center">
            <q-img
                class="active"
                :src="tree.iconUrl"
                width="32px"
                height="32px"
            />
            <span class="q-ml-sm text-h6 text-weight-medium">{{
                tree.name
            }}</span>
        </div>
        <Keystones
            v-if="!props.secondary"
            class="q-mt-lg"
            :keystones="tree.keystones"
        />
        <Perks class="q-mt-md self-center" :perks="tree.perks" />
    </div>
</template>

<script setup lang="ts">
import Keystones from 'components/Lobby/RuneCollection/Keystones.vue';
import { ref, Ref } from 'vue';
import { lolbuild } from 'app/wailsjs/go/models';
import RuneTreeData = lolbuild.RuneTree;
import { whenever } from '@vueuse/core';
import { storeToRefs } from 'pinia';
import { useApplicationStore } from 'stores/application-store';
import { GetRuneTree } from 'app/wailsjs/go/lolbuild/Loader';
import Perks from 'components/Lobby/RuneCollection/Perks.vue';

const props = defineProps({
    secondary: {
        type: Boolean,
        default: false,
    },
});

const tree: Ref<RuneTreeData | null> = ref(null);

const application = useApplicationStore();
const { selectedBuild } = storeToRefs(application);

const loadTree = async () => {
    if (selectedBuild.value) {
        const name = props.secondary
            ? selectedBuild.value.secondary.name
            : selectedBuild.value.primary.name;
        tree.value = await GetRuneTree(name);
    }
};

loadTree();

whenever(selectedBuild, async () => {
    await loadTree();
});
</script>

<style lang="scss">
.rune-tree {
    flex-basis: 50%;

    .q-img {
        filter: grayscale(100%);
        opacity: 0.35;
    }
}
</style>
