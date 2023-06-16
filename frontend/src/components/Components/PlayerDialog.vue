<template>
    <Dialog   modal :draggable="false" :closable="true" header="" showHeader class="w-11" @update:visible="emit('closeDialog')" style="height:fit-content">
    <div class=" grid" v-if="dataLoaded">
            <!-- <div class="grid"> -->
            <div class="ml-6 col" style="font-family: Didot;">
                <span class="text-4xl" >{{ playerAvgs.playerName }}</span>
                <br />
                <span class="text-lg align-self-center">{{ playerAvgs.position }}</span>
                <span class="">
                    
                </span>
            </div>
            <div class="col" style="font-family: Didot;">
                <span v-tooltip.top='{value: getCountryName(playerAvgs.nationality)}' class="inline-block" style="height: fit-content">
                    <CountryFlag size="big" rounded shadow :country="getCountryCode(playerAvgs.nationality)" />
                </span>
                <span>&nbsp; &nbsp;</span>
                <span v-tooltip.top='{value: getCountryName(playerAvgs.secondNationality.String)}' v-if="playerAvgs.secondNationality.Valid == true" class="inline-block" style="height: fit-content">
                    <CountryFlag size="big" rounded shadow :country="getCountryCode(playerAvgs.secondNationality.String)"  />
                </span>
                <!-- <br/> -->
                <span>&nbsp; &nbsp;</span>
                <span class="text-lg inline-block align-self-center vertical-align-super">D.O.B.: {{ playerAvgs.birthdate }}</span>
            </div>
            
            <div class="col flex flex-column align-items-end justify-content-end" style="font-family: Didot;">
                <span class="text-lg">From save: <span class="font-bold text-xl">{{ playerAvgs.saveName }}</span></span>
                <span class="text-lg">Game Version: <span class="font-bold text-xl">{{ playerAvgs.gameVersion }}</span></span>
                <span class="text-lg">Number of seasons: <span class="font-bold text-xl">{{ playerAvgs.seasons }}</span></span>
                
            </div>
        </div>
        
        <div style="font-family: Didot">
            <table class="totalsTable">
                    <thead>
                         <tr class="justify-content-center">
                            <th class="firstCol"></th>
                            <th class="statCell" v-tooltip.top="'Appearances (Subs)'">Apps.</th>
                            <th class="statCell" v-tooltip.top="'Minutes'">Mins.</th>
                            <th class="statCell" v-if="playerAvgs.position != 'GK'" v-tooltip.top="'Goals'">Gls.</th>
                            <th class="statCell" v-if="playerAvgs.position != 'GK'" v-tooltip.top="'Assists'">Asts.</th>
                            <th class="statCell" v-if="playerAvgs.position != 'GK'" v-tooltip.top="'Yellow Cards'">Yell.</th>
                            <th class="statCell" v-if="playerAvgs.position != 'GK'" v-tooltip.top="'Red Cards'">Red</th>
                            <th class="statCell" v-if="playerAvgs.position == 'GK'" >Shutouts</th>
                            <th class="statCell" v-tooltip.top="'Average Rating'">Avg.</th>
                            <th class="statCell" v-tooltip.top="'Player of the Match'">P.o.M.</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr >
                            <td class="firstCol">Total:</td>
                            <td class="statCell">{{ playerAvgs.totStart }} ({{ playerAvgs.totSubs }})</td>
                            <td class="statCell">{{ playerAvgs.totMin }}</td>
                            <td class="statCell" v-if="playerAvgs.position != 'GK'">{{ playerAvgs.totGls }}</td>
                            <td class="statCell" v-if="playerAvgs.position != 'GK'">{{ playerAvgs.totAst }}</td>
                            <td class="statCell" v-if="playerAvgs.position != 'GK'">{{ playerAvgs.totYel }}</td>
                            <td class="statCell" v-if="playerAvgs.position != 'GK'">{{ playerAvgs.totRed }}</td>
                            <td class="statCell" v-if="playerAvgs.position == 'GK'" v-tooltip.top="'Shutouts'">{{ playerAvgs.totShutouts }}</td>
                            <td class="statCell">{{ numberFormatterSQ.format(playerAvgs.avgRat) }}</td>
                            <td class="statCell">{{ playerAvgs.totPOM }}</td>
                        </tr>
                    </tbody>
                </table>
        </div>
    
        <Divider class="fullWidth topDiv" style=""/>
    <!-- </div> -->
    <div class="">
        <TabView >
            <TabPanel header="Season Stats">
                <DataTable :rows="15" scrollHeight="flex" class="p-datatable-sm" :value="playerSeasons" 
                stripedRows reorderableColumns removableSort sortMode="multiple" v-model:expandedRows="expandedSeasons">
                <Column expander style="width: 3rem" />    
                <Column field="season" header="Season" class="min-w-min justify-content-center text-center" style="min-width: 90px!important;" :reorderableColumn="false" sortable></Column>
                    <Column field="teamName" header="Team" class="min-w-min" :reorderableColumn="false"></Column>
                    <Column field="minutes" header="Minutes" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="starts" header="Apps" sortField="starts" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ slotProps.data.starts }} ({{ slotProps.data.subs }})
                        </template>
                    </Column>
                    <Column field="goals" header="Goals" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="assists" header="Assists" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="yellowCards" header="Yellow Cards" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="redCards" header="Red Cards" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="avgRating" header="Avg. Rating" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ numberFormatterSQ.format(slotProps.data.avgRating) }}
                        </template>
                    </Column>
                    <Column field="playerOfTheMatch" header="Player of the Match" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="avgPassP" header="Pass %" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ numberFormatterSQ.format(slotProps.data.passPerc) }}%
                        </template>
                    </Column>
                    <Column field="avgWinP" header="Win %" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ numberFormatterSQ.format(slotProps.data.winPerc) }}%
                        </template>
                    </Column>
                    <template #expansion="slotProps">
                        <div class=" grid flex justify-content-evenly">
                            <Card class="col-3  attrCard" v-if="playerAvgs.position != 'GK'">
                                <template #header class="flex text-center"><h3>Technical</h3></template>
                                <template #content>
                                    <table class="attrTable">
                                        <tbody>
                                            <tr v-for="attr in techAttr">
                                                <td>{{ attr.name }}</td>
                                                <td>{{ slotProps.data[attr.attr] }}</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </template>
                            </Card>
                            <Card class="col-3  attrCard" v-if="playerAvgs.position == 'GK'">
                                <template #header class="flex text-center"><h3>Goalkeeping</h3></template>
                                <template #content>
                                    <table class="attrTable">
                                        <tbody>
                                            <tr v-for="attr in gkAttr">
                                                <td>{{ attr.name }}</td>
                                                <td>{{ slotProps.data[attr.attr] }}</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </template>
                            </Card>
                            <Card class="col-3  attrCard">
                                <template #header ><h3>Mental</h3></template>
                                <template #content>
                                    <table class="attrTable">
                                        <tbody>
                                            <tr v-for="attr in mentalAttr">
                                                <td>{{ attr.name }}</td>
                                                <td>{{ slotProps.data[attr.attr] }}</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </template>
                            </Card>
                            <Card class="col-3 attrCard">
                                <template #header ><h3>Physical</h3></template>
                                <template #content>
                                    <table class="attrTable">
                                        <tbody>
                                            <tr v-for="attr in phyAttr">
                                                <td>{{ attr.name }}</td>
                                                <td>{{ slotProps.data[attr.attr] }}</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </template>
                            </Card>
                        </div>
                    </template>
                </DataTable>
            </TabPanel>
            <TabPanel header="Attribute Progress">

            </TabPanel>
        </TabView>
    </div>
    </Dialog>
    
</template>

<script setup lang="ts">
import { ref, defineProps, defineEmits, onMounted, onBeforeMount } from 'vue'
import { number } from 'yup';
import { GetSinglePlayer } from '../../../wailsjs/go/main/App'

import countryCodes from 'country-code-info';

// @ts-ignore
// import convertFIFACountryCode from '/country-code-converter'
// import getCountryISO2 from "country-iso-3-to-2";

const props = defineProps<{
    playerID: number,
}>()
const emit = defineEmits(
    ['closeDialog', 'beError']
)
const playerID = ref()
const playerSeasons = ref()
const playerAvgs = ref()
const dataLoaded = ref(false)
const expandedSeasons = ref([])
let numberFormatterSQ: Intl.NumberFormat = new Intl.NumberFormat(navigator.language, {
    style: "decimal",
    maximumFractionDigits: 2
})

const techAttr: {attr: string, name: string}[] = [
    {attr: 'cor', name: 'Corners'},
    {attr: 'cro', name: 'Crossing'},
    {attr: 'dri', name: 'Dribbling' },
    {attr: 'fin', name: 'Finishing'},
    {attr: 'fir', name: 'First Touch'},
    {attr: 'fre', name: 'Free Kicks'},
    {attr: 'hea', name: 'Heading'},
    {attr: 'lon', name: 'Long Shots'},
    {attr: 'lth', name: 'Long Throws'},
    {attr: 'mar', name: 'Marking'},
    {attr: 'pas', name: 'Passing'},
    {attr: 'pen', name: 'Penalties'},
    {attr: 'tck', name: 'Tackling'},
    {attr: 'tec', name: 'Technique'}
]

const gkAttr: {attr: string, name: string}[] = [
    {attr: 'aer', name: 'Aerial Reach'},
    {attr: 'cmd', name: 'Command of Area'},
    {attr: 'com', name: 'Communication' },
    {attr: 'ecc', name: 'Eccentricity'},
    {attr: 'fir', name: 'First Touch'},
    {attr: 'han', name: 'Handling'},
    {attr: 'kic', name: 'Kicking'},
    {attr: 'ovo', name: 'One On Ones'},
    {attr: 'pas', name: 'Passing'},
    {attr: 'pun', name: 'Punching (Tendency)'},
    {attr: 'ref', name: 'Reflexes'},
    {attr: 'tro', name: 'Rushing Out (Tendency)'},
    {attr: 'thr', name: 'Throwing'}
]
const mentalAttr: {attr: string, name: string}[] = [
    {attr: 'agg', name: 'Aggression'},
    {attr: 'ant', name: 'Anticipation'},
    {attr: 'bra', name: 'Bravery'},
    {attr: 'cmp', name: 'Composure'},
    {attr: 'cnt', name: 'Concentration'},
    {attr: 'dec', name: 'Decisions'},
    {attr: 'det', name: 'Determination'},
    {attr: 'fla', name: 'Flair'},
    {attr: 'ldr', name: 'Leadership'},
    {attr: 'otb', name: 'Off the Ball'},
    {attr: 'pos', name: 'Positioning'},
    {attr: 'tea', name: 'Teamwork'},
    {attr: 'vis', name: 'Vision'},
    {attr: 'wor', name: 'Work Rate'}
]
const phyAttr: {attr: string, name: string}[] = [
    {attr: 'acc', name: 'Acceleration'},
    {attr: 'agi', name: 'Agility'},
    {attr: 'bal', name: 'Balance'},
    {attr: 'jum', name: 'Jumping Reach'},
    {attr: 'nat', name: 'Natural Fitness'},
    {attr: 'pac', name: 'Pace'},
    {attr: 'sta', name: 'Stamina'},
    {attr: 'str', name: 'Strength'}
]

const difCountryCodes: Map<string, string> = new Map([
    ['ENG', 'GB-ENG'],
    ['WAL', 'GB-WLS'],
    ['SCO', 'GB-SCO'],
    ['NIR', 'GB-NIR'],
    ['KOS', 'XK']
])

const getCountryCode = (nationality: string) => {
    if (difCountryCodes.has(nationality)) {
        return difCountryCodes.get(nationality)
    }
    // return getCountryISO2(nationality)
    return countryCodes.findCountry({'fifa': nationality})!.a3
}

const difCountryNames: Map<string, string> = new Map([
    ['ENG', 'England'],
    ['WAL', 'Wales'],
    ['SCO', 'Scotland'],
    ['NIR', 'Northern Ireland'],
    ['KOS', 'Kosovo']
])

const getCountryName = (nationality: string) => {
    if (difCountryNames.has(nationality)) {
        return difCountryNames.get(nationality)
    }
    return countryCodes.findCountry({'fifa': nationality})!.name
}

onBeforeMount( () => {
    playerID.value = props.playerID
    GetSinglePlayer(props.playerID).then( (response) => {
        dataLoaded.value = true
        if (response.Error != "") {
            emit('beError', response.Error)
            return
        }
        playerSeasons.value = response.OnePlayer
        playerAvgs.value = response.PlayerAvgSum
        // console.log(playerAvgs.value.secondNationality.String)
        //TODO: Add `COUNT(playerSeason.seasonID)` to query where getting the averages from 
        console.log(playerSeasons.value)
    })
})

</script>

<style scoped lang="scss">
    // .totalsTable tbody td:nth-of-type(odd), thead th:nth-of-type(odd) {
    //     background: var(--surface-100);
    // }
    // .totalsTable tbody td:first, thead th:first {
    //     background: none;
    // }
    // .totalsTable td, th {
    //     // border-collapse: collapse;
    //     border: .5px solid grey
    // }

    .firstCol {
        border: none!important
    }

    .statCell {
        padding: 5px;
        text-align: center;
    }

    .totalsTable {
        border-collapse: collapse;
    }

    .totalsTable thead tr {
        border-bottom: 1px solid var(--surface-400);
    }

    .attrCard {
        background: none!important;;
    }

    .attrCard .p-card-header h3 {
        text-align: center;
    }

    .attrCard :deep(.p-card-body) {
        padding-top: 0%!important;
        padding-bottom: 0%!important;
    }

    .attrCard :deep(.p-card-content) {
        padding-top: 0%!important;
        padding-bottom: 0%!important;
    }

    .attrTable{
        width: 100%;
        border-collapse: collapse;
    }
    .attrTable tbody tr {
        border-bottom: 1px solid var(--surface-border);
    }

    .attrTable td {
        font-size: smaller;
    }

</style>