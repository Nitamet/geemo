<template>
    <q-page class="row items-center justify-evenly">
        <NotInLobby :message="application.getLeagueStateMessage" />
    </q-page>
</template>

<script setup lang="ts">
import { useApplicationStore } from 'stores/application-store';
import NotInLobby from 'components/NotInLobby.vue';

const application = useApplicationStore();
</script>
