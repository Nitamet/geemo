<template>
    <div class="in-lobby q-mt-xl row items-center justify-evenly">
        <div class="lobby-info">
            <div class="row justify-between items-center">
                <SummonerInfo />
                <div class="current-game-mode">
                    <span class="text-subtitle1">Summoner's Rift</span>
                </div>
            </div>
            <div class="main-container q-mt-lg row items-stretch">
                <ChampionBuilds />
                <div class="build-info q-pa-lg">
                    <span class="text-h5 text-weight-bold">RUNES</span>
                    <div class="rune-trees row no-wrap q-mt-lg q-gutter-x-sm">
                        <RuneTree />
                        <RuneTree secondary />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import ChampionBuilds from 'components/Lobby/ChampionBuilds.vue';
import SummonerInfo from 'components/Lobby/SummonerInfo.vue';
import RuneTree from 'components/Lobby/RuneCollection/RuneTree.vue';
</script>

<style lang="scss">
@import '../css/variables.scss';

.main-container {
    min-height: 600px;
    width: 100%;
    background-color: $build-info-background-color;
}

.lobby-info {
    width: 80%;
    max-width: 1600px;
    min-width: 1200px;
}

.build-info {
    flex: 1;
}
</style>
