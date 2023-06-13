<template>
    <div class="comp-size pr-4 pb-4">
        <TabView>
            <TabPanel header="Outfield Players">
                <DataTable stripedRows :rows="15" scrollable scrollHeight="flex" class="p-datatable-sm"
                    paginator :rowsPerPageOptions="[5,15,20,30,50]" :value="outfield" tableStyle="min-width: 1812px"
                    filterDisplay="menu" v-model:filters="filters" :lazy="false" removableSort sortMode="multiple" reorderableColumns>
                    <template #empty>
                        <InlineMessage severity="warn">
                            No data found! Add new season by clicking the "New Season" button to the left.
                        </InlineMessage>
                    </template>
                    <Column field="playerName" header="Player" sortable frozen class="min-w-min" style="min-width: 205px!important;" :reorderableColumn="false">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column field="position" header="Position" class="min-w-min" style="min-width: 150px!important;" :reorderableColumn="false">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column field="teamName" header="Team" class="min-w-min" :reorderableColumn="false">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column field="seasons" header="Seasons" class="min-w-min justify-content-center text-center" style="min-width: 90px!important;" :reorderableColumn="false" sortable>
                    </Column>
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
                            {{ numberFormatterSQ.format(slotProps.data.avgPassP) }}%
                        </template>
                    </Column>
                    <Column field="avgWinP" header="Win %" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ numberFormatterSQ.format(slotProps.data.avgWinP) }}%
                        </template>
                    </Column>
                </DataTable>
            </TabPanel>
            <TabPanel header="Goalies">
                <DataTable stripedRows :rows="15" scrollable scrollHeight="flex" class="p-datatable-sm "
                    paginator :rowsPerPageOptions="[5,15,20,30,50]" :value="gks" tableStyle="min-width: 1552px"
                    filterDisplay="menu" v-model:filters="GKfilters" :lazy="false" removableSort sortMode="multiple" reorderableColumns>
                    <template #empty>
                        <InlineMessage severity="warn">
                            No data found! Add new season by clicking the "New Season" button to the left.
                        </InlineMessage>
                    </template>
                    <Column field="playerName" header="Player" sortable class="min-w-min" style="min-width: 205px!important;" :reorderableColumn="false">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column field="position" header="Position" class="min-w-min" style="min-width: 150px!important;" :reorderableColumn="false">
                    </Column>
                    <Column field="teamName" header="Team" class="min-w-min" :reorderableColumn="false">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column field="seasons" header="Seasons" class="min-w-min justify-content-center text-center" :showFilterMatchModes="false" :reorderableColumn="false" sortable>
                    </Column>
                    <Column field="minutes" header="Minutes" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="starts" header="Appearances" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="shutouts" header="Shutouts" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="avgSaveP" header="Save %" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ numberFormatterSQ.format(slotProps.data.avgSaveP) }}%
                        </template>
                    </Column>
                    <Column field="avgRating" header="Avg. Rating" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ numberFormatterSQ.format(slotProps.data.avgRating) }}
                        </template>
                    </Column>
                    <Column field="playerOfTheMatch" header="Player of the Match" class="min-w-min justify-content-center text-center" sortable></Column>
                    <Column field="avgPassP" header="Pass %" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ numberFormatterSQ.format(slotProps.data.avgPassP) }}%
                        </template>
                    </Column>
                    <Column field="avgWinP" header="Win %" class="min-w-min justify-content-center text-center" sortable>
                        <template #body="slotProps">
                            {{ numberFormatterSQ.format(slotProps.data.avgWinP) }}%
                        </template>
                    </Column>
                    
                </DataTable>
            </TabPanel>
        </TabView>
    </div>
</template>

<script setup lang="ts">
import { GetSavePlayersTotals } from '../../../wailsjs/go/main/App';
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { backend } from '../../../wailsjs/go/models'
import { FilterMatchMode, FilterOperator } from 'primevue/api';

const route = useRoute()
const gks = ref()
const outfield = ref()
const filters = ref()
const GKfilters = ref()
const uniqueYears = ref()
const emit = defineEmits(['beError'])
let numberFormatterSQ: Intl.NumberFormat 

onMounted( () => {
    GetSavePlayersTotals(+route.params.id).then( (response) => {
        if (response.Error != "") {
            emit('beError', response.Error)
            return
        }
        outfield.value = response.OutTotals
        gks.value = response.GKTotals
        numberFormatterSQ = new Intl.NumberFormat(navigator.language, {
            style: "decimal",
            maximumFractionDigits: 2

        })
    })
})

const initFilters = () => {
    filters.value = {
        teamName: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        playerName: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        position: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  }
    }
    GKfilters.value = {
        teamName: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        playerName: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        position: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  }
    }
}

initFilters()
</script>

<style scoped lang="scss">
    .comp-size {
        height: calc(100vh - 79px)!important;
        width: calc(100vw - 205px)!important;
    }

    .p-tabview :deep(.p-tabview-nav li.p-highlight .p-tabview-nav-link) {
        // background: none!important;
        color: #5eead4
    }

    .p-tabview :deep(.p-tabview-nav li .p-tabview-nav-link) {
        background: none!important;
        color:white;
    }

    .p-tabview :deep(.p-tabview-panels) {
        background: none!important;
        height:100%!important
    }

    .p-tabview :deep(.p-tabview-nav) {
        border: none!important;
    }

    :deep(.p-tabview-panels) {
        padding: 0%!important;
        height: 100%!important
    }

    :deep(.p-tabview-panel) {
        height: calc(100vh - 79px - 56px - 16px)!important;
    }

</style>