<template>
    <div class="max-order">
        <div>{{ $t('skillMaxOrder') }}</div>
        <div class="text-bold">
            <Skill :skill="props.skillMaxOrder[0]" /> →
            <Skill :skill="props.skillMaxOrder[1]" /> →
            <Skill :skill="props.skillMaxOrder[2]" />
        </div>
    </div>
</template>

<style lang="scss">
.max-order {
    font-size: 16px;
}
</style>

<script setup lang="ts">
import Skill from 'components/Lobby/Skill.vue';

interface Props {
    skillMaxOrder: number[];
}

const props = defineProps<Props>();
</script>
