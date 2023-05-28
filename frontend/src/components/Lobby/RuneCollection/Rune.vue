<template>
    <q-img
        :src="props.iconUrl"
        :width="size"
        :height="size"
        :class="[isActive ? 'active' : '']"
    />
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

const size = props.isKeystone ? '96px' : '64px';
const findPerk = () => {
    if (selectedBuild.value) {
        console.log(
            selectedBuild.value.selectedPerks.find(
                (perk) => perk.id === props.id
            )
        );
        return selectedBuild.value.selectedPerks.find(
            (perk) => perk.id === props.id
        );
    }

    return null;
};
const isActive = findPerk() != null;
</script>
