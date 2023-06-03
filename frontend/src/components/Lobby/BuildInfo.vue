<template>
    <div v-if="null !== selectedBuild" class="build-info q-pa-lg row no-wrap">
        <div class="runes">
            <span class="text-h5 text-weight-bold">RUNES</span>
            <div class="rune-trees row no-wrap q-mt-lg q-gutter-x-sm">
                <RuneTree />
                <div class="rune-tree column">
                    <RuneTree secondary />
                    <StatMods />
                </div>
            </div>
        </div>
        <div class="q-ml-lg column">
            <div class="summoner-spells">
                <span class="text-h5 text-weight-bold">SPELLS</span>
                <div class="row q-mt-md q-gutter-x-md">
                    <SummonerSpell
                        v-for="spell in selectedBuild.summonerSpells"
                        :name="spell.name"
                        :icon-url="spell.iconUrl"
                        :key="spell.id"
                    />
                </div>
                <q-separator class="q-mt-md full-width separator" />
            </div>
            <div class="items q-mt-md">
                <span class="text-h5 text-weight-bold">ITEMS</span>
                <div class="column q-gutter-y-sm">
                    <ItemGroup
                        group-name="Starting"
                        :items="selectedBuild.items.starting"
                    />
                    <ItemGroup
                        group-name="Core"
                        :items="[
                            selectedBuild.items.mythic,
                            ...selectedBuild.items.core,
                        ]"
                    />
                    <ItemGroup
                        group-name="Fourth"
                        :items="selectedBuild.items.fourth"
                    />
                    <ItemGroup
                        group-name="Fifth"
                        :items="selectedBuild.items.fifth"
                    />
                    <ItemGroup
                        group-name="Sixth"
                        :items="selectedBuild.items.sixth"
                    />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import StatMods from 'components/Lobby/RuneCollection/StatMods.vue';
import RuneTree from 'components/Lobby/RuneCollection/RuneTree.vue';
import { useApplicationStore } from 'stores/application-store';
import { storeToRefs } from 'pinia';
import SummonerSpell from 'components/Lobby/SummonerSpell.vue';
import Item from 'components/Lobby/Item.vue';
import ItemGroup from 'components/Lobby/ItemGroup.vue';

const application = useApplicationStore();
const { selectedBuild } = storeToRefs(application);
</script>

<style lang="scss">
@import '../../css/variables.scss';

.runes {
    border-right: 1px solid $divider-color;
    min-width: 75%;
}
</style>
