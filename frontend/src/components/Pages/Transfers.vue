<template>
    <div class="comp-size pr-3 pb-4">
        <TabView class="">
            <TabPanel header="Transfers In" class="">
                <!-- <p>In Transfers</p> -->
                <DataTable stripedRows :rows="15" scrollable scrollHeight="flex" class="p-datatable-sm "
                    paginator :rowsPerPageOptions="[5,15,20,30,50]" :value="inTransfers"
                    filterDisplay="menu" :rowClass="({ loan }) => loan === 1 ? 'font-italic' : null">
                    <Column field="date" header="Date"></Column>
                    <Column field="year" header="Year"></Column>
                    <Column field="team" header="From"></Column>
                    <Column field="player" header="Player"></Column>
                    <Column header="Fee">
                        <template #body="slotProps">
                            {{ formatFee(slotProps.data.fee, slotProps.data.potentialFee, slotProps.data.free, slotProps.data.loan) }}
                        </template>
                    </Column>
                </DataTable>
            </TabPanel>
            <TabPanel header="Transfers Out">
                <!-- <p>Out Transfers</p> -->
                <DataTable stripedRows :rows="10" scrollable scrollHeight="flex" class="p-datatable-sm "
                    paginator :rowsPerPageOptions="[5,10,20,30,50]" :value="outTransfers"
                    filterDisplay="menu" :rowClass="({ loan }) => loan === 1 ? 'font-italic' : null">
                    <Column field="date" header="Date"></Column>
                    <Column field="year" header="Year"></Column>
                    <Column field="team" header="To"></Column>
                    <Column field="player" header="Player"></Column>
                    <Column header="Fee">
                        <template #body="slotProps">
                            {{ formatFee(slotProps.data.fee, slotProps.data.potentialFee, slotProps.data.free, slotProps.data.loan) }}
                            <!-- <span v-if="slotProps.data.loan == 0"> Loan</span> -->
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
import { ColumnFilterModelType } from 'primevue/column'

const emit = defineEmits(['beError'])
// const currCode: string | null 
let numberFormatterComp: Intl.NumberFormat 
// = new Intl.NumberFormat(navigator.language, {
//     style: 'currency',
//     // currency: 
//     notation: 'compact'
// })

onMounted( () => {
    GetSaveTransfers(+route.params.id).then( (response) => {
        if (response.Error != "") {
            emit('beError', response.Error)
            
        }
        console.log(response)
        inTransfers.value = response.InTransfers
        outTransfers.value = response.OutTransfers
        currency.value = response.Currency
        numberFormatterComp = new Intl.NumberFormat(navigator.language, {
            style: 'currency',
            currency: currency.value,
            notation: "compact"
        })
    })
})

const formatFee = (fee: number, potFee: {Int64: number, Valid: boolean}, free: number, loan: number) => {
    if (free) {
        if (loan) {
            return "Loan"
        }
        return "Free"
    } 
    let allFees = numberFormatterComp.format(fee)
    if (potFee.Valid) {
        allFees = allFees + ' (' + numberFormatterComp.format(potFee.Int64) + ')'
    }
    if (loan) {
        allFees = allFees + ' (Loan)'
    }
    return allFees

}



const route = useRoute()
const inTransfers = ref()
const outTransfers = ref()
const currency = ref()
const filters = ref()


</script>

<style>
    .comp-size {
        height: calc(100vh - 79px)!important;
        width: calc(100vw - 205px)!important;
    }

    .p-tabview .p-tabview-nav li.p-highlight .p-tabview-nav-link {
        background: none!important;
    }

    .p-tabview .p-tabview-nav li .p-tabview-nav-link {
        background: none!important;
        color:white;
    }

    .p-tabview .p-tabview-panels {
        background: none!important;
        height:100%!important
    }

    .p-tabview .p-tabview-nav {
        border: none!important;
    }

    /* .p-tabview .p-component {
        height: 100%!important
    } */
    .p-tabview-panels {
        padding: 0%!important;
        height: 100%!important
    }

    .dt-size {
        height: calc(100vh - 79px - 56px - 16px)!important;
        width: calc(100vw - 205px)!important;
    }

    .p-tabview-panel {
        height: calc(100vh - 79px - 56px - 16px)!important;
    }

    /* :deep .p-datatable-flex-scrollable .p-datatable-scrollable-wrapper {
        overflow: hidden;
    } */
</style>