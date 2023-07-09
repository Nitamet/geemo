<template>
    <q-layout view="lHh Lpr lFf">
        <q-page-container>
            <router-view />
            <Settings />
        </q-page-container>
    </q-layout>
</template>

<script setup lang="ts">
import Settings from 'components/Settings.vue';
</script>
