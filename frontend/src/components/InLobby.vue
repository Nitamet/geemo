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
                <div class="champion-and-builds"></div>
                <div class="build-info"></div>
            </div>
        </div>
    </div>
</template>

<style lang="scss">
.lobby-info {
    width: 80%;
    max-width: 1600px;
    min-width: 1200px;
}

.main-container {
    min-height: 600px;
    width: 100%;
    background-color: red;
}

.champion-and-builds {
    width: 30%;
    max-width: 300px;
    background-color: green;
}
</style>
<script setup lang="ts">
import SummonerInfo from 'components/Lobby/SummonerInfo.vue';
</script>
