import { defineStore } from 'pinia';
import { GetLCUState } from 'app/wailsjs/go/main/App';

export enum LeagueState {
    NotLaunched = 'NotLaunched',
    NotInLobby = 'NotInLobby',
    InLobby = 'InLobby',
    Playing = 'Playing',
    Unknown = 'Unknown',
}

export const useApplicationStore = defineStore('application', {
    state: () => ({
        leagueState: LeagueState.NotLaunched,
    }),
    getters: {
        getLeagueStateMessage(): string {
            switch (this.leagueState) {
                case LeagueState.NotLaunched:
                    return 'League of Legends is not launched.';
                case LeagueState.NotInLobby:
                    return 'You are not in lobby.';
                case LeagueState.Unknown:
                    return 'Unknown state.';
                default:
                    return 'An error occurred. Try to restart the app.';
            }
        },
    },
    actions: {
        async startCheckingLeagueState() {
            async function delay(ms: number) {
                return await new Promise((resolve) => setTimeout(resolve, ms));
            }
            await delay(2000);
            console.log('checking');
            const newState = await GetLCUState();
            this.leagueState =
                LeagueState[newState as keyof typeof LeagueState] ??
                LeagueState.Unknown;
            await this.startCheckingLeagueState();
        },
    },
});
