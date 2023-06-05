<template>
    <Dialog header="Add Save" @hide="resetForm" :modal="true" :draggable="false" :closable="false"
        class="w-5">
        <form method="POST" id='addForm' @submit.prevent="newSaveAdded(!v$.$invalid)">
            <div class="field">
                <div class="p-float-label mt-4 field">
                    <InputText class="w-full" id="saveName" type="text" :disabled="addingToDB"
                        v-model="v$.saveName.$model" autofocus
                        :class="{ 'p-invalid': v$.saveName.$invalid && submitted }" />
                    <label for="saveName">Save Name</label>
                    <small v-if="(v$.saveName.$invalid && submitted) || v$.saveName.$pending.$response"
                        class="p-error">{{ v$.saveName.required.$message.replace('Value', 'Save Name') }}</small>
                </div>
            </div>
            <div class="p-float-label mt-4 field">
                <InputText class="w-full" id="managerName" type="text" :disabled="addingToDB"
                    v-model="v$.managerName.$model" :class="{ 'p-invalid': v$.managerName.$invalid && submitted }" />
                <label for="managerName">Manager Name</label>
                <small v-if="(v$.managerName.$invalid && submitted) || v$.managerName.$pending.$response"
                    class="p-error">{{ v$.managerName.required.$message.replace('Value', 'Manager Name') }}</small>
            </div>
            <!-- <div class="formgroup-inline grid"> -->
                <div class="field mt-4 p-float-label">
                    <Dropdown class="w-full" v-model="v$.gameVersion.$model" id="gVersion" :disabled="addingToDB"
                        :options="versions" optionLabel="name" optionValue="name" 
                        :class="{ 'p-invalid': v$.gameVersion.$invalid && submitted }" />
                    <label for="gVersion">Game Version</label>
                    <small v-if="(v$.gameVersion.$invalid && submitted) || v$.gameVersion.$pending.$response"
                        class="p-error">{{ v$.gameVersion.required.$message.replace('Value', 'Game Version') }}</small>
                </div>
                <!-- <div class="field mt-4 p-float-label">
                    <InputText class="w-full" v-model="v$.currency.$model" id="curr" :disabled="addingToDB"
                        :class="{ 'p-invalid': v$.currency.$invalid && submitted }" />
                    <label for="curr">Currency (i.e. $, €, £)</label>
                    <small v-if="(v$.currency.$invalid && submitted) || v$.currency.$pending.$response"
                    class="p-error">{{ v$.currency.required.$message.replace('Value', 'Currency') }}</small>
                </div> -->
                <div class="field mt-4 p-float-label">
                    <Dropdown class="w-full" v-model="v$.currency.$model" id="currency" :disabled="addingToDB"
                    :options="currencies" :filterFields="['name','symbol','code']" filter placeholder="Select Save Currency"
                    :class="{ 'p-invalid': v$.currency.$invalid && submitted }">
                        <template #value="slotProps">
                            <div v-if="slotProps.value">{{ slotProps.value.name }} - {{ slotProps.value.symbol }}</div>
                            <span v-else>{{ slotProps.placeholder }}</span>
                        </template>
                        <template #option="slotProps">
                            <div>{{ slotProps.option.name }} - {{ slotProps.option.symbol }}</div>
                        </template>
                    </Dropdown>
                    <label for="currency">Currency</label>
                    <small v-if="(v$.currency.$invalid && submitted) || v$.currency.$pending.$response"
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

<script lang="ts">
import { required } from '@vuelidate/validators';
import { useVuelidate } from "@vuelidate/core";
import { AddNewSave } from "../../../wailsjs/go/main/App";
import { string } from 'yup';
import currencies from '../../../../currency-names.json'

const curr = currencies



export default {
    setup: () => ({ v$: useVuelidate() }),
    data() {
        return {
            addSaveModal: false,
            saveName: '',
            gameVersion: '',
            managerName: '',
            currency: {"name":'',"symbol":'',"code":''},
            currencyToSubmit: '',
            blockedDocument: false,
            submitted: false,
            versions: [
                { name: "2023" },
            ],
            currFilters: [
                "name",
                "symbol",
                "code"
            ],
            currencies: currencies,
            addingToDB: false,
        }
    },
    methods: {
        openAddSave() {
            this.addSaveModal = true
        },
        addSave() {
            AddNewSave(this.saveName, this.managerName, +this.gameVersion, this.currency.code).then((response) => {
                // console.log(response)
                // return
                // response.Error = "just a quick test thanks\nwe're going multiple lines!"
                // if (response.Error != "") {
                //     console.log('we got here?')
                //     this.$emit('beError', response.Error)
                // }
                if (response.Integer == 0) {
                    this.$emit('saveAdded', response.Integer)
                } else {
                    this.$emit('saveAdded', response.Integer)
                }
                // this.resetForm()
                
            })

        },
        newSaveAdded(isFormValid: boolean) {
            this.submitted = true;
            // const saveNo = localStorage.getItem("saves")
            
            console.log(this.currency.code)

            if (!isFormValid) {
                return;
            }
            this.blockDocument();
        },
        blockDocument() {
            this.addingToDB = true
            this.addSave()

            setTimeout(() => {
                this.addSaveModal = false;
                this.addingToDB = false;
                this.resetForm()
                this.$emit('closeDialog')
            }, 3000);

        },
        resetForm() {
            this.saveName = ''
            this.gameVersion = ''
            this.managerName = ''
            this.currency = {"name":'',"symbol":'',"code":''}
            this.submitted = false;
        },
    },
    // eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
    validations() {
        return {
            saveName: {
                required
            },
            managerName: {
                required
            },
            gameVersion: {
                required
            },
            currency: {
                required
            }

        }
    }
}
</script>