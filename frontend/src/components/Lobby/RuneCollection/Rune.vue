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
    isKeystone: {
        type: Boolean,
        default: false,
    },
});

const application = useApplicationStore();
const { selectedBuild } = storeToRefs(application);

let size = '44px';
if (props.isKeystone) {
    size = '64px';
}

const findPerk = () => {
    if (selectedBuild.value) {
        return selectedBuild.value.selectedPerks.find(
            (perk) => perk.id == props.id
        );
    }

    return null;
};
const isActive = findPerk() != null;
</script>
