<template>
    <div class="champion-builds q-pa-lg">
        <Champion
            :champion-name="currentChampionName"
            :champion-icon-url="currentChampionIconUrl"
        />
    </div>
</template>

<script setup lang="ts">
import Champion from 'components/Lobby/ChampionBuilds/Champion.vue';
import { delay } from 'src/util/misc';
import { GetCurrentChampion } from 'app/wailsjs/go/main/App';
import { computed, ref } from 'vue';
import { whenever } from '@vueuse/core';

let currentChampion = ref(-1);

let currentChampionName = ref('Champion');

const currentChampionIconUrl = computed(() => {
    return `https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/${currentChampion.value}.png`;
});

const startCheckingCurrentChampion = async () => {
    await delay(3000);
    const champion = await GetCurrentChampion();
    if (champion !== 0) {
        currentChampion.value = champion;
    }
    await startCheckingCurrentChampion();
};
startCheckingCurrentChampion();

whenever(currentChampion, async () => {
    const resp = await fetch(
        `https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champions/${currentChampion.value}.json`
    );
    const json: { name: string } = await resp.json();
    currentChampionName.value = json.name;
});
</script>

<style lang="scss">
@import '../../css/variables.scss';

.champion-builds {
    width: 30%;
    max-width: 300px;
    background-color: $build-selection-background-color;
    border-right: 2px solid $divider-color;
}
</style>
