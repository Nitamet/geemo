<template>
    <div class="keystones row justify-between no-wrap">
        <Rune
            v-for="keystone in props.keystones"
            :id="keystone.id"
            :icon-url="keystone.iconUrl"
            :name="keystone.name"
            :key="keystone.id"
            is-keystone
        />
    </div>
</template>

<script setup lang="ts">
import Rune from 'components/Lobby/RuneCollection/Rune.vue';
import { lolbuild } from 'app/wailsjs/go/models';
import RuneData = lolbuild.Rune;

const props = defineProps<{ keystones: RuneData[] }>();
</script>
