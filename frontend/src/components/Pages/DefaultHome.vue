<template>
    <div class="pr-4 pb-4" >
        <div class="mt-3 ml-1">
            <span class="text-5xl" style="font-family: Didot;">Welcome to the Save Tracker</span>
        </div>
        <Divider class="fullWidth topDiv" style=""/>
        <NoDefaultSave v-if="numOfSaves > 0 && numOfSeasons <= 0"/>
        <NoSaves v-if="numOfSaves <= 0" />
        <div class="grid fullWidth mr-0 gap-0" v-if="numOfSaves > 0 && numOfSeasons > 0">
            <div v-if="trophies != null" class="col-12" > 
                        <Carousel :value="trophies" circular :numVisible="5" :numScroll="1">
                            <template #item="slotProps">
                                <div class="border-2  border-round border-yellow-900 m-2 text-center py-1 px-1 flex flex-column align-items-center justify-content-center" >
                                    <div class="">
                                        <Button label="" class="trophyImgButton" @click="newTrophyImage(slotProps.data)">
                                            <img :src="slotProps.data[1]['b64']"  style="max-height: 100px!important;"/>
                                            <FontAwesomeIcon icon="trophy"  style="color: var(--text-color); height: 100px!important;" v-if="!slotProps.data[1]['b64']"/>
                                        </Button>                                        
                                    </div>
                                    <div >
                                        {{ slotProps.data[0] }}<br/>X{{ slotProps.data[1]["years"].length }}
                                    </div>
                                </div>
                            </template>
                        </Carousel> 
            </div> 
            <div class="col-4 statsDiv">
                <Card class="statsCard" style="max-height: 250px!important;">
                    <template #header><h1 class="flex align-content-center justify-content-center my-0">Transfers</h1></template>
                    <template #content>
                        <TabView class="playerStatView">
                            <TabPanel header="Transfers">
                                <table class="statTable">
                                    <thead>
                                        <tr>
                                            <th>Team</th>
                                            <th>Count (Total)</th>
                                        </tr>
                                    </thead>
                                    <tr v-for="team in mostTrfs">
                                        <td>{{ team.teamName }}</td>
                                        <td class="transfer">{{ team.numTransfers }} ({{ numberFormmaterCur.format(team.totFee) }})</td>
                                    </tr>
                                </table>
                            </TabPanel>
                            <TabPanel header="Loans">
                                <table class="statTable">
                                    <thead>
                                        <tr>
                                            <th>Team</th>
                                            <th>Count (Total)</th>
                                        </tr>
                                    </thead>
                                    <tr v-for="team in mostLoans">
                                        <td>{{ team.teamName }}</td>
                                        <td class="transfer">{{ team.numLoans }} ({{ numberFormmaterCur.format(team.totFee) }})</td>
                                    </tr>
                                </table>
                            </TabPanel>
                        </TabView>
                        
                    </template>
                </Card>
            </div>
            <div class="col-4 statsDiv">
                <Card class="statsCard">
                    <template #header><h1 class="flex align-content-center justify-content-center my-0">Results</h1></template>
                    <template #content>
                        <TabView class="playerStatView">
                            <TabPanel header="Best">
                                <table class="statTable">
                                    <thead>
                                        <tr>
                                            <th>Team</th>
                                            <th>Win %</th>
                                            <th>Record</th>
                                        </tr>
                                    </thead>
                                    <tr v-for="team in sortedResults_best">
                                        <td>{{ team[0] }}</td>
                                        <td>{{ numberFormmaterDec.format(team[1].WinPerc) }}%</td>
                                        <td >{{ team[1].Record.W }}-{{ team[1].Record.D }}-{{ team[1].Record.L }}</td>
                                    </tr>
                                </table>
                            </TabPanel>
                            <TabPanel header="Worst">
                                <table class="statTable">
                                    <thead>
                                        <tr>
                                            <th>Team</th>
                                            <th>Win %</th>
                                            <th>Record</th>
                                        </tr>
                                    </thead>
                                    <tr v-for="team in sortedResults_worst">
                                        <td>{{ team[0] }}</td>
                                        <td>{{ numberFormmaterDec.format(team[1].WinPerc) }}%</td>
                                        <td >{{ team[1].Record.W }}-{{ team[1].Record.D }}-{{ team[1].Record.L }}</td>
                                    </tr>
                                </table>
                            </TabPanel>
                        </TabView>
                    </template>
                </Card>
            </div>
            <div class="col-4 statsDiv" style="max-height:250px!important ;" v-if="numPlayerSeasons > 0">
                <Card class="statsCard" style="max-height: 250px!important;">
                    <template #header ><h1 class="flex align-content-center justify-content-center my-0">Player Stats</h1></template>
                    <!-- <h1 class="flex align-content-center justify-content-center my-0">Player Stats</h1> -->
                    <template #content class="pt-0" style="max-height: 250px!important;padding-top: 0px!important;">
                        <TabView class="playerStatView" style="height: 207px!important;" >
                            <TabPanel header="Gls.">
                                <table class="statTable">
                                    <tr v-for="player in topGls">
                                        <td><Button class="playerButton" :label="player.playerName" link @click="openPlayerDialog(player.playerID)"/></td>
                                        <td class="stat">{{ player.goals }}</td>
                                    </tr>
                                </table>
                            </TabPanel>
                            <TabPanel header="Ast.">
                                <table class="statTable">
                                    <tr v-for="player in topAsts">
                                        <td><Button class="playerButton" :label="player.playerName" link @click="openPlayerDialog(player.playerID)"/></td>
                                        <td class="stat">{{ player.assists }}</td>
                                    </tr>
                                </table>
                            </TabPanel>
                            <TabPanel header="Apps">
                                <table class="statTable">
                                    <tr v-for="player in topApps">
                                        <td><Button class="playerButton" :label="player.playerName" link @click="openPlayerDialog(player.playerID)"/></td>
                                        <td class="stat">{{ player.apps }}</td>
                                    </tr>
                                </table>
                            </TabPanel>
                            <TabPanel header="Rating">
                                <table class="statTable">
                                    <tr v-for="player in topRat">
                                        <td><Button class="playerButton" :label="player.playerName" link @click="openPlayerDialog(player.playerID)"/></td>
                                        <td class="stat">{{ numberFormmaterDec.format(player.avgRating) }}</td>
                                    </tr>
                                </table>
                            </TabPanel>
                        </TabView>
                        
                    </template>
                </Card>
            </div>
        </div>
    </div>
    <PlayerDialog :visible="playerDialog" v-if="playerDialog" :playerID="playerDialogID" @closeDialog="playerDialog = false"/>
</template>

<script lang="ts" async setup>
import { ref, onMounted } from 'vue';
import { GetImage, GetAllRankings, SelectNewTrophyImage, GetNumSaves, GetNumSeasons, GetAllPSeasons } from '../../../wailsjs/go/main/App'
import PlayerDialog from '../Components/PlayerDialog.vue';
import NoDefaultSave from './NoDefaultSave.vue';
import NoSaves from './NoSaves.vue';
import { backend, main } from '../../../wailsjs/go/models'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { useRouter } from 'vue-router';

const playerDialog = ref(false)
const playerDialogID = ref(0)
const sortedResults_worst = ref()
const sortedResults_best = ref()
const topGls = ref()
const topAsts = ref()
const topApps = ref()
const topRat = ref()
const mostTrfs = ref()
const mostLoans = ref()
const avgInFee = ref()
const avgOutFee = ref()
const trophies = ref()
const numOfSaves = ref()
const numOfSeasons = ref()
const numPlayerSeasons = ref()

const router = useRouter()
let imgPlaceholder: string | undefined
let save = ref({ saveID: 0, managerName: "", gameVersion: 0, saveName: '', image: imgPlaceholder })
const emit = defineEmits(['beError'])
let numberFormmaterDec: Intl.NumberFormat = new Intl.NumberFormat(navigator.language, {
    style: "decimal",
    maximumFractionDigits: 2
})
let numberFormmaterCur: Intl.NumberFormat

const beError = (error: string) => {
    emit('beError', error)
}

function openPlayerDialog(playerID: number) {
    playerDialogID.value = playerID
    playerDialog.value = true
}

function newTrophyImage(trophyData: {0: string, 1: {"id": number, "image": string, "years": string[], "b64": string}}) {
    SelectNewTrophyImage(trophyData[1].id).then( (response: main.ErrorReturn) => {
        if (response.Error) {
            emit('beError', response.Error)
            return
        }
        trophies.value.forEach( function (trophy: {0: string, 1: {"id": number, "image": string, "years": string[], "b64": string}}) {
            if (trophy[1].id == trophyData[1].id) {
                trophy[1].image = response.ImageFile
            }
        })
        GetImage(response.ImageFile).then( (response) => {
            if (response.Error != "") {
                emit('beError', response.Error)
                return
            }
            trophyData[1].b64 = response.b64Image
        })
    })
}

onMounted(  async () => {
    GetNumSaves().then( (totSaves: number) => {
        numOfSaves.value = totSaves
            if (totSaves <= 0){
                router.replace({name: 'No Saves', replace: true})
                return
            }
        GetNumSeasons().then( (totSeasons: number) => {
            numOfSeasons.value = totSeasons
            // console.log(numOfSeasons.value)
            if (totSeasons <= 0) {
                router.replace({name: 'No Seasons',replace: true})
                return
            }
            GetAllPSeasons().then( (playerSeasonNum: number) => {
                numPlayerSeasons.value = playerSeasonNum
                // console.log(numPlayerSeasons.value > 0)
                
                GetAllRankings(numPlayerSeasons.value > 0).then( (response: main.ErrorReturn) => {
                    if (response.Error != "") {
                        emit('beError', response.Error)
                        return
                    }
                    topApps.value = response.TopApps
                    topAsts.value = response.TopAsts
                    topGls.value = response.TopGls
                    topRat.value = response.TopAvg
                    mostTrfs.value = response.TopTrfs
                    mostLoans.value = response.TopLoans
                    avgInFee.value = response.AvgInFee
                    avgOutFee.value = response.AvgOutFee
                    numberFormmaterCur = new Intl.NumberFormat(navigator.language, {
                        style: "currency",
                        currency: response.TopTrfs[0].currency,
                        notation: "compact"
                    })
                    let teamsMap = new Map<string, {"W": 0, "D": 0, "L": 0, "GF": 0, "GA": 0}>()
                    let resultsMap = new Map<string, {"WinPerc": number, "Record": {"W": 0, "D": 0, "L": 0}}>()
                    let matches: backend.Match[] = response.Matches

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
                        if (games >= 2) {
                            resultsMap.set(team, {"WinPerc": (results.W / games) * 100, "Record": results})
                        }
                        
                    })
                    sortedResults_worst.value = [...resultsMap].sort(([k, v], [k2, v2]) => {
                        if (v.WinPerc > v2.WinPerc) {
                            return 1
                        }
                        if (v.WinPerc < v2.WinPerc) {
                            return -1
                        }
                        if (v.WinPerc == v2.WinPerc) {
                            if (v.Record.L > v2.Record.L) {
                                return 1
                            }
                            if (v.Record.L < v2.Record.L) {
                                return -1
                            }
                            if (v.Record.L == v2.Record.L) {
                                if (v.Record.D > v2.Record.D) {
                                    return 1
                                }
                                if (v.Record.D < v2.Record.D) {
                                    return -1
                                } 
                                if (v.Record.D == v2.Record.D) {
                                    if (v.Record.W < v2.Record.W) {
                                        return -1
                                    }
                                    if (v.Record.W > v2.Record.W) {
                                        return 1
                                    }
                                }
                            }
                        }
                        return 0
                    })
                    sortedResults_best.value = [...resultsMap].sort(([k, v], [k2, v2]) => {
                        if (v.WinPerc < v2.WinPerc) {
                            return 1
                        }
                        if (v.WinPerc > v2.WinPerc) {
                            return -1
                        }
                        if (v.WinPerc == v2.WinPerc) {
                            if (v.Record.W < v2.Record.W) {
                                return 1
                            }
                            if (v.Record.W > v2.Record.W) {
                                return -1
                            }
                            if (v.Record.W == v2.Record.W) {
                                if (v.Record.D < v2.Record.D) {
                                    return 1
                                }
                                if (v.Record.D > v2.Record.D) {
                                    return -1
                                } 
                                if (v.Record.D == v2.Record.D) {
                                    if (v.Record.L > v2.Record.L) {
                                        return -1
                                    }
                                    if (v.Record.L < v2.Record.L) {
                                        return 1
                                    }
                                }
                            }
                        }
                        return 0
                    })
                    sortedResults_worst.value = sortedResults_worst.value.slice(0,5)
                    sortedResults_best.value = sortedResults_best.value.slice(0,5)
                    if (response.Trophies == null) {
                        return
                    }
                    let combinedTrophies = new Map<string, {"years": string[], "image": string, "id": number, "b64": string}>()
                    response.Trophies.forEach( function (trophy) {
                        if (!combinedTrophies.has(trophy.trophyName)) {
                            combinedTrophies.set(trophy.trophyName, {"years": [trophy.season], "image": trophy.trophyImage, "id": trophy.trophyID, "b64": ""})
                        } else if (combinedTrophies.has(trophy.trophyName)) {
                            let trophyInfo: {"years":string[], "image": string, "id": number, "b64": string} = combinedTrophies.get(trophy.trophyName)!
                            trophyInfo.years.push(trophy.season)
                            combinedTrophies.set(trophy.trophyName, trophyInfo)
                        }
                    })
                    trophies.value = [...combinedTrophies]
                    
                    trophies.value.forEach( function (trophy: {"years": string[], "image": string, "id": number, "b64": string}[]) {
                        if (trophy[1]['image'] != null) {
                            GetImage(trophy[1]['image']).then( (response) => {
                                if (response.Error != "") {
                                    emit('beError', response.Error)
                                    return
                                }
                                trophy[1]['b64'] = response.b64Image
                            })
                        }
                    })
                })
            })
        })
    })
    
})

</script>

<style scoped lang="scss">

.comp-size {
        height: calc(100vh - 79px)!important;
        width: calc(100vw - 205px - 20px)!important;
    }

.fullWidth {
    // height: calc(100vw - 79px)!important;
    width: calc(100vw - 205px - 50px)!important;
}

.isHovered {
    background-color:hotpink!important;
}

.story p:first-of-type {
    margin-top: 0px!important;
}
.statsCard {
    background: none!important;
    box-shadow: none!important;
    color: white!important
}

.statTable {
    width: 100%!important;
    // border: 1px solid white;
    border-collapse: collapse;
}

.stat {
    width: 3rem!important;
}

.transfer {
    width: 6rem!important;
}

.statTable tr {
    border-bottom: 1px solid var(--surface-200);
}

.playerStatView .p-tabview-panels, .p-tabview-nav, .p-tabview-nav-link {
    background: none!important;
    border: 0px!important;
    padding-top: 0px!important;
}

.p-tabview-panel {
    padding: 0px!important;
    padding-top: 5px!important;
}

.p-tabview-panels {
    background: none!important;
}

:deep(.p-tabview-panels) {
    padding:0px!important;
    background: none!important;
    padding-top: .5rem!important;
}

.statsCard :deep(.p-card-body), :deep(.p-card-content) {
    padding-top: 0px!important;
}

.col-4 {
    padding: 0.25rem!important
}

.p-tabview :deep(.p-tabview-nav li.p-highlight .p-tabview-nav-link) {
    color: #5eead4
}
.p-tabview :deep(.p-tabview-nav li .p-tabview-nav-link) {
    background: transparent!important;
    color:white;
    padding: 10px!important
}

.p-carousel {
    height: 100%!important;
} 

.trophyImgButton {
    background-color: transparent!important;
    border: none!important
    
}

.trophyImgButton:click{
    border:none!important
}

.trophyImgButton:hover {
    background-color: grey!important;
}

.editButton {
    height: 1rem!important;
}

.playerButton {
    padding: 0%!important;
    color: var(--surface-900)!important
}

</style>