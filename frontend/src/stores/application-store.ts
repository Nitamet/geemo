import { defineStore } from 'pinia';
import { GetState } from 'app/wailsjs/go/main/App';
import { delay } from 'src/delay';
import { lolbuild } from 'app/wailsjs/go/models';
import Build = lolbuild.Build;
import { i18nInstance } from 'boot/i18n';

export enum LeagueState {
    NotLaunched = 'NotLaunched',
    NotInLobby = 'NotInLobby',
    NotSupportedGameMode = 'NotSupportedGameMode',
    InLobby = 'InLobby',
    InGame = 'InGame',
    //  Playing = 'Playing',
    Unknown = 'Unknown',
}

export const useApplicationStore = defineStore('application', {
    state: () => ({
        leagueState: LeagueState.NotLaunched,
        selectedBuild: null as Build | null,
        selectedBuildSource: null as string | null,
        currentChampionId: null as number | null,
    }),
    getters: {
        getLeagueStateMessage(): string {
            switch (this.leagueState) {
                case LeagueState.NotLaunched:
                    return i18nInstance.t('gameNotLaunched');
                case LeagueState.NotSupportedGameMode:
                    return i18nInstance.t('notSupportedGameMode');
                case LeagueState.NotInLobby:
                    return i18nInstance.t('notInLobby');
                case LeagueState.InGame:
                    return i18nInstance.t('inGame');
                case LeagueState.Unknown:
                    return i18nInstance.t('unknownState');
                default:
                    return i18nInstance.t('errorOccurred');
            }
        },
    },
    actions: {
        async startCheckingLeagueState() {
            await delay(1000);

            const newState = await GetState();
            this.leagueState =
                LeagueState[newState as keyof typeof LeagueState] ??
                LeagueState.Unknown;
            await this.startCheckingLeagueState();
        },
    },
});
