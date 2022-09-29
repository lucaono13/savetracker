import { defineStore } from 'pinia'

export const useSavesStore = defineStore('saves', {
    state: () => ({
        saves: 1,
    }),
    getters: {
        saveCount: (state) => state.saves,
    },
    actions: {
        add() {
            this.saves++
        }
    }
})