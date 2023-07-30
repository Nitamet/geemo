<template>
    <div class="league-state row items-center justify-evenly">
        <q-img
            src="~assets/lol-icon.jpg"
            :class="[
                'lol-icon',
                leagueState !== LeagueState.NotLaunched ? 'active' : '',
            ]"
        />
        <h3>{{ props.message }}</h3>
        <q-btn
            class="button"
            label="Retry"
            size="xl"
            color="primary"
            padding="xs lg"
            push
            @click="updateState"
        />
    </div>
</template>

<script setup>
import { GetState } from 'app/wailsjs/go/main/App';
import { LeagueState, useApplicationStore } from 'stores/application-store';
import { computed } from 'vue';

const props = defineProps({ message: String });

const application = useApplicationStore();
const leagueState = computed({
    get() {
        return application.leagueState;
    },
    set(val) {
        application.leagueState = LeagueState[val] ?? LeagueState.Unknown;
    },
});
const updateState = async () => {
    const leagueStateString = await GetState();

    leagueState.value = LeagueState[leagueStateString] ?? LeagueState.Unknown;
};
</script>

<style lang="scss">
.league-state {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.lol-icon {
    width: 256px;
    height: 256px;
    filter: grayscale(1);
    transition: all 1s ease;
}
</style>
