<template>
    <div class="row q-gutter-x-sm items-center q-mt-lg">
        <q-avatar size="42px" rounded>
            <img :src="getCoreItem().iconUrl" :alt="getCoreItem().name" />
        </q-avatar>
        <span>{{ props.build.name }}</span>
        <q-img
            class="q-ml-auto"
            width="38px"
            height="38px"
            :src="`https://ddragon.leagueoflegends.com/cdn/img/${props.build.selectedPerks[0].iconUrl}`"
            rounded
        />
    </div>
</template>

<script setup lang="ts">
import { lolbuild } from 'app/wailsjs/go/models';

import Build = lolbuild.Build;

const props = defineProps<{ build: Build }>();

const getCoreItem = () => {
    if ('' !== props.build.items.mythic.name) {
        return props.build.items.mythic;
    }

    // If we don't have a mythic item, we'll use the second core item
    // Because the first core item is the boots
    return props.build.items.core[1];
};
</script>
