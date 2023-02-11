<template>
    <Dialog header="Add Save" @hide="resetForm" :modal="true" :draggable="false" :closable="false"
        class="w-5">
        <form method="POST" id='addForm' @submit.prevent="newSaveAdded(!v$.$invalid)">
            <div class="field">
                <div class="p-float-label mt-4 field">
                    <InputText class="w-full" id="saveName" type="text" :disabled="addingToDB"
                        v-model="v$.saveName.$model" autofocus
                        :class="{ 'p-invalid': v$.saveName.$invalid && submitted }" />
                    <label for="saveName">Save Name*</label>
                    <small v-if="(v$.saveName.$invalid && submitted) || v$.saveName.$pending.$response"
                        class="p-error">{{ v$.saveName.required.$message.replace('Value', 'Save Name') }}</small>
                </div>
            </div>
            <div class="p-float-label mt-4 field">
                <InputText class="w-full" id="managerName" type="text" :disabled="addingToDB"
                    v-model="v$.managerName.$model" :class="{ 'p-invalid': v$.managerName.$invalid && submitted }" />
                <label for="managerName">Manager Name*</label>
                <small v-if="(v$.managerName.$invalid && submitted) || v$.managerName.$pending.$response"
                    class="p-error">{{ v$.managerName.required.$message.replace('Value', 'Manager Name') }}</small>
            </div>
            <div class="field mt-4">
                <Dropdown class="w-full" v-model="v$.gameVersion.$model" id="gVersion" :disabled="addingToDB"
                    :options="versions" optionLabel="name" optionValue="name" placeholder="Select Game Version"
                    :class="{ 'p-invalid': v$.gameVersion.$invalid && submitted }" />
                <small v-if="(v$.gameVersion.$invalid && submitted) || v$.gameVersion.$pending.$response"
                    class="p-error">{{ v$.gameVersion.required.$message.replace('Value', 'Game Version') }}</small>
            </div>
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
export default {
    setup: () => ({ v$: useVuelidate() }),
    data() {
        return {
            addSaveModal: false,
            saveName: '',
            gameVersion: '',
            managerName: '',
            blockedDocument: false,
            submitted: false,
            versions: [
                { name: "2023" },
                { name: "2022" },
                { name: "2021" },
                { name: "2020" }
            ],
            addingToDB: false,
        }
    },
    methods: {
        openAddSave() {
            this.addSaveModal = true
        },
        addSave() {
            AddNewSave(this.saveName, this.managerName, +this.gameVersion).then((response) => {
                if (response == 0) {
                    this.$emit('saveAdded')
                } else {
                    this.$emit('saveAdded', response)
                }
                // this.resetForm()
                
            })

        },
        newSaveAdded(isFormValid: boolean) {
            this.submitted = true;
            // const saveNo = localStorage.getItem("saves")
            
            

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
            }

        }
    }
}
</script>