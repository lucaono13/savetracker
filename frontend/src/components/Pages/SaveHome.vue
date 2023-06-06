<template>
    <!-- <Sidebar/> -->
    <!-- <div class="col w-full"> -->
        <!-- <div class="grid"> -->
            <!-- <img class="col-2" :src="save.image"  v-if="save.image" /> -->

            <!-- <Card class="col">
                <template #header>

                </template>
                <template #title>
                    {{ save.saveName }}
                </template>
            </Card> -->
            <!-- <div class="col-10 justify-content-center"> -->
                <p>{{ save.saveName }}</p>
                <p>{{ save.managerName }}</p>
            <!-- </div> -->
            <!-- <router-view></router-view> -->
        <!-- </div> -->

    <!-- </div> -->
</template>

<script lang="ts" setup>
import Sidebar from '../Components/Sidebar.vue'
import { useRoute } from 'vue-router';
import { nextTick, ref, onMounted } from 'vue';
import { GetImage, SingleSave, GetSaveResults, GetSaveHomeRankings } from '../../../wailsjs/go/main/App'
import { backend, main } from '../../../wailsjs/go/models'
// export interface Props {
//     id?: number
// }
// const props = withDefaults(defineProps<Props>(), {
//     id: 0
// })

// const props = defineProps({
//     id: String
// })
// var saveID: number | null
// if (props.id == null) {
//     saveID = +localStorage.getItem("defaultSave")
// }
const route = useRoute()
const sortedResults = ref()
const topGls = ref()
const topAsts = ref()
const topApps = ref()
const topRat = ref()
// console.log(route.params.id)
let imgPlaceholder: string | undefined
let save = ref({ saveID: 0, managerName: "", gameVersion: 0, saveName: '', image: imgPlaceholder })
const emit = defineEmits(['beError'])


onMounted( () => {
    SingleSave(+route.params.id).then((response) => {
        if (response.Error != "") {
            emit('beError', response.Error)
            emit('beError','HERE')
            return
        }
        let result = response.Save
        // let sessionVal: string = response.id.toString() + '_save_'
        save.value.saveID = result.id
        save.value.managerName = result.managerName
        save.value.gameVersion = result.gameVersion
        save.value.saveName = result.saveName
        save.value.image = result.saveImage
        // localStorage.setItem("saveCurrency", result.currency)
        if (result.saveImage) {
            GetImage(result.saveImage).then(async (result) => {
                save.value.image = result
            })
        }

        nextTick()
    })
    // GetSaveResults(+route.params.id).then( (response) => {
    GetSaveHomeRankings(+route.params.id).then( (response: main.ErrorReturn) => {
        if (response.Error != "") {
            emit('beError', response.Error)
            emit('beError', "MAYBE")
            return
        }
        topApps.value = response.TopApps
        topAsts.value = response.TopAsts
        topGls.value = response.TopGls
        topRat.value = response.TopAvg
        let teamsMap = new Map<string, {"W": 0, "D": 0, "L": 0, "GF": 0, "GA": 0}>()
        let resultsMap = new Map<string, {"WinPerc": number, "Record": {"W": 0, "D": 0, "L": 0}}>()
        let matches: backend.Match[] = response.Matches
        // teams = Array.from(new Map([...]))

        matches.forEach( function (match) {
            if (!teamsMap.has(match.opposition)) {               
                teamsMap.set(match.opposition, {"W": 0, "D": 0, "L": 0, "GF": 0, "GA": 0})
            }
            let matchResults: {"W": 0, "D": 0, "L": 0, "GF": 0, "GA": 0} = teamsMap.get(match.opposition)!
            switch (match.result) {
                case "W": {
                    matchResults.W += 1
                    break
                }
                case "D": {
                    matchResults.D += 1
                    break
                }
                case "L": {
                    matchResults.L += 1
                    break
                }
            }
            matchResults.GF += match.goalsFor
            matchResults.GA += match.goalsAgainst 
            teamsMap.set(match.opposition, matchResults)
        })

        teamsMap.forEach( (results: {"W": 0, "D": 0, "L": 0, "GF": 0, "GA": 0}, team: string) => {
            let games: number = results.W + results.D + results.L
            if (games > 2) {
                resultsMap.set(team, {"WinPerc": (results.W / games) * 100, "Record": results})
            }
            
        })
        sortedResults.value = [...resultsMap].sort(([k, v], [k2, v2]) => {
            if (v.WinPerc > v2.WinPerc) {
                return 1
            }
            if (v.WinPerc < v2.WinPerc) {
                return -1
            }
            return 0
        })
    })
    
})


    // console.log(SingleSave(+route.params.id))
    // const save = SingleSave(+this.$route.params.id)
</script>