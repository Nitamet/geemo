<template>
    <div class="champion-builds q-pa-lg">
        <Champion
            :champion-name="currentChampionName"
            :champion-icon-url="currentChampionIconUrl"
        />
        <q-separator class="q-mt-md separator" />
        <div class="sources q-mt-sm">
            <div
                v-for="buildCollection in builds"
                :key="buildCollection.source"
            >
                <span class="text-subtitle2">{{ buildCollection.source }}</span>
                <div class="builds q-gutter-y-md">
                    <Build
                        v-for="build in buildCollection.runes"
                        :key="build.name"
                        :build="build"
                        :champion-name="currentChampionName"
                        :source="buildCollection.source"
                    />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import Champion from 'components/Lobby/ChampionBuilds/Champion.vue';
import { delay } from 'src/util/misc';
import { GetCurrentChampion } from 'app/wailsjs/go/main/App';
import { computed, Ref, ref } from 'vue';
import { whenever } from '@vueuse/core';
import { LeagueState, useApplicationStore } from 'stores/application-store';
import { storeToRefs } from 'pinia';
import { lolbuild } from 'app/wailsjs/go/models';
import BuildCollection = lolbuild.BuildCollection;
import { LoadBuilds } from 'app/wailsjs/go/lolbuild/Loader';
import Build from 'components/Lobby/ChampionBuilds/Build.vue';

let currentChampion = ref(-1);
let currentChampionName = ref('Champion');
const currentChampionIconUrl = computed(() => {
    return `https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/${currentChampion.value}.png`;
});

const application = useApplicationStore();
const { leagueState } = storeToRefs(application);
const startCheckingCurrentChampion = async () => {
    await delay(3000);
    const champion = await GetCurrentChampion();
    if (champion !== 0) {
        currentChampion.value = champion;
    }

    if (leagueState.value === LeagueState.InLobby) {
        await startCheckingCurrentChampion();
    }
};
startCheckingCurrentChampion();

let builds: Ref<BuildCollection[]> = ref([]);

whenever(currentChampion, async () => {
    const resp = await fetch(
        `https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champions/${currentChampion.value}.json`
    );
    const json: { name: string } = await resp.json();
    currentChampionName.value = json.name;

    builds.value = await LoadBuilds(currentChampionName.value, ['ugg']);
});
</script>

<style lang="scss">
@import '../../css/variables.scss';

.champion-builds {
    width: 30%;
    max-width: 340px;
    background-color: $build-selection-background-color;
    border-right: 2px solid $divider-color;
}
</style>
