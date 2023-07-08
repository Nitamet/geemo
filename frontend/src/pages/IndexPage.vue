<template>
    <q-page
        :class="
            application.leagueState !== LeagueState.InLobby
                ? 'row items-center justify-evenly'
                : 'row items-center'
        "
    >
        <InLobby v-if="application.leagueState === LeagueState.InLobby" />
        <NotInLobby v-else :message="application.getLeagueStateMessage" />
    </q-page>
</template>

<script setup lang="ts">
import { useApplicationStore } from 'stores/application-store';
import { LeagueState } from 'stores/application-store';
import NotInLobby from 'components/NotInLobby.vue';
import InLobby from 'components/InLobby.vue';

const application = useApplicationStore();
</script>
