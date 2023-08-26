<template>
    <div class="champion-builds column q-pa-lg">
        <div class="row items-center justify-between">
            <Champion
                :champion-name="currentChampionName"
                :champion-icon-url="currentChampionIconUrl"
            />
            <div>
                <RolePicker
                    v-if="props.gameMode !== GameMode.ARAM"
                    @roleChanged="(newRole) => (selectedRole = newRole)"
                    :assigned-role="selectedRole"
                />
            </div>
        </div>
        <q-separator class="q-mt-md separator" />
        <div class="sources q-mt-sm">
            <div
                v-for="buildCollection in builds"
                :key="buildCollection.source"
            >
                <span class="text-subtitle2 text-bold">
                    {{ buildCollection.source }}
                </span>
                <div class="builds q-gutter-y-md">
                    <Build
                        v-for="build in buildCollection.builds"
                        :key="build.name"
                        :build="build"
                        :champion-name="currentChampionName"
                        :source="buildCollection.source"
                        :is-selected="selectedBuild === build"
                        @build-clicked="
                            (event) => selectBuild(event.build, event.source)
                        "
                    />
                </div>
            </div>
        </div>
        <div class="footer">
            <q-btn
                v-if="selectedBuild !== null"
                class="button q-mt-lg q-mr-lg full-width"
                :label="$t('importSelectedBuild')"
                size="16px"
                color="primary"
                push
                @click="importSelectedBuild()"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import Champion from 'components/Lobby/ChampionBuilds/Champion.vue';
import { delay } from 'src/delay';
import {
    ApplyItemSet,
    ApplyRunes,
    ApplySummonerSpells,
    GetAssignedRole,
    GetCurrentChampion,
} from 'app/wailsjs/go/main/App';
import { computed, onBeforeMount, Ref, ref } from 'vue';
import { whenever } from '@vueuse/core';
import { LeagueState, useApplicationStore } from 'stores/application-store';
import { storeToRefs } from 'pinia';
import { lcu, lolbuild } from 'app/wailsjs/go/models';
import { GetChampionName, LoadBuilds } from 'app/wailsjs/go/lolbuild/Loader';
import Build from 'components/Lobby/ChampionBuilds/Build.vue';
import { GameMode, Role } from 'components/models';
import RolePicker from 'components/RolePicker.vue';
import { useSettingsStore } from 'stores/settings-store';
import BuildCollection = lolbuild.BuildCollection;
import BuildInfo = lolbuild.Build;
import ItemSet = lcu.ItemSet;

interface Props {
    gameMode: GameMode;
}

const settingsStore = useSettingsStore();
const { language, activeSources } = storeToRefs(settingsStore);

const props = defineProps<Props>();
const championNone = -1;

let currentChampion = ref(championNone);
let currentChampionName = ref('Champion');

const currentChampionIconUrl = computed(() => {
    return `https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/${currentChampion.value}.png`;
});

const buildCollections = new Map<string, BuildCollection[]>();

const loadBuildCollection = async () => {
    if (currentChampion.value === championNone) {
        return [];
    }

    const championName = await GetChampionName(
        currentChampion.value,
        language.value
    );
    currentChampionName.value = championName.name;

    const buildCollection = (
        await LoadBuilds(
            championName.slug,
            activeSources.value,
            selectedRole.value,
            language.value
        )
    ).sort((a, b) => a.source.localeCompare(b.source));

    buildCollections.set(
        `${currentChampion.value}-${selectedRole.value}`,
        buildCollection
    );

    return buildCollection;
};

let selectedRole = ref<Role>(Role.Mid);

const loadBuilds = async () => {
    const buildCollection = buildCollections.get(
        `${currentChampion.value}-${selectedRole.value}`
    );

    if (!buildCollection) {
        builds.value = await loadBuildCollection();
    } else {
        builds.value = buildCollection;
    }

    if (builds.value.length > 0) {
        selectBuild(builds.value[0].builds[0], builds.value[0].source);
    }
};

whenever(selectedRole, async () => {
    if (currentChampion.value === -1) {
        return;
    }

    await loadBuilds();
});

const application = useApplicationStore();
const { leagueState, selectedBuild, selectedBuildSource } =
    storeToRefs(application);
const startCheckingCurrentChampion = async () => {
    await delay(3000);
    const champion = await GetCurrentChampion();
    if (champion !== 0) {
        currentChampion.value = champion;
    }

    if (leagueState.value === LeagueState.InLobby) {
        await startCheckingCurrentChampion();
    }
};
startCheckingCurrentChampion();

let builds: Ref<BuildCollection[]> = ref([]);

whenever(currentChampion, async () => {
    await loadBuilds();
});

whenever(leagueState, () => {
    // Clear the selected build when the game starts or the lobby is closed
    if (leagueState.value !== LeagueState.InLobby) {
        selectedBuild.value = null;
    }
});

whenever(language, async () => {
    buildCollections.clear();
    selectedBuild.value = null;
    await loadBuilds();
});

whenever(activeSources, async () => {
    buildCollections.clear();
    selectedBuild.value = null;
    await loadBuilds();
});

const selectBuild = (build: BuildInfo, source: string) => {
    selectedBuild.value = build;
    selectedBuildSource.value = source;

    if (settingsStore.autoImport) {
        importSelectedBuild();
    }
};

const getItemSet = (build: lolbuild.Build) => {
    const itemBlocks: {
        type: string;
        items: { id: string; count: number }[];
    }[] = [];
    for (const itemGroup of build.itemGroups) {
        itemBlocks.push({
            type: itemGroup.name,
            items: itemGroup.items.map((item) => ({
                id: item.id.toString(),
                count: 1,
            })),
        });
    }

    return {
        title: `${currentChampionName.value}: ${build.name}`,
        associatedChampions: [],
        associatedMaps: [],
        type: 'custom',
        map: 'any',
        mode: 'any',
        startedFrom: 'blank',
        uid: '1',
        preferredItemSlots: [],
        sortrank: 0,
        blocks: itemBlocks,
    };
};

const importBuild = (build: BuildInfo, source: string) => {
    const selectedPerks = build.selectedPerks.map((perk) => perk.id);

    const runePage = {
        name: `${source}: ${build.name} ${currentChampionName.value}`,
        primaryStyleId: build.primary.id,
        selectedPerkIds: selectedPerks,
        subStyleId: build.secondary.id,
        current: true,
    };

    ApplyRunes(runePage);

    const summonerSpells = {
        firstSpellId: build.summonerSpells.at(0)?.id ?? 0,
        secondSpellId: build.summonerSpells.at(1)?.id ?? 0,
    };

    ApplySummonerSpells(
        summonerSpells.firstSpellId,
        summonerSpells.secondSpellId
    );

    const itemSetRaw = getItemSet(build);

    ApplyItemSet(ItemSet.createFrom(itemSetRaw));
};

const importSelectedBuild = () => {
    if (selectedBuild.value) {
        importBuild(
            selectedBuild.value,
            selectedBuildSource.value ?? 'Unknown'
        );
    }
};

onBeforeMount(async () => {
    const leagueAssignedRole = await GetAssignedRole();
    if ('' !== leagueAssignedRole) {
        selectedRole.value = leagueAssignedRole as Role;
    } else {
        selectedRole.value =
            GameMode.ARAM === props.gameMode ? Role.ARAM : Role.Mid;
    }
});
</script>

<style lang="scss">
@import '../../css/variables.scss';

.champion-builds {
    width: 30%;
    background-color: $build-selection-background-color;
    border-right: 2px solid $divider-color;
}

.footer {
    margin: auto 0 0;
}
</style>
