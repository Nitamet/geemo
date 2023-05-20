<template>
    <div class="summoner-info row q-gutter-md items-center">
        <q-avatar color="blue" rounded size="64px">
            <img
                :src="`https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/profile-icons/${data.profileIconId}.jpg`"
                alt="Summoner Icon"
            />
        </q-avatar>
        <span class="text-h3">{{ data.displayName }}</span>
    </div>
</template>

<script setup lang="ts">
import { onMounted, reactive } from 'vue';
import { GetSummoner } from 'app/wailsjs/go/main/App';
import { useApplicationStore } from 'stores/application-store';
import { storeToRefs } from 'pinia';
import { whenever } from '@vueuse/core';

const data = reactive({
    displayName: 'Summoner',
    profileIconId: 0,
});

const updateSummoner = async () => {
    const summoner = await GetSummoner();
    if (!summoner) return;
    data.displayName = summoner.displayName;
    data.profileIconId = summoner.profileIconId;
};

const application = useApplicationStore();
const { leagueState } = storeToRefs(application);
whenever(leagueState, async () => {
    await updateSummoner();
});

onMounted(async () => {
    await updateSummoner();
});
</script>
