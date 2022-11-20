<template>
<div class="flex align-items-center justify-content-center h-full">
    
    <BlockUI :blocked="blockedDocument" :fullScreen="true">
        
    </BlockUI>
    <Card style="width:500px">
        <template #title>
            No Saves DB Found
        </template>
        <template #content>
            Either this is the first time loading/using the app or you deleted all your saves! (Hopefully not by accident).
            Click on the "Add Save" button below to add a save!
        </template>
        <template #footer>
            <Button label="Add Save" class="p-button" @click="openAddSave" />
        </template>
    </Card>
    <Dialog header="Add Save" v-model:visible="addSaveModal" @hide="resetForm" :modal="true" :draggable="false" class="w-5">
            <!-- <h3 class="mb-5">Add Save</h3> -->
            <form method="POST" id='addForm' @submit.prevent="newSaveAdded(!v$.$invalid)">
                <div class="field">
                    <div class="p-float-label mt-4 field">
                        <InputText class="w-full" id="saveName" type="text" :disabled="addingToDB" v-model="v$.saveName.$model" autofocus :class="{'p-invalid':v$.saveName.$invalid && submitted}"/>
                        <label for="saveName">Save Name*</label>
                        <small v-if="(v$.saveName.$invalid && submitted) || v$.saveName.$pending.$response" class="p-error">{{v$.saveName.required.$message.replace('Value', 'Save Name')}}</small>
                    </div>
                </div>
                <div class="p-float-label mt-4 field">
                    <InputText class="w-full"  id="managerName" type="text" :disabled="addingToDB" v-model="v$.managerName.$model" :class="{'p-invalid':v$.managerName.$invalid && submitted}" />
                    <label for="managerName">Manager Name*</label>
                    <small v-if="(v$.managerName.$invalid && submitted) || v$.managerName.$pending.$response" class="p-error">{{v$.managerName.required.$message.replace('Value', 'Manager Name')}}</small>
                </div>
                <div class="field mt-4" >
                    <!-- <label for="version">Game Version</label> -->
                    <Dropdown class="w-full" v-model="v$.gameVersion.$model" id="gVersion" :disabled="addingToDB" :options="versions" optionLabel="name" optionValue="name" placeholder="Select Game Version" :class="{'p-invalid':v$.gameVersion.$invalid && submitted}"/>
                    <small v-if="(v$.gameVersion.$invalid && submitted) || v$.gameVersion.$pending.$response" class="p-error">{{v$.gameVersion.required.$message.replace('Value', 'Game Version')}}</small>
                </div>
            </form>
            <template #footer >
                <div class="flex align-content-center justify-content-center">
                    <Button label="Cancel" id="cancelB" :disabled="addingToDB" class='p-button-text' @click="closeAddSave" />
                    <Button label="Add" id="submitB" :disabled="addingToDB" type="submit" form="addForm"/>
                </div>
                <div v-if="addingToDB" class="mt-3 justify-content-start">
                    <p class="mb-1">Adding...</p>
                    <ProgressBar style="height: 0.3em" mode="indeterminate" />
                </div>
            </template>
        </Dialog>
</div>

</template>

<script lang="ts">
    import { required } from "@vuelidate/validators";
    import { useVuelidate } from "@vuelidate/core";
    export default {
        setup: () => ({ v$: useVuelidate()}),
        data() {
            return {
                addSaveModal: false,
                saveName: '',
                gameVersion: '',
                managerName: '',
                blockedDocument: false,
                submitted: false,
                versions: [
                    {name: "2023"},
                    {name: "2022"},
                    {name: "2021"},
                    {name: "2020"}
                ],
                addingToDB: false,
            }
        },
        methods: {
            openAddSave() {
                this.addSaveModal = true
            },
            closeAddSave() {
                this.resetForm()
                this.addSaveModal = false;
            },
            newSaveAdded(isFormValid: boolean) {
                this.submitted = true;

                if (!isFormValid) {
                    return;
                }
                console.log(this.saveName)
                // this.disableForm()
                this.blockDocument();
            },
            blockDocument() {
                // this.blockedDocument = true;
                this.addingToDB = true

                setTimeout(() => {
                    // this.blockedDocument = false;
                    this.addingToDB = false;
                    this.closeAddSave()
                    this.resetForm()
                }, 3000);
                
            },
            resetForm() {
                this.saveName =''
                this.gameVersion = ''
                this.managerName = ''
                this.submitted = false;
            },
            disableForm() {
                let inputs = [
                    'saveName',
                    'managerName',
                    'gVersion',
                    'cancelB',
                    'submitB'
                ]
                // for (let i = 0; i < inputs.length; i++) {
                //     let elem = document.getElementById(inputs[i]);
                //     elem?.ariaDisabled = true
                // }
            }
            // handleSubmit(isFormValid) {
            //     this.submitted = true;

            //     if (isFormValid) {
            //         return;
            //     }
            //     this.blockDocument();
            // }
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