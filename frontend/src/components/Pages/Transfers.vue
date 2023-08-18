<template>
    <div class="comp-size pr-4 pb-4">
        <TabView>
            <TabPanel header="Transfers In" class="">
                <DataTable stripedRows :rows="15" scrollable scrollHeight="flex" class="p-datatable-sm "
                    paginator :rowsPerPageOptions="[5,15,20,30,50]" :value="inTransfers" tableStyle="min-width: 906px"
                    filterDisplay="menu" v-model:filters="filters" :rowClass="(loan:number) => loan === 1 ? 'font-italic' : null!"
                    v-on:filter="GetTotalSpent"  :lazy="false" removableSort>
                    <template #header>Total Spent: {{ totalSpent }} ({{ totalPotSpent }})</template>
                    <template #empty>
                        <InlineMessage severity="warn">
                            No data found! Add new season by clicking the "New Season" button to the left.
                        </InlineMessage>
                    </template>
                    <Column field="date" header="Date" class="min-w-min"></Column>
                    <Column field="year" header="Year" class="min-w-min" :showFilterMatchModes="false">
                        <template #filter="{ filterModel }">
                            <MultiSelect v-model="filterModel.value" :options="uniqueYears" placeholder="Select Year" class="p-column-filter" />
                        </template>
                    </Column>
                    <Column field="team" header="From" class="min-w-min">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column field="player" header="Player" class="min-w-min">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column header="Fee" class="min-w-min" sortable sortField="fee">
                        <template #body="slotProps">
                            {{ formatFee(slotProps.data.fee, slotProps.data.potentialFee, slotProps.data.free, slotProps.data.loan) }}
                        </template>
                    </Column>
                    
                </DataTable>
            </TabPanel>
            <TabPanel header="Transfers Out">
                <DataTable stripedRows :rows="15" scrollable scrollHeight="flex" class="p-datatable-sm "
                    paginator :rowsPerPageOptions="[5,15,20,30,50]" :value="outTransfers" tableStyle="min-width: 906px"
                    filterDisplay="menu" :rowClass="(loan: number) => loan === 1 ? 'font-italic' : null!" v-model:filters="filters"
                    v-on:filter="GetTotalReceived"  :lazy="false" removableSort>
                    <template #header>Total Received: {{ totalReceived }} ({{ totalPotReceived }})</template>
                    <template #empty>
                        <InlineMessage severity="warn">
                            No data found! Add new season by clicking the "New Season" button to the left.
                        </InlineMessage>
                    </template>
                    <Column field="date" header="Date" class="min-w-min"></Column>
                    <Column field="year" header="Year" class="min-w-min" :showFilterMatchModes="false">
                        <template #filter="{ filterModel }">
                            <MultiSelect v-model="filterModel.value" :options="uniqueYears" placeholder="Select Year" class="p-column-filter" />
                        </template>
                    </Column>
                    <Column field="team" header="To" class="min-w-min">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column field="player" header="Player" class="min-w-min">
                        <template #filter="{ filterModel }">
                            <InputText v-model="filterModel.value" type="text" class="p-column-filter" placeholder="Search by name" />
                        </template>
                    </Column>
                    <Column header="Fee" class="min-w-min" sortable sortField="fee">
                        <template #body="slotProps">
                            <p v-show="false">{{ slotProps.data.fee }}</p>
                            {{ formatFee(slotProps.data.fee, slotProps.data.potentialFee, slotProps.data.free, slotProps.data.loan) }}
                        </template>
                    </Column>
                    
                </DataTable>
            </TabPanel>
        </TabView>
    </div>
    
</template>


<script setup lang="ts">
import { GetSaveTransfers } from '../../../wailsjs/go/main/App'
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { backend } from '../../../wailsjs/go/models'
import { FilterMatchMode, FilterOperator } from 'primevue/api'

const route = useRoute()
const inTransfers = ref()
const outTransfers = ref()
const currency = ref()
const filters = ref()
const totalSpent = ref()
const totalPotSpent = ref()
const totalReceived = ref()
const totalPotReceived = ref()
const uniqueYears = ref()
const emit = defineEmits(['beError'])
let numberFormatterTR: Intl.NumberFormat 

onMounted( () => {
    GetSaveTransfers(+route.params.id).then( (response) => {
        if (response.Error != "") {
            emit('beError', response.Error)
            return
        }
        inTransfers.value = response.InTransfers
        outTransfers.value = response.OutTransfers
        uniqueYears.value = new Set(response.InTransfers.map((item: backend.Transfer) => item.year ))
        let outYears = new Set(response.OutTransfers.map((item: backend.Transfer) => item.year ))
        outYears.forEach(uniqueYears.value.add, uniqueYears.value)
        uniqueYears.value = ([...new Set(uniqueYears.value)])
        currency.value = response.Currency
        numberFormatterTR = new Intl.NumberFormat(navigator.language, {
            style: 'currency',
            currency: currency.value,
            notation: "compact"
        })
    })
})


const formatFee = (fee: number, potFee: {Int64: number, Valid: boolean}, free: number, loan: number) => {
    if (fee == 0 && !free && !loan) {
        return "Undisclosed"
    }
    if (free) {
        if (loan) {
            return "Loan"
        }
        return "Free"
    } 
    let allFees = numberFormatterTR.format(fee)
    if (potFee.Valid) {
        allFees = allFees + ' (' + numberFormatterTR.format(potFee.Int64) + ')'
    }
    if (loan) {
        allFees = allFees + ' (Loan)'
    }
    return allFees
}

const initFilters = () => {
    filters.value = {
        year: { value: null, matchMode: FilterMatchMode.IN },
        team: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
        player: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.CONTAINS }],  },
    }
}

initFilters()

const formatTotal = (total: number) => {
    let formattedTot: string = numberFormatterTR.format(total)
    return formattedTot
}

function GetTotalSpent( e: {originalEvent: Event, filteredValue: backend.Transfer[]}) {
    let spent = 0
    let potSpent = 0
    e.filteredValue.forEach(function (transfer) {
        spent += transfer.fee
        if (transfer.potentialFee.Valid == true) {
            potSpent += transfer.potentialFee.Int64
        } else {
            potSpent += transfer.fee
        }
    })
    totalSpent.value = numberFormatterTR.format(spent)
    totalPotSpent.value = numberFormatterTR.format(potSpent)
    
}

function GetTotalReceived( e: {originalEvent: Event, filteredValue: backend.Transfer[]}) {
    let receieved = 0
    let potReceived = 0
    e.filteredValue.forEach(function (transfer) {
        receieved += transfer.fee
        
        if (transfer.potentialFee.Valid == true) {
            potReceived += transfer.potentialFee.Int64
        } else {
            potReceived += transfer.fee
        }
    })
    totalReceived.value = numberFormatterTR.format(receieved)
    totalPotReceived.value = numberFormatterTR.format(potReceived)
}



</script>

<style scoped lang="scss">
    .comp-size {
        height: calc(100vh - 79px)!important;
        width: calc(100vw - 205px)!important;
    }

    .p-tabview :deep(.p-tabview-nav li.p-highlight .p-tabview-nav-link) {
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

    .dt-size {
        height: calc(100vh - 79px - 56px - 16px)!important;
        width: calc(100vw - 205px)!important;
    }

    :deep(.p-tabview-panel) {
        height: calc(100vh - 79px - 56px - 16px)!important;
    }

</style>