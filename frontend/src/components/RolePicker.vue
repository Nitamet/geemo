<template>
    <q-btn-dropdown
        color="primary"
        :icon="`img:${roleIcons[selectedRole]}`"
        class="role-picker"
        content-class="role-picker-menu"
    >
        <q-list>
            <q-item clickable v-close-popup @click="changeRole(Role.Top)">
                <q-item-section avatar>
                    <q-img :src="roleIcons['top']" />
                </q-item-section>
                <q-item-section>
                    <q-item-label>Top</q-item-label>
                </q-item-section>
            </q-item>

            <q-item clickable v-close-popup @click="changeRole(Role.Jungle)">
                <q-item-section avatar>
                    <q-img :src="roleIcons['jungle']" />
                </q-item-section>
                <q-item-section>
                    <q-item-label>Jungle</q-item-label>
                </q-item-section>
            </q-item>

            <q-item clickable v-close-popup @click="changeRole(Role.Mid)">
                <q-item-section avatar>
                    <q-img :src="roleIcons['mid']" />
                </q-item-section>
                <q-item-section>
                    <q-item-label>Middle</q-item-label>
                </q-item-section>
            </q-item>

            <q-item clickable v-close-popup @click="changeRole(Role.ADC)">
                <q-item-section avatar>
                    <q-img :src="roleIcons['adc']" />
                </q-item-section>
                <q-item-section>
                    <q-item-label>ADC</q-item-label>
                </q-item-section>
            </q-item>

            <q-item clickable v-close-popup @click="changeRole(Role.Support)">
                <q-item-section avatar>
                    <q-img :src="roleIcons['support']" />
                </q-item-section>
                <q-item-section>
                    <q-item-label>Support</q-item-label>
                </q-item-section>
            </q-item>
        </q-list>
    </q-btn-dropdown>
</template>

<script setup lang="ts">
import { Role } from 'components/models';
import { ref } from 'vue';

const emit = defineEmits<{
    (e: 'roleChanged', value: Role): void;
}>();

let selectedRole = ref(Role.Mid);

const changeRole = (role: Role) => {
    selectedRole.value = role;
    emit('roleChanged', role);
};

const roleIcons = {
    top: 'https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-clash/global/default/assets/images/position-selector/positions/icon-position-top.png',
    jungle: 'https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-clash/global/default/assets/images/position-selector/positions/icon-position-jungle.png',
    mid: 'https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-clash/global/default/assets/images/position-selector/positions/icon-position-middle.png',
    adc: 'https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-clash/global/default/assets/images/position-selector/positions/icon-position-bottom.png',
    support:
        'https://raw.communitydragon.org/latest/plugins/rcp-fe-lol-clash/global/default/assets/images/position-selector/positions/icon-position-utility.png',
};
</script>
