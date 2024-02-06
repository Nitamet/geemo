<template>
    <span :class="getSkillClass(props.skill)">
        {{ getSkillKey(props.skill) }}
    </span>
</template>

<style lang="scss">
.skill-q {
    color: #a04bd5;
}

.skill-w {
    color: #bdbb2f;
}

.skill-e {
    color: #42d782;
}

.skill-r {
    color: cyan;
}
</style>

<script setup lang="ts">
interface Props {
    skill: number;
}

const props = defineProps<Props>();

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
