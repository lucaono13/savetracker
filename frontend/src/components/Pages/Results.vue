<template>
    <div class="card dt-size pr-4 pb-4">
        <DataTable stripedRows :rows="10" scrollable scrollHeight="flex" class="p-datatable-sm"
            paginator :rowsPerPageOptions="[5, 10, 20, 30, 50]" :value="results"
            tableStyle="min-width: 906px" filterDisplay="menu" v-model:filters="filters" :lazy="false" v-on:filter="GetRecord"
            >
            <template #header>W-D-L: {{ wins }}-{{ draws }}-{{ losses }} ({{ totalGames }} GP) </template>
            <!-- <Column v-for="col of columns"  :key="col.field" :field="col.field" :header="col.header"></Column> -->
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
                <template #body="{ data }">{{ data.opposition }}</template>
                <!-- <template #filter="{ filterModel }">
                    <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search Opposition" />
                </template> -->
                <template #filter="{ filterModel }">
                    <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                </template>
                <!-- <template #filterapply="{ filterCallback, filterModel }">
                    <Button type="button" icon="pi pi-check" @click="filterCallback()" severity="success"></Button>
                </template> -->
            </Column>
            <Column header="Score" filterField="result" class="min-w-min" :showFilterMatchModes="false">
                <template #body="slotProps">
                    <Badge class="p-badge-dot" :severity="getResult(slotProps.data.result)"></Badge>
                    <p>&nbsp;</p>
                    <p v-show="false">{{ slotProps.data.result }}</p>
                    <p v-if="slotProps.data.venue == 'H' || slotProps.data.venue == 'N'">{{ slotProps.data.goalsFor }} - {{ slotProps.data.goalsAgainst }}</p>
                    <p v-if="slotProps.data.venue == 'A'">{{ slotProps.data.goalsAgainst }} - {{ slotProps.data.goalsFor }}</p>
                    <p v-if="slotProps.data.penalties == 1 || slotProps.data.extraTime == 1"> &nbsp;</p>
                    <p v-if="slotProps.data.penalties == 1">pens</p>
                    <p v-if="slotProps.data.extraTime == 1">a.e.t</p>
                </template>
                <template #filter="{ filterModel }" >
                    <MultiSelect v-model="filterModel.value" :options="resultType" placeholder="Select Result" class="p-column-filter" />
                </template>
            </Column>
            <!-- <template #footer>{{ results ? results.length : 0 }}</template> -->
        </DataTable>
    </div>
</template>
<!-- filterDisplay="row" v-model:filters="filters" :globalFilterFields="['result']" -->

<script setup lang="ts">
import { GetSaveResults } from '../../../wailsjs/go/main/App'
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { backend } from '../../../wailsjs/go/models'
import { FilterMatchMode, FilterOperator } from 'primevue/api'
import { ColumnFilterModelType } from 'primevue/column'

onMounted( () => {
    console.log('mounted')
    GetSaveResults(+route.params.id).then( (response) => {
        uniqueYears.value = ([...new Set(response.Matches.map((item: backend.Match ) => item.year))])
        // console.log(uniqueYears.value)
        results.value = response.Matches
        // GetRecord()
        // GetTotalRecords()
    })
    // console.log([...new Set(results.value.map((item: backend.Match ) => item.year))])
})

const route = useRoute()
const results = ref()
const wins = ref()
const losses = ref()
const draws = ref()
const totalGames = ref()
const columns = [
    { field: "date", header: "Date" },
    { field: "year", header: "Year" },
    { field: "competition", header: "Competition" },
    { field: "stadium", header: "Stadium" },
    { field: "venue", header: "Venue" },
    { field: "opposition", header: "Opponent" },
    { field: "result", header: "Result" },
    { field: "goalsFor", header: "GF" },
    { field: "goalsAgainst", header: "GA" },
    { field: "penalties", header: "Penalties" },
    { field: "extraTime", header: "extraTime" }
]
const resultType = ref(['W','D','L'])
const venueType = ref(['H','A','N'])
const filters = ref()
const uniqueYears = ref()
// GetSaveResults(+route.params.id).then( (response) => {
//         const unique = ref([...new Set(response.Matches.map((item: backend.Match ) => item.year))])
//         console.log('sure')
//         results.value = response.Matches
//         GetRecord()
//     })
// const initFilters = () => {
//     filters.value = {
//         opposition: { value: null, matchMode: FilterMatchMode.}
//     }
// }
const initFilters = () => {
    filters.value = {
        opposition: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        year: { value: null, matchMode: FilterMatchMode.IN },
        result: { value: null, matchMode: FilterMatchMode.IN },
        venue : {value: null, matchMode: FilterMatchMode.IN},
        // competition: {value: null, matchMode: FilterMatchMode.CONTAINS}
        competition: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        stadium: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  }
    }
}

function testing(model: ColumnFilterModelType) {
    console.log(filters.value.opposition)
    console.log(model)
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
    // console.log(e)
    // console.log(e.filteredValue)
    totalGames.value = e.filteredValue.length
    if (e.filteredValue.length > 0) {
        console.log('1')
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
    
    // if (e != null) {
    //     console.log('2')
    //     if (e.filteredValue.length > 0) {
    //         console.log(e.filteredValue.length)
    //         // wins.value = e.filteredValue.filter( (obj: backend.Match) => {
    //         //     return obj.result == "W"
    //         // }).length
    //         // losses.value = e.filteredValue.filter( (obj: backend.Match) => {
    //         //     return obj.result == "L"
    //         // }).length
    //         // draws.value = e.filteredValue.filter( (obj: backend.Match) => {
    //         //     return obj.result == "D"
    //         // }).length
    //     }
    // }
}

function onFilter(event: { originalEvent: Event; filteredValue: []}) {
    console.log(event.filteredValue)
}
    // wins.value = results.value.filter( (obj: backend.Match) => {
    //     return obj.result == "W"
    // }).length
//     losses.value = results.value.filter( (obj: backend.Match) => {
//         return obj.result == "L"
//     }).length
//     draws.value = results.value.filter( (obj: backend.Match) => {
//         return obj.result == "D"
//     }).length
// }

// TODO: attempt to filter records properly to get right amount to update record for each filter
// function GetTotalRecords( model: ColumnFilterModelType) {
//     // https://bobbyhadz.com/blog/typescript-filter-array-of-objects
//     console.log(model["constraints"][0].value)
//     console.log(model.value)
//     // const values: backend.Match[] = allMatches.filter( (obj) => {

//     // })
    
// }

// const allMatches: backend.Match[] = results.value
// // const uniqueYears: any[] = [...new Set(allMatches.map((item: backend.Match ) => item.year))]
// console.log(allMatches.map(item => item.year))

</script>

<style lang="scss">
    // .p-datatable .p-datatable-tbody>tr{
    //     // background: rgba(0,0,0,0)!important;
    // }

    .dt-size {
        height: calc(100vh - 79px)!important;
        width: calc(100vw - 205px)!important;
    //     // width: 100%;
    //     overflow: scroll;
    }

    :deep .p-datatable-flex-scrollable .p-datatable-scrollable-wrapper {
        overflow: hidden;
    }
</style>