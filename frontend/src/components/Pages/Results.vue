<template>
    <div class="card dt-size pr-4 pb-4">
        <DataTable stripedRows :rows="30" scrollable scrollHeight="flex" class="p-datatable-sm "
            paginator :rowsPerPageOptions="[5, 10, 20, 30, 50, 75]" :value="results"
            tableStyle="min-width: 1000px" filterDisplay="menu" v-model:filters="filters" :lazy="false" v-on:filter="GetRecord">
            <template #empty>
                        <InlineMessage severity="warn">
                            No data found! Add new season by clicking the "New Season" button to the left.
                        </InlineMessage>
                    </template>
            <template #header>W-D-L: {{ wins }}-{{ draws }}-{{ losses }} ({{ totalGames }} GP)  
                GF: {{ NumberFormatterR.format(GF) }} (AVG: {{ NumberFormatterR.format(GF / totalGames) }}) GA: {{ NumberFormatterR.format(GA) }} (AVG: {{ NumberFormatterR.format(GA / totalGames) }})</template>
            <Column field="date" header="Date" class="min-w-min">
            </Column>
            <Column filterField="year" header="Year" class="min-w-min" :showFilterMatchModes="false">
                <template #body="{ data }">{{ data.year }}</template>
                <template #filter="{ filterModel }">
                    <MultiSelect v-model="filterModel.value" :options="uniqueYears" placeholder="Select Year" class="p-column-filter" />
                </template>
            </Column>
            <Column field="competition" header="Competition" class="min-w-min">
                <template #filter="{ filterModel }">
                    <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by competition" />
                </template>
            </Column>
            <Column field="stadium" header="Stadium" class="min-w-min">
                <template #filter="{ filterModel }">
                    <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by stadium" />
                </template>
            </Column>
            <Column field="venue"  header="Venue" class="min-w-min" :showFilterMatchModes="false">
                <template #filter="{ filterModel }">
                    <MultiSelect v-model="filterModel.value" :options="venueType" placeholder="Select Venue" class="p-column-filter" />
                </template>
            </Column>
            <Column field="opposition" header="Opponent" class="min-w-min" >
                <template #filter="{ filterModel }">
                    <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                </template>
            </Column>
            <Column header="Score" filterField="result" class="min-w-min" :showFilterMatchModes="false">
                <template #body="slotProps">
                    <Badge class="p-badge-dot" :severity="getResult(slotProps.data.result)"></Badge>
                    <span>&nbsp;</span>
                    <span v-if="slotProps.data.venue == 'H' || slotProps.data.venue == 'N'">{{ slotProps.data.goalsFor }} - {{ slotProps.data.goalsAgainst }}</span>
                    <span v-if="slotProps.data.venue == 'A'">{{ slotProps.data.goalsAgainst }} - {{ slotProps.data.goalsFor }}</span>
                    <span v-if="slotProps.data.penalties == 1 || slotProps.data.extraTime == 1"> &nbsp;</span>
                    <span v-if="slotProps.data.penalties == 1">pens</span>
                    <span v-if="slotProps.data.extraTime == 1">a.e.t</span>
                </template>
                <template #filter="{ filterModel }" >
                    <MultiSelect v-model="filterModel.value" :options="resultType" placeholder="Select Result" class="p-column-filter" />
                </template>
            </Column>
        </DataTable>
    </div>
</template>

<script setup lang="ts">
import { GetSaveResults } from '../../../wailsjs/go/main/App'
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { backend } from '../../../wailsjs/go/models'
import { FilterMatchMode, FilterOperator } from 'primevue/api'

onMounted( () => {
    GetSaveResults(+route.params.id).then( (response) => {
        uniqueYears.value = ([...new Set(response.Matches.map((item: backend.Match ) => item.year))])
        results.value = response.Matches
    })
})

const route = useRoute()
const results = ref()
const wins = ref()
const losses = ref()
const draws = ref()
const totalGames = ref()
const resultType = ref(['W','D','L'])
const venueType = ref(['H','A','N'])
const filters = ref()
const uniqueYears = ref()
const GF = ref(0)
const GA = ref(0)
let NumberFormatterR: Intl.NumberFormat = new Intl.NumberFormat(navigator.language, {
    style: "decimal",
    maximumFractionDigits: 2
})
const initFilters = () => {
    filters.value = {
        opposition: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        year: { value: null, matchMode: FilterMatchMode.IN },
        result: { value: null, matchMode: FilterMatchMode.IN },
        venue : {value: null, matchMode: FilterMatchMode.IN},
        competition: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        stadium: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  }
    }
}

initFilters()
const getResult = (match: string) => {
    switch (match) {
        case 'W':
            return 'success'
        case 'D':
            return 'warning'
        case 'L':
            return 'danger'
        default:
            return 'info'
    }
}



function GetRecord(e: {originalEvent: Event, filteredValue: backend.Match[]}) {
    GF.value = 0
    GA.value = 0
    totalGames.value = e.filteredValue.length
    if (e.filteredValue.length > 0) {
        wins.value = e.filteredValue.filter( (obj: backend.Match) => {
            return obj.result == "W"
        }).length
        losses.value = e.filteredValue.filter( (obj: backend.Match) => {
            return obj.result == "L"
        }).length
        draws.value = e.filteredValue.filter( (obj: backend.Match) => {
            return obj.result == "D"
        }).length
    }
    e.filteredValue.forEach( function (value) {
        GF.value += value.goalsFor
        GA.value += value.goalsAgainst
    })
}


</script>

<style scoped lang="scss">
    .dt-size {
        height: calc(100vh - 79px)!important;
        width: calc(100vw - 205px)!important;
    }

    :deep .p-datatable-flex-scrollable .p-datatable-scrollable-wrapper {
        overflow: hidden;
    }
</style>
