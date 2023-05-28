<template>
    <div class="rune-grid">
        <Rune
            v-for="perk in props.perks"
            :id="perk.id"
            :key="perk.id"
            :icon-url="perk.iconUrl"
            :name="perk.name"
            :is-keystone="false"
        />
    </div>
</template>

<script setup lang="ts">
import Rune from 'components/Lobby/RuneCollection/Rune.vue';
import { lolbuild } from 'app/wailsjs/go/models';
import RuneData = lolbuild.Rune;

const props = defineProps<{ perks: RuneData[] }>();
</script>

<style lang="scss">
.rune-grid {
    width: 80%;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 22px 0;
    justify-items: center;
}
</style>
