
<!-- import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'; -->

<template>
    <Dialog header="Add Season" @hide="" :modal="true" :draggable="false" :closable="true" class="w-7">
        <form method="POST" id="addSave" @submit.prevent="addSeason(!v$.$invalid)" class="">
            <div class="formgroup-inline pb-0">
                <span class="field w-6">
                    <label for="teamName" id="teamLabel">Team Name</label>
                    <InputText v-model="v$.teamName.$model" class="w-full" id="teamName" type="text" 
                        placeholder="Borussia Mönchengladbach"
                        :class="{ 'p-invalid': v$.teamName.$invalid && submitted }"
                    />
                    
                </span>
                <div class=" field w-5 gap-2">
                    <label for="shortName">Short Name</label>
                    <small class="font-italic" id="shortOptional">Optional</small>
                    <span class="p-inputgroup">
                        <InputText v-model="v$.shortName.$model" class="full" id="shortName" type="text" 
                            placeholder="Borussia M'gladbach" aria-describedby="shortOptional"
                        />
                        <Button severity="info" @click="explainSN" ><font-awesome-icon icon="fa-solid fa-question"/></Button>
                        <OverlayPanel ref="explain">
                            TODO: Add explanation
                    </OverlayPanel>
                    </span>
                    
                </div>
            </div>
            <div class="formgroup-inline ">
                <div class=" field w-3">
                    <label for="season">Season</label>
                    <InputMask v-model="v$.season.$model" class="w-full" id="season" type="text"  
                        mask="9999-99" placeholder="2023-24"
                        :class="{ 'p-invalid': v$.season.$invalid && submitted }"
                    />
                </div>
                <div class="field  w-8">
                    <label for="country">Country</label>
                    <InputText v-model="v$.country.$model" class="w-full" id="country" type="text" 
                        placeholder="Germany"
                        :class="{ 'p-invalid': v$.country.$invalid && submitted }"
                    />
                    
                </div>
            </div>
            <div class="formgroup-inline">
                <div class="field  w-12">
                    <label for="trophiesWon" class="">Trophies Won</label>
                    <small id="separateBy" class="font-italic">Separate by commas. Enter N/A if none</small>
                    <Textarea v-model="v$.trophies.$model" class="w-full fileText" id="trophiesWon" 
                        placeholder="Bundesliga, DFB-Pokal, Champions League" aria-describedby="separateBy"
                        :class="{ 'p-invalid': v$.trophies.$invalid && submitted }"
                    />
                </div>
            </div>
            <!-- <Divider/> -->
            <p style="font-size: small ;"><a type="button" @click="openLink($event, 'https://drive.google.com/drive/folders/1v7EZSIHoykRcBMWbPxGfl5Ux7Q2PosXv?usp=sharing')" href="">Use the views for Squad and Transfer in from this link to export the data properly.</a></p>
            <div class="mt-4 p-inputgroup fileGroup" >
                <Button style="width: 141px" class="flex-none justify-content-center" @click="GetFile('squadFile', 'Squad')" >Squad File</Button>
                <Textarea v-model="v$.squadFile.$model" class="fileText" id="squadFile"
                    cols="500" disabled
                    :class="{ 'p-invalid': v$.squadFile.$invalid && submitted }"
                />
            </div>
            <div class="p-inputgroup mt-4 fileGroup" >
                <Button style="width: 141px" class="justify-content-center" @click="GetFile('scheduleFile', 'Schedule')" >Schedule File</Button>
                <Textarea v-model="v$.scheduleFile.$model" class="fileText" id="scheduleFile" 
                    cols="500" disabled
                    :class="{ 'p-invalid': v$.scheduleFile.$invalid && submitted }"
                />
            </div>
            <div class="p-inputgroup mt-4 fileGroup" >
                <Button style="width: 141px" class="justify-content-center" @click="GetFile('transfersFile', 'Transfers')" >Transfers File</Button>
                <Textarea v-model="v$.transfersFile.$model" class="fileText" id="transfersFile"
                    cols="500" disabled
                    :class="{ 'p-invalid': v$.transfersFile.$invalid && submitted }"
                />
            </div>
            
        </form>
        <template #footer>
            <!-- <pre>{{ errors.country }}</pre> -->
            <div class="flex align-content-center justify-content-center">
                <Button label="Cancel" id="cancelB" class="p-button-text" @click="$emit('closeDialog')" />
                <Button label="Add" id="submitSeasonB"  form="addSave" type="submit"/>
            </div>
        </template>
    </Dialog>
</template>

<script setup lang="ts">
import { nextTick, ref, reactive } from 'vue'
import { BrowserOpenURL } from '../../../wailsjs/runtime/runtime';
import { useRoute } from 'vue-router'
import { SelectSquadFile, SelectScheduleFile, SelectTransfersFile, SelectFileParse, AddNewSeason } from '../../../wailsjs/go/main/App'
import { useForm, useField } from 'vee-validate';
import { useVuelidate } from "@vuelidate/core"
import { required } from '@vuelidate/validators'
import { main } from '../../../wailsjs/go/models'
import * as yup from 'yup';
import InputText from 'primevue/inputtext';

const route = useRoute()

const value = ref()
const explain = ref()
const submitted = ref(false)
const squadFileInput = ref<HTMLInputElement | null>()
const scheduleFileInput = ref<InputText | null>()
const transfersFileInput = ref<InputText | null>()

const explainSN = (event: any) => {
    explain.value.toggle(event)
}
const openLink = (event: any, url: string) => {
    event.preventDefault();
    BrowserOpenURL(url)
}

function GetFile(textArea: string, fileType: string) {
    SelectFileParse(fileType).then( (response) => {
        if (response.length > 0) {
            switch (textArea) {
                case 'squadFile': {
                    v$.value.squadFile.$model = response
                    break
                }
                case 'scheduleFile': {
                    v$.value.scheduleFile.$model = response
                    break
                }
                case 'transfersFile': {
                    v$.value.transfersFile.$model = response
                    break
                }
            }
        }
    })
}


const state = reactive({
    teamName: '',
    shortName: '',
    season: '',
    country: '',
    trophies: '',
    squadFile: '',
    scheduleFile: '',
    transfersFile: '',
})

const rules = {
    teamName: { required },
    shortName: {},
    season: { required },
    country: { required },
    trophies: { required },
    squadFile: { required },
    scheduleFile: { required },
    transfersFile: { required }
}


const v$ = useVuelidate(rules, state)

function addSeason(isValid: boolean) {
    console.log(v$.value.squadFile.$model, v$.value.scheduleFile.$model, v$.value.transfersFile.$model)
    let season: main.NewSeason = {
        teamName: v$.value.teamName.$model,
        shortName: v$.value.shortName.$model,
        season: v$.value.season.$model,
        country: v$.value.country.$model,
        trophiesWon: v$.value.trophies.$model,
        squadFile: v$.value.squadFile.$model,
        scheduleFile: v$.value.scheduleFile.$model,
        transfersFile: v$.value.transfersFile.$model,
    }
    submitted.value = true
    // console.log(v$.value.teamName.$model)
    if (!isValid) {
        return;
    }
    // ADD NEW SEASON BACKEND FUNCTION
}

</script>

<style lang="scss">

.p-dialog .p-dialog-header {
    padding-bottom: 0%!important;
}
.files {
    text-overflow: ellipsis;
    overflow: hidden;
    width: 500px
}

.fileGroup {
    height: 47px;
}

.fileText {
    max-height: 47px;
    min-height: 47px;
}

</style>