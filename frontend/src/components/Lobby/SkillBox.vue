<template>
    <div :class="['box', 'skill-box']">
        <Skill :skill="props.skill" v-if="maxOnThisLevel" />
    </div>
</template>

<style lang="scss">
@import '../../css/variables.scss';

.skill-box {
    background-color: $build-selection-background-color;
}
</style>

<script setup lang="ts">
import { computed } from 'vue';
import Skill from 'components/Lobby/Skill.vue';

interface Props {
    skillOrder: number[];
    index: number;
    skill: number;
}

const props = defineProps<Props>();

const maxOnThisLevel = computed(
    () => props.skillOrder[props.index - 1] === props.skill
);

const getSkillKey = (skill: number) => {
    switch (skill) {
        case 1:
            return 'Q';
        case 2:
            return 'W';
        case 3:
            return 'E';
        case 4:
            return 'R';
        default:
            return '';
    }
};

const getSkillClass = (skill: number) => {
    switch (skill) {
        case 1:
            return 'skill-q';
        case 2:
            return 'skill-w';
        case 3:
            return 'skill-e';
        case 4:
            return 'skill-r';
        default:
            return '';
    }
};
</script>
