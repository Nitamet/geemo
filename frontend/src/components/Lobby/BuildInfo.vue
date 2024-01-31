<template>
    <div v-if="null !== selectedBuild" class="build-info q-pa-lg row no-wrap">
        <div class="runes-skills">
            <span class="text-h5 text-weight-bold text-uppercase">{{
                $t('runes')
            }}</span>
            <div
                class="rune-trees row no-wrap q-mt-lg q-gutter-x-md justify-center"
            >
                <RuneTree />
                <div class="rune-tree column">
                    <RuneTree secondary />
                    <StatMods />
                </div>
            </div>
            <span class="text-h5 text-weight-bold text-uppercase">{{
                $t('skills')
            }}</span>
            <SkillOrder :skill-order="selectedBuild.skillOrder" />
        </div>
        <div class="q-ml-lg column">
            <div class="summoner-spells">
                <span class="text-h5 text-weight-bold text-uppercase">{{
                    $t('spells')
                }}</span>
                <div class="row q-mt-sm q-gutter-x-sm">
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
                <span class="text-h5 text-weight-bold text-uppercase">{{
                    $t('items')
                }}</span>
                <div class="column q-gutter-y-sm">
                    <ItemGroup
                        v-for="itemGroup in nonEmptyItemGroups"
                        :key="itemGroup.name"
                        :group-name="itemGroup.name"
                        :items="itemGroup.items"
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
import ItemGroup from 'components/Lobby/ItemGroup.vue';
import { computed } from 'vue';
import SkillOrder from 'components/Lobby/SkillOrder.vue';

const application = useApplicationStore();
const { selectedBuild } = storeToRefs(application);

const nonEmptyItemGroups = computed(() =>
    selectedBuild.value?.itemGroups.filter((group) => group.items.length > 0)
);
</script>

<style lang="scss">
@import '../../css/variables.scss';

.runes-skills {
    border-right: 1px solid $divider-color;
    min-width: 75%;
}
</style>
