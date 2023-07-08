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
        <div class="q-ml-auto row align-center">
            <span
                :class="['winrate', props.build.winrate > 50 ? 'good' : 'bad']"
                ><strong>{{ props.build.winrate }}%</strong>
            </span>
            <q-img
                class="q-ml-sm"
                width="38px"
                height="38px"
                :src="`https://ddragon.leagueoflegends.com/cdn/img/${props.build.selectedPerks[0].iconUrl}`"
                rounded
            />
        </div>
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

.winrate {
    background-color: red;
    border-radius: 4px;
    padding: 6px;
}

.good {
    background-color: #135936;
    color: #17f84a;
}

.bad {
    background-color: #6b1c22;
    color: #ff7984;
}
</style>
