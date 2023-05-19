<template>
    <q-page>
        <!--<NotInLobby :message="application.getLeagueStateMessage" />-->
        <InLobby />
    </q-page>
</template>

<script setup lang="ts">
import { useApplicationStore } from 'stores/application-store';
import NotInLobby from 'components/NotInLobby.vue';
import InLobby from 'components/InLobby.vue';

const application = useApplicationStore();
</script>
