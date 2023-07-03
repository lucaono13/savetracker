<template>
    <Dialog header="Add Save" @hide="" :modal="true" :draggable="false" :closable="false"
        class="w-5">
        <form method="POST" id='addForm' @submit.prevent="newSaveAdded(!v$.$invalid)">
            <div class="field">
                <div class="p-float-label mt-4 field">
                    <InputText class="w-full" id="saveName" type="text" :disabled="addingToDB"
                        v-model="v$.saveName.$model" autofocus
                        :class="{ 'p-invalid': v$.saveName.$invalid && submitted }" />
                    <label for="saveName">Save Name</label>
                    <small v-if="(v$.saveName.$invalid && submitted) || v$.saveName.$pending"
                        class="p-error">{{ v$.saveName.required.$message.replace('Value', 'Save Name') }}</small>
                </div>
            </div>
            <div class="p-float-label mt-4 field">
                <InputText class="w-full" id="managerName" type="text" :disabled="addingToDB"
                    v-model="v$.managerName.$model" :class="{ 'p-invalid': v$.managerName.$invalid && submitted }" />
                <label for="managerName">Manager Name</label>
                <small v-if="(v$.managerName.$invalid && submitted) || v$.managerName.$pending"
                    class="p-error">{{ v$.managerName.required.$message.replace('Value', 'Manager Name') }}</small>
            </div>
                <div class="field mt-4 p-float-label">
                    <Dropdown class="w-full" v-model="v$.gameVersion.$model" id="gVersion" :disabled="addingToDB"
                        :options="versions" optionLabel="name" optionValue="name" 
                        :class="{ 'p-invalid': v$.gameVersion.$invalid && submitted }" />
                    <label for="gVersion">Game Version</label>
                    <small v-if="(v$.gameVersion.$invalid && submitted) || v$.gameVersion.$pending"
                        class="p-error">{{ v$.gameVersion.required.$message.replace('Value', 'Game Version') }}</small>
                </div>
                <div class="field mt-4 p-float-label">
                    <Dropdown class="w-full" v-model="v$.currency.$model" id="currency" :disabled="addingToDB"
                    :options="currencies" :filterFields="['name','symbol','code']" filter placeholder="Select Save Currency"
                    :class="{ 'p-invalid': v$.currency.$invalid && submitted }">
                        <template #value="slotProps">
                            <div v-if="slotProps.value">{{ slotProps.value.name }} - {{ slotProps.value.symbol }}</div>
                        </template>
                        <template #option="slotProps">
                            <div >{{ slotProps.option.name }} - {{ slotProps.option.symbol }}</div>
                        </template>
                    </Dropdown>
                    <label for="currency">Currency</label>
                    <small v-if="(v$.currency.$invalid && submitted) || v$.currency.$pending"
                        class="p-error">{{ v$.gameVersion.required.$message.replace('Value', 'Currency') }}</small>
                </div>
            <!-- </div> -->
            
        </form>
        <template #footer>
            <div class="flex align-content-center justify-content-center">
                
                <Button label="Cancel" id="cancelB" :disabled="addingToDB" class='p-button-text' @click="$emit('closeDialog')" />
                <Button label="Add" id="submitB" :disabled="addingToDB" type="submit" form="addForm" />
            </div>
            <div v-if="addingToDB" class="mt-3 justify-content-start">
                <p class="mb-1">Adding...</p>
                <ProgressBar style="height: 0.3em" mode="indeterminate" />
            </div>
        </template>
    </Dialog>
</template>



<script setup lang="ts">
    import { ref, reactive } from 'vue'
    import { required } from '@vuelidate/validators';
    import { useVuelidate } from "@vuelidate/core";
    import { AddNewSave } from "../../../wailsjs/go/main/App";
    import currencies from '../../../../currency-names.json'

    const emit = defineEmits(['saveAdded', 'beError', 'closeDialog'])

    const state = reactive({
        saveName: '',
        managerName: '',
        gameVersion: '',
        currency: {"name": '', "symbol": '', "code": ''},
    })

    const rules = {
        saveName: { required },
        managerName: { required },
        gameVersion: { required },
        currency: { required },
    }

    const versions = ref([
        { name: '2023'},
    ])
    const blockedDocument = ref(false)
    const addingToDB = ref(false)
    const submitted = ref(false)

    const v$ = useVuelidate(rules, state)
    
    function AddSave() {
        let fields = v$.value
        AddNewSave(fields.saveName.$model, fields.managerName.$model, +fields.gameVersion.$model, fields.currency.$model.name).then( (response) => {
            if (response.Error != '') {
                emit('beError', response.Error)
                return
            }
            emit('saveAdded', response.Integer)
        })
    }

    function newSaveAdded(isFormValid: boolean) {
        submitted.value = true

        if (!isFormValid) {
            return
        }
        BlockDocument()
    }

    function BlockDocument() {
        addingToDB.value = true
        setTimeout( () => {
            AddSave()
            addingToDB.value = false
            resetForm()
            emit('closeDialog')
        }, 2000)
    }

    function resetForm() {
        v$.value.currency.$model = {"name": '', "symbol": '', "code": ''}
        v$.value.gameVersion.$model = ''
        v$.value.saveName.$model = ''
        v$.value.managerName.$model = ''
    }

</script>