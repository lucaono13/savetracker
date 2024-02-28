<template>
    <Dialog modal :draggable="false" :closable="true" header="" showHeader class="w-11" @update:visible="emit('closeDialog')" style="height:fit-content">
    <div class=" grid" v-if="dataLoaded">
            <!-- <div class="grid"> -->
            <div class="ml-6 col" style="font-family: Didot;">
                <span class="text-4xl" >{{ playerAvgs.playerName }}</span>
                <br />
                <span class="text-lg align-self-center">{{ playerAvgs.position }}</span>
                <span class="">
                    
                </span>
            </div>
            <div class="col align-items-start" style="font-family: Didot;">
                <span v-tooltip.top='{value: getCountryName(playerAvgs.nationality)}' class="inline-block" style="height:fit-content" v-if="playerAvgs.secondNationality.Valid == false">
                    <CountryFlag size="big" rounded shadow :country="getCountryCode(playerAvgs.nationality)" />
                </span>
                <Splitter layout="vertical" :gutter-size="0" v-if="playerAvgs.secondNationality.Valid == true" class="inline-block justify-content-center" style="width: 55px">
                    <SplitterPanel v-tooltip.top='{value: getCountryName(playerAvgs.nationality)}'>
                        <CountryFlag size="big" rounded shadow :country="getCountryCode(playerAvgs.nationality)" />
                    </SplitterPanel>
                    <SplitterPanel v-tooltip.top='{value: getCountryName(playerAvgs.secondNationality.String)}'>
                        <CountryFlag size="big" rounded shadow :country="getCountryCode(playerAvgs.secondNationality.String)" />
                    </SplitterPanel>
                </Splitter>
                <span>&nbsp; &nbsp;</span>
                <div class="text-lg inline align-items-start vertical-align-top" style=""><span style="vertical-align: sub;">D.O.B.: {{ playerAvgs.birthdate }}</span></div>
            </div>
            
            <div class="col flex flex-column align-items-end justify-content-end" style="font-family: Didot;">
                <span class="text-lg">From save: <span class="font-bold text-xl">{{ playerAvgs.saveName }}</span></span>
                <span class="text-lg">Game Version/Save Type: <span class="font-bold text-xl">{{ playerAvgs.gameVersion }}</span></span>
                <span class="text-lg">Number of seasons: <span class="font-bold text-xl">{{ playerAvgs.seasons }}</span></span>
                
            </div>
        </div>
        <div style="font-family: Didot" v-if="dataLoaded">
            <table class="totalsTable">
                    <thead>
                         <tr class="justify-content-center">
                            <th class="firstCol"></th>
                            <th class="statCell" v-tooltip.top="'Appearances (Subs)'">Apps.</th>
                            <th class="statCell" v-tooltip.top="'Minutes'">Mins.</th>
                            <th class="statCell" v-if="!isGK" v-tooltip.top="'Goals'">Gls.</th>
                            <th class="statCell" v-if="!isGK" v-tooltip.top="'Assists'">Asts.</th>
                            <th class="statCell" v-if="!isGK" v-tooltip.top="'Yellow Cards'">Yell.</th>
                            <th class="statCell" v-if="!isGK" v-tooltip.top="'Red Cards'">Red</th>
                            <th class="statCell" v-if="isGK" >Shutouts</th>
                            <th class="statCell" v-tooltip.top="'Average Rating'">Avg.</th>
                            <th class="statCell" v-tooltip.top="'Player of the Match'">P.o.M.</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr >
                            <td class="firstCol">Total:</td>
                            <td class="statCell">{{ playerAvgs.totStart }} ({{ playerAvgs.totSubs }})</td>
                            <td class="statCell">{{ playerAvgs.totMin }}</td>
                            <td class="statCell" v-if="!isGK">{{ playerAvgs.totGls }}</td>
                            <td class="statCell" v-if="!isGK">{{ playerAvgs.totAst }}</td>
                            <td class="statCell" v-if="!isGK">{{ playerAvgs.totYel }}</td>
                            <td class="statCell" v-if="!isGK">{{ playerAvgs.totRed }}</td>
                            <td class="statCell" v-if="isGK" v-tooltip.top="'Shutouts'">{{ playerAvgs.totShutouts }}</td>
                            <td class="statCell">{{ numberFormatterSQ.format(playerAvgs.avgRat) }}</td>
                            <td class="statCell">{{ playerAvgs.totPOM }}</td>
                        </tr>
                    </tbody>
                </table>
        </div>
    
        <Divider class="fullWidth topDiv" style=""/>
    <!-- </div> -->
    <div class="" v-if="dataLoaded">
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
                    <Column v-if="!isGK" field="goals" header="Goals" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column v-if="!isGK" field="assists" header="Assists" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column v-if="!isGK" field="yellowCards" header="Yellow Cards" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column v-if="!isGK" field="redCards" header="Red Cards" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column v-if="isGK" field="shutouts" header="Shutouts" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column v-if="isGK" field="savePerc" header="Save %" class="min-w-min justify-content-center text-center" sortable></Column>
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
                            <Card class="col-3  attrCard" v-if="!isGK">
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
                            <Card class="col-3  attrCard" v-if="isGK">
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
                <div class="flex flex-column justify-content-center">
                    
                    <Chart id="attrChart" type="line" :data="chartData" :options="chartOptions"  class="h-25rem" />
                    <div class="w-full flex justify-content-center"><Button class="w-4" text raised label="Choose Attributes to Show" @click="toggle" /></div>
                    <OverlayPanel ref="attrChoices" :showCloseIcon="true">
                        <div class="flex ">
                            <Card v-if="!isGK" class="attrCard selectionCard">
                                <template #header><h3>Technical</h3></template>
                                <template #content>
                                    <table class="attrTable ">
                                        <tbody>
                                            <tr v-for="attr in techAttr">
                                                <td>{{ attr.name }}</td>
                                                <td><InputSwitch v-model="attr.visible.value" @change="updateChartData"/></td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </template>
                            </Card>
                            <Card v-if="isGK" class="attrCard selectionCard">
                                <template #header><h3>Goalkeeping</h3></template>
                                <template #content>
                                    <table class="attrTable">
                                        <tbody>
                                            <tr v-for="attr in gkAttr">
                                                <td>{{ attr.name }}</td>
                                                <td><InputSwitch v-model="attr.visible.value" @change="updateChartData"/></td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </template>
                            </Card>
                            <Card class="attrCard selectionCard">
                                <template #header><h3>Mental</h3></template>
                                <template #content>
                                    <table class="attrTable">
                                        <tbody>
                                            <tr v-for="attr in mentalAttr">
                                                <td>{{ attr.name }}</td>
                                                <td><InputSwitch v-model="attr.visible.value" @change="updateChartData"/></td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </template>
                            </Card>
                            <Card  class="attrCard selectionCard">
                                <template #header><h3>Physical</h3></template>
                                <template #content>
                                    <table class="attrTable ">
                                        <tbody>
                                            <tr v-for="attr in phyAttr">
                                                <td>{{ attr.name }}</td>
                                                <td><InputSwitch v-model="attr.visible.value" @change="updateChartData"/></td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </template>
                            </Card>
                        </div>
                    </OverlayPanel>
                    
                </div>
            </TabPanel>
        </TabView>
    </div>
    </Dialog>
    
</template>

<script setup lang="ts">
import { ref, defineProps, defineEmits, onBeforeMount, Ref } from 'vue'
import { GetSinglePlayer } from '../../../wailsjs/go/main/App'
import { backend } from '../../../wailsjs/go/models';
import countryCodes from 'country-code-info';


const props = defineProps<{
    playerID: number,
}>()
const emit = defineEmits(
    ['closeDialog', 'beError']
)
const documentStyle = document.styleSheets[0].cssRules[0]
const playerID = ref()
const playerSeasons = ref()
const playerAvgs = ref()
const dataLoaded = ref(false)
const expandedSeasons = ref([])
const chartOptions = ref()
const chartData = ref()
const chartPlugins = ref()
const seasons = ref()
const isGK = ref(false)
const attrChoices = ref()
let numberFormatterSQ: Intl.NumberFormat = new Intl.NumberFormat(navigator.language, {
    style: "decimal",
    maximumFractionDigits: 2
})

const toggle = (event: any) => {
    attrChoices.value.toggle(event)
}

const techAttr: {attr: string, name: string, visible: Ref<boolean>}[] = [
    {attr: 'cor', name: 'Corners', visible: ref(false)},
    {attr: 'cro', name: 'Crossing', visible: ref(false)},
    {attr: 'dri', name: 'Dribbling', visible: ref(false)},
    {attr: 'fin', name: 'Finishing', visible: ref(false)},
    {attr: 'fir', name: 'First Touch', visible: ref(false)},
    {attr: 'fre', name: 'Free Kicks', visible: ref(false)},
    {attr: 'hea', name: 'Heading', visible: ref(false)},
    {attr: 'lon', name: 'Long Shots', visible: ref(false)},
    {attr: 'lth', name: 'Long Throws', visible: ref(false)},
    {attr: 'mar', name: 'Marking', visible: ref(false)},
    {attr: 'pas', name: 'Passing', visible: ref(false)},
    {attr: 'pen', name: 'Penalties', visible: ref(false)},
    {attr: 'tck', name: 'Tackling', visible: ref(false)},
    {attr: 'tec', name: 'Technique', visible: ref(false)}
]
const gkAttr: {attr: string, name: string, visible: Ref<boolean>}[] = [
    {attr: 'aer', name: 'Aerial Reach', visible: ref(false)},
    {attr: 'cmd', name: 'Command of Area', visible: ref(false)},
    {attr: 'com', name: 'Communication', visible: ref(false)},
    {attr: 'ecc', name: 'Eccentricity', visible: ref(false)},
    {attr: 'fir', name: 'First Touch', visible: ref(false)},
    {attr: 'han', name: 'Handling', visible: ref(false)},
    {attr: 'kic', name: 'Kicking', visible: ref(false)},
    {attr: 'ovo', name: 'One On Ones', visible: ref(false)},
    {attr: 'pas', name: 'Passing', visible: ref(false)},
    {attr: 'pun', name: 'Punching (Tendency)', visible: ref(false)},
    {attr: 'ref', name: 'Reflexes', visible: ref(false)},
    {attr: 'tro', name: 'Rushing Out (Tendency)', visible: ref(false)},
    {attr: 'thr', name: 'Throwing', visible: ref(false)}
]
const mentalAttr: {attr: string, name: string, visible: Ref<boolean>}[] = [
    {attr: 'agg', name: 'Aggression', visible: ref(false)},
    {attr: 'ant', name: 'Anticipation', visible: ref(false)},
    {attr: 'bra', name: 'Bravery', visible: ref(false)},
    {attr: 'cmp', name: 'Composure', visible: ref(false)},
    {attr: 'cnt', name: 'Concentration', visible: ref(false)},
    {attr: 'dec', name: 'Decisions', visible: ref(false)},
    {attr: 'det', name: 'Determination', visible: ref(false)},
    {attr: 'fla', name: 'Flair', visible: ref(false)},
    {attr: 'ldr', name: 'Leadership', visible: ref(false)},
    {attr: 'otb', name: 'Off the Ball', visible: ref(false)},
    {attr: 'pos', name: 'Positioning', visible: ref(false)},
    {attr: 'tea', name: 'Teamwork', visible: ref(false)},
    {attr: 'vis', name: 'Vision', visible: ref(false)},
    {attr: 'wor', name: 'Work Rate', visible: ref(false)}
]
const phyAttr: {attr: string, name: string, visible: Ref<boolean>}[] = [
    {attr: 'acc', name: 'Acceleration', visible: ref(false)},
    {attr: 'agi', name: 'Agility', visible: ref(false)},
    {attr: 'bal', name: 'Balance', visible: ref(false)},
    {attr: 'jum', name: 'Jumping Reach', visible: ref(false)},
    {attr: 'nat', name: 'Natural Fitness', visible: ref(false)},
    {attr: 'pac', name: 'Pace', visible: ref(false)},
    {attr: 'sta', name: 'Stamina', visible: ref(false)},
    {attr: 'str', name: 'Strength', visible: ref(false)}
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

const setChartPlugins = () => {
    return {
        plugins: {
            legend: {
                labels: {
                    color: getComputedStyle(document.documentElement).getPropertyValue('--text-color')
                },
            }
        }
    }
}

const setChartOptions = () => {
    const documentStyle = getComputedStyle(document.documentElement)
    const textColor = documentStyle.getPropertyValue('--text-color')
    const textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary')
    const surfaceBorder = documentStyle.getPropertyValue('--surface-border')

    return {
        maintainAspectRatio: false,
        aspectRatio: 0.6,
        layout: {
            padding: 0
        },
        plugins: {
            legend: {
                position: 'bottom',
                display: true,
                labels: {
                    filter: function(item: any) {
                        return !item.hidden
                    }
                },
                onClick: function() {
                    return
                }
            }
        },
        scales: {
            x: {
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder
                }
            },
            y: {
                ticks: {
                    color:textColorSecondary
                },
                grid: {
                    color: surfaceBorder
                },
            }
        },
        parsing: {
            xAxisKey: 'season',
        },
        elements: {
            point: {
                radius: 6,
                hitRadius: 3,
                pointStyle: 'rectRounded',
                spanGaps: true
            }
        }
    }
}

const setChartData = () => {
    let listOfAttr: {attr: string, name: string, visible: Ref<boolean>}[][]
    if (playerAvgs.value.position != 'GK') {
        listOfAttr = [techAttr, mentalAttr, phyAttr]
    } else {
        listOfAttr = [gkAttr, mentalAttr, phyAttr]
    }
    let attributes: {label: string, data: {}[], parsing: { yAxisKey: string}, hidden: boolean}[] = []
    listOfAttr.forEach( (attrList) => {
        attrList.forEach( (attribute) => {
            attributes.push({
                label: attribute.name,
                data: playerSeasons.value,
                parsing: {
                    yAxisKey:attribute.attr
                },
                hidden: !attribute.visible.value
            })
        })
    })
    return {
        datasets: attributes
    }
}

const updateChartData = () => {
    chartData.value = setChartData()
}

onBeforeMount( () => {
    playerID.value = props.playerID
    GetSinglePlayer(props.playerID).then( (response) => {
        if (response.Error != "") {
            emit('beError', response.Error)
            return
        }
        playerSeasons.value = response.OnePlayer
        playerAvgs.value = response.PlayerAvgSum
        seasons.value = ([...new Set(playerSeasons.value.map((item: backend.PlayerPageInfo ) => item.season))])
        // console.log(playerSeasons.value)
        if (playerAvgs.value.position == 'GK') {
            isGK.value = true
        }
        dataLoaded.value = true
        chartData.value = setChartData()
        chartOptions.value = setChartOptions()
        chartPlugins.value = setChartPlugins()
    })
})

</script>

<style scoped lang="scss">
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
        background: none!important;
        border: 0px!important;
        box-shadow: none!important;
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

    .selectionCard :deep(.p-card-header) {
        border-bottom: 1px solid var(--surface-border);
        padding-bottom: 3px!important;
    }

    .p-splitter {
        border: none;
        background: none;
    }


</style>