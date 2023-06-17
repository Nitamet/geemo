<template>
    <q-img
        :src="props.iconUrl"
        :width="size"
        :height="size"
        :class="[isActive ? 'active' : '']"
    >
        <q-tooltip class="secondary text-body2" :offset="[10, 10]">
            {{ props.name }}
        </q-tooltip>
    </q-img>
</template>

<script setup lang="ts">
import { useApplicationStore } from 'stores/application-store';
import { storeToRefs } from 'pinia';

const props = defineProps({
    id: {
        type: Number,
        required: true,
    },
    name: {
        type: String,
        required: true,
    },
    iconUrl: {
        type: String,
        required: true,
    },
    row: {
        type: Number,
        required: true,
    },
});

const application = useApplicationStore();
const { selectedBuild } = storeToRefs(application);

const size = '32px';

const findStatMod = () => {
    if (selectedBuild.value) {
        const selectedStatMods = selectedBuild.value.selectedPerks.slice(-3);

        return selectedStatMods.at(props.row - 1)?.id == props.id;
    }

    return null;
};
const isActive = findStatMod();
</script>
