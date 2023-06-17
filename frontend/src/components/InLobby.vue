<template>
    <div class="in-lobby q-mt-xl row items-center justify-evenly">
        <div class="lobby-info">
            <div class="row justify-between items-center">
                <SummonerInfo />
                <div class="current-game-mode">
                    <span class="text-h5">{{ gameModeName }}</span>
                </div>
            </div>
            <div class="main-container q-mt-lg row items-stretch">
                <ChampionBuilds :game-mode="gameMode" />
                <BuildInfo />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import ChampionBuilds from 'components/Lobby/ChampionBuilds.vue';
import SummonerInfo from 'components/Lobby/SummonerInfo.vue';
import BuildInfo from 'components/Lobby/BuildInfo.vue';
import { GetGameMode } from 'app/wailsjs/go/main/App';
import { onBeforeMount, ref } from 'vue';
import { GameMode } from 'components/models';

let gameMode = ref<GameMode>(GameMode.None);
let gameModeName = ref<string>('');

onBeforeMount(async () => {
    const gameModeInfo = await GetGameMode();
    gameMode.value = gameModeInfo[0] as GameMode;
    gameModeName.value = gameModeInfo.at(1) ?? '';
});
</script>

<style lang="scss">
@import '../css/variables.scss';

.main-container {
    min-height: 600px;
    width: 100%;
    background-color: $build-info-background-color;
}

.lobby-info {
    width: 95%;
    max-width: 1600px;
    min-width: 1200px;
}

.build-info {
    flex: 1;
}
</style>
