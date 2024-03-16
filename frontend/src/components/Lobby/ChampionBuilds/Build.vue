<template>
    <div
        :class="[
            'build',
            'row',
            'q-gutter-x-sm',
            'items-center',
            'q-py-sm',
            props.isSelected ? 'selected' : '',
        ]"
        @click="
            emit('buildClicked', { build: props.build, source: props.source })
        "
    >
        <q-avatar size="42px" rounded>
            <img :src="getCoreItem().iconUrl" :alt="getCoreItem().name" />
        </q-avatar>
        <div class="build-name column">
            <span class="text-bold">
                {{ props.build.name }}
            </span>
            <span class="text-caption">
                ({{ matches }} {{ $t('matches') }})
            </span>
        </div>
        <div class="q-ml-auto row align-center">
            <span
                :class="['winrate', props.build.winrate > 50 ? 'good' : 'bad']"
            >
                {{ props.build.winrate }}%
            </span>
            <q-img
                class="q-ml-sm"
                width="38px"
                height="38px"
                :src="props.build.selectedPerks[0].iconUrl"
                rounded
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { lolbuild } from 'app/wailsjs/go/models';
import Build = lolbuild.Build;
import { useI18n } from 'vue-i18n';
import { computed } from 'vue';

interface Props {
    build: Build;
    championName: string;
    source: string;
    isSelected: boolean;
}

const props = defineProps<Props>();

const emit = defineEmits<{
    (e: 'buildClicked', value: { build: Build; source: string }): void;
}>();

const i18n = useI18n();
const matches = computed(() =>
    new Intl.NumberFormat(i18n.locale.value).format(props.build.matches)
);

const getCoreItem = () => {
    if ('' !== props.build.coreItem.name) {
        return props.build.coreItem;
    }

    const coreItemGroup = props.build.itemGroups.find(
        (itemGroup) => itemGroup.name === 'Core'
    );

    if (undefined === coreItemGroup) {
        throw new Error('Core item group not found');
    }

    // If we don't have a mythic item, we'll use the second core item
    // Because the first core item is the boots
    return coreItemGroup.items[1];
};
</script>

<style scoped lang="scss">
@import '../../../css/variables.scss';

.build {
    border-radius: 4px;
    transition: background-color 0.1s ease-in-out;
}

.build:hover {
    cursor: pointer;
}

.build-name {
    flex: 1;
}

@media (max-width: $breakpoint-md-max) {
    .build-name {
        font-size: 12px;
    }
}

.selected {
    background-color: $divider-color;
}

.winrate {
    background-color: red;
    border-radius: 4px;
    padding: 6px;
}

.good {
    background-color: #135936;
    color: #17f84a;
}

.bad {
    background-color: #6b1c22;
    color: #ff7984;
}
</style>
