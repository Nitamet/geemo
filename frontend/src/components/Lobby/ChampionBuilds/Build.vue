<template>
    <div
        :class="[
            'build',
            'row',
            'q-gutter-x-sm',
            'items-center',
            'q-py-sm',
            props.isSelected ? 'selected' : '',
        ]"
        @click="
            emit('buildClicked', { build: props.build, source: props.source })
        "
    >
        <q-avatar size="42px" rounded>
            <img :src="getCoreItem().iconUrl" :alt="getCoreItem().name" />
        </q-avatar>
        <span>{{ props.build.name }}</span>
        <q-img
            class="q-ml-auto"
            width="38px"
            height="38px"
            :src="`https://ddragon.leagueoflegends.com/cdn/img/${props.build.selectedPerks[0].iconUrl}`"
            rounded
        />
    </div>
</template>

<script setup lang="ts">
import { lolbuild } from 'app/wailsjs/go/models';
import Build = lolbuild.Build;

interface Props {
    build: Build;
    championName: string;
    source: string;
    isSelected: boolean;
}

const props = defineProps<Props>();

const emit = defineEmits<{
    (e: 'buildClicked', value: { build: Build; source: string }): void;
}>();

const getCoreItem = () => {
    if ('' !== props.build.items.mythic.name) {
        return props.build.items.mythic;
    }

    // If we don't have a mythic item, we'll use the second core item
    // Because the first core item is the boots
    return props.build.items.core[1];
};
</script>

<style scoped lang="scss">
@import '../../../css/variables.scss';

.build {
    border-radius: 4px;
    transition: background-color 0.1s ease-in-out;
}

.build:hover {
    cursor: pointer;
}

.selected {
    background-color: $divider-color;
}
</style>
