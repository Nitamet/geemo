<template>
    <div
        :class="[
            'build',
            'row',
            'q-gutter-x-sm',
            'items-center',
            'q-py-sm',
            isSelected ? 'selected' : '',
        ]"
        @click="selectBuild(props.build)"
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
import { lcu, lolbuild } from 'app/wailsjs/go/models';

import Build = lolbuild.Build;
import {
    ApplyRunes,
    ApplySummonerSpells,
    ApplyItemSet,
} from 'app/wailsjs/go/main/App';
import { useApplicationStore } from 'stores/application-store';
import { storeToRefs } from 'pinia';
import { ref } from 'vue';
import { whenever } from '@vueuse/core';
import ItemSet = lcu.ItemSet;

const props = defineProps<{
    build: Build;
    championName: string;
    source: string;
}>();

const application = useApplicationStore();
const { selectedBuild } = storeToRefs(application);

let isSelected = ref(false);

const getCoreItem = () => {
    if ('' !== props.build.items.mythic.name) {
        return props.build.items.mythic;
    }

    // If we don't have a mythic item, we'll use the second core item
    // Because the first core item is the boots
    return props.build.items.core[1];
};

const selectBuild = (build: Build) => {
    isSelected.value = true;
    selectedBuild.value = build;
    const selectedPerks = build.selectedPerks.map((perk) => perk.id);

    const runePage = {
        name: `${props.source}: ${build.name} ${props.championName}`,
        primaryStyleId: build.primary.id,
        selectedPerkIds: selectedPerks,
        subStyleId: build.secondary.id,
        current: true,
    };

    ApplyRunes(runePage);

    const summonerSpells = {
        firstSpellId: build.summonerSpells.at(0)?.id ?? 0,
        secondSpellId: build.summonerSpells.at(1)?.id ?? 0,
    };

    ApplySummonerSpells(
        summonerSpells.firstSpellId,
        summonerSpells.secondSpellId
    );

    const convertToItemSet = (items: lolbuild.Item[]) =>
        items.map((item) => ({
            id: item.id.toString(),
            count: 1,
        }));

    const itemBlocks = [
        {
            type: 'Starting',
            items: convertToItemSet(build.items.starting),
        },
        {
            type: 'Core',
            items: convertToItemSet(
                build.items.core.concat(build.items.mythic)
            ),
        },
        {
            type: 'Fourth',
            items: convertToItemSet(build.items.fourth),
        },
        {
            type: 'Fifth',
            items: convertToItemSet(build.items.fifth),
        },
        {
            type: 'Sixth',
            items: convertToItemSet(build.items.sixth),
        },
    ];

    const itemSet = {
        title: `${props.championName}: ${build.name}`,
        associatedChampions: [],
        associatedMaps: [],
        type: 'custom',
        map: 'any',
        mode: 'any',
        startedFrom: 'blank',
        uid: '1',
        preferredItemSlots: [],
        sortrank: 0,
        blocks: itemBlocks,
    };

    ApplyItemSet(ItemSet.createFrom(itemSet));
};

whenever(selectedBuild, () => {
    if (selectedBuild.value !== props.build) {
        isSelected.value = false;
    }
});
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
