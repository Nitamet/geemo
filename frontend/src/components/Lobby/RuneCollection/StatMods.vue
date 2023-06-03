<template>
    <div class="stat-mod-grid self-center q-mt-lg">
        <StatMod
            row="1"
            id="5008"
            name="Adaptive Force"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsAdaptiveForceIcon.png"
        />
        <StatMod
            row="1"
            id="5005"
            name="Attack Speed"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsAttackSpeedIcon.png"
        />
        <StatMod
            row="1"
            id="5007"
            name="Ability Haste"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsCDRScalingIcon.png"
        />
        <StatMod
            row="2"
            id="5008"
            name="Adaptive Force"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsAdaptiveForceIcon.png"
        />
        <StatMod
            row="2"
            id="5002"
            name="Armor"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsArmorIcon.png"
        />
        <StatMod
            row="2"
            id="5003"
            name="Magic Resist"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsMagicResIcon.png"
        />
        <StatMod
            row="3"
            id="5001"
            name="Health"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsHealthScalingIcon.png"
        />
        <StatMod
            row="3"
            id="5002"
            name="Armor"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsArmorIcon.png"
        />
        <StatMod
            row="3"
            id="5003"
            name="Magic Resist"
            icon-url="https://ddragon.leagueoflegends.com/cdn/img/perk-images/StatMods/StatModsMagicResIcon.png"
        />
    </div>
</template>

<script setup lang="ts">
import StatMod from 'components/Lobby/RuneCollection/StatMod.vue';
</script>

<style lang="scss">
.stat-mod-grid {
    width: 80%;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 22px 0;
    justify-items: center;
}
</style>
