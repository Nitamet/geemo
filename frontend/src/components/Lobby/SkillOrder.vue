<template>
    <div class="skill-order" v-if="championData !== null">
        <div class="row no-wrap">
            <q-img class="box" :src="championData.passive.image.full">
                <q-tooltip class="secondary text-body2" :offset="[10, 10]">
                    {{ championData.passive.name }}
                </q-tooltip>
            </q-img>
            <div v-for="index in 18" :key="index">
                <div class="box">{{ index }}</div>
            </div>
        </div>
        <div class="row no-wrap">
            <q-img class="box" :src="championData.spells[0].image.full">
                <q-tooltip class="secondary text-body2" :offset="[10, 10]">
                    {{ championData.spells[0].name }}
                </q-tooltip>
            </q-img>
            <div v-for="index in 18" :key="index">
                <SkillBox
                    :skill-order="props.skillOrder"
                    :skill="1"
                    :index="index"
                />
            </div>
        </div>
        <div class="row no-wrap">
            <q-img class="box" :src="championData.spells[1].image.full">
                <q-tooltip class="secondary text-body2" :offset="[10, 10]">
                    {{ championData.spells[1].name }}
                </q-tooltip>
            </q-img>
            <div v-for="index in 18" :key="index">
                <SkillBox
                    :skill-order="props.skillOrder"
                    :skill="2"
                    :index="index"
                />
            </div>
        </div>
        <div class="row no-wrap">
            <q-img class="box" :src="championData.spells[2].image.full">
                <q-tooltip class="secondary text-body2" :offset="[10, 10]">
                    {{ championData.spells[2].name }}
                </q-tooltip>
            </q-img>
            <div v-for="index in 18" :key="index">
                <SkillBox
                    :skill-order="props.skillOrder"
                    :skill="3"
                    :index="index"
                />
            </div>
        </div>
        <div class="row no-wrap">
            <q-img class="box" :src="championData.spells[3].image.full">
                <q-tooltip class="secondary text-body2" :offset="[10, 10]">
                    {{ championData.spells[3].name }}
                </q-tooltip>
            </q-img>
            <div v-for="index in 18" :key="index">
                <SkillBox
                    :skill-order="props.skillOrder"
                    :skill="4"
                    :index="index"
                />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import SkillBox from 'components/Lobby/SkillBox.vue';
import { GetChampionData } from 'app/wailsjs/go/lolbuild/Loader';
import { storeToRefs } from 'pinia';
import { useSettingsStore } from 'stores/settings-store';
import { useApplicationStore } from 'stores/application-store';
import { Ref, ref, watch } from 'vue';
import { lolbuild } from 'app/wailsjs/go/models';
import ChampionData = lolbuild.ChampionData;

interface Props {
    skillOrder: number[];
}

const props = defineProps<Props>();

const settingsStore = useSettingsStore();
const { language } = storeToRefs(settingsStore);

const applicationStore = useApplicationStore();
const { currentChampionId } = storeToRefs(applicationStore);

const championData: Ref<ChampionData | null> = ref(null);

const getChampionData = async () => {
    if (currentChampionId.value !== null && currentChampionId.value !== -1) {
        championData.value = await GetChampionData(
            currentChampionId.value,
            language.value
        );
    }

    return null;
};

getChampionData();

watch(currentChampionId, async () => {
    await getChampionData();
});
</script>
