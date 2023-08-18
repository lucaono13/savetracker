<template>
  <Toast />
  <Menubar class="w-full  top-0 relative left-0" style="width: 98vw!important; padding-right: none!important;">
    <template #start style="width:250px!important" class="flex flex-row">
      <div class="flex align-items-center gap-2">
      <Button @click="GoHome"><FontAwesomeIcon :icon="['fas','house-chimney']"/></Button>
      <CascadeSelect style="width:80%" id="change" @change="SaveSelected" v-model="selectedSave" :options="isDataLoaded"
        optionLabel="name" optionGroupLabel="gameVersion" :option-group-children="['saves']" placeholder="Select save">
        <template #option="slotProps">
          <div class="flex align-items-center justify-content-center">
            <FontAwesomeIcon icon="fa-solid fa-list" v-if="slotProps.option.gameVersion" />
            <span style="padding-left: 5px;" v-if="slotProps.option.gameVersion">{{
              slotProps.option.gameVersion
            }}</span>
            <span style="padding-left: 5px;" v-if="slotProps.option.name">{{ slotProps.option.name }}</span>
          </div>
        </template>
      </CascadeSelect>
    </div>
    </template>
    <template #end>
      <Button class="p-button-help p-button-outlined" @click="addSaveModal = true">New Save</Button>
      <AddSaveDialog v-model:visible="addSaveModal" @beError="beError" @saveAdded="GetSaves" @closeDialog="addSaveModal = false" />
    </template>
  </Menubar>

  <div class="grid mt-0" style="width:100vw!important">
    <router-view @saveAdded="GetSaves" v-slot="{ Component, route }">
      <component :is="Component" :key="route.params.id" @getSaves="GetSaves" @beError="beError"></component>
    </router-view>
  </div>
  <Toast position="bottom-left" group="update" sticky>
    <template #message="slotProps" >
      <div class="flex flex-column align-items-center flex-1">
        <span class="p-toast-summary">{{ slotProps.message.summary }}</span>
        <a class='p-toast-detail' href="" @click="openLink($event, slotProps.message.detail)" style="color:var(--orange-500)">Click here to download.</a>
      </div>
      
    </template>
  </Toast>
  <Dialog :visible="aboutDialog" modal header="About Save Tracker" :draggable="false" :style="{ width: '50vw' }" @update:visible="aboutDialog = false">
    <!-- <p>Save Tracker</p> -->
    <p>Author: <a href="" @click="openLink($event, 'https://github.com/lucaono13')">Gianluca</a></p>
    <p><span class="font-bold">Version: </span><span class="font-italic">{{ appVersion }}</span></p>
    <p><span class="font-bold">DB Version: </span><span class="font-italic">{{ dbVersion }}</span></p>
    
    <p><span>Technologies Used:</span>
      <ul>
        <li><a href="" @click="openLink($event, 'https://wails.io')">Wails Framework</a></li>
        <li><a href="" @click="openLink($event, 'https://primevue.org')">PrimeVue Components</a></li>
        <li><a href="" @click="openLink($event, 'https://fontawesome.com')">Font Awesome Icons</a></li>
      </ul>
    </p>
  </Dialog>
</template>

<script lang="ts" async setup>
import { nextTick, ref, computed, reactive } from 'vue';
import { RetrieveSaves, GetImage } from '../wailsjs/go/main/App';
import Sidebar from './components/Components/Sidebar.vue'
import AddSaveDialog from "./components/Components/AddSaveDialog.vue";
import AddSeasonDialog from './components/Components/AddSeasonDialog.vue';
import { useRoute, useRouter } from 'vue-router';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { useToast } from 'primevue/usetoast';
import { BrowserOpenURL, EventsOn } from '../wailsjs/runtime/runtime';

const toast = useToast()
const route = useRoute()
const router = useRouter()
const startup = ref(true)
const addSaveModal = ref(false)
const addSeasonModal = ref(false)
const aboutDialog = ref(false)
const appVersion = ref()
const dbVersion = ref()

const newVersionAvail = (url: string) => {
  toast.add({
    severity: 'warn',
    summary: 'New Version Available',
    detail: url,
    group: 'update'
  })
}

EventsOn('updateVersion', (url: string) => newVersionAvail(url))
EventsOn('githubIssues', () => BrowserOpenURL( "https://github.com/lucaono13/savetracker/issues"))
EventsOn('saveDeleted', () => {
  selectedSave.value = {}
  GetSaves()
})
EventsOn('aboutDialog', (version: string, databaseVersion: string) => {
  appVersion.value = version
  dbVersion.value = databaseVersion
  aboutDialog.value = true
})


const openLink = (event: any, url: string) => {
    event.preventDefault();
    BrowserOpenURL(url)
}

// EventsOn('manageSaves', function() {
//   console.log('managing now!')
// })

// Getting number of saves, persists between opening of app
var noOfSaves: null | string = localStorage.getItem("saves")
if (noOfSaves == null) {
  noOfSaves = "0"
}
// Sidebar Visibility (1+ saves)
let sbVisible = ref(false)
if (+noOfSaves > 0) {
  sbVisible.value = true
}
let blank: string = 'hi'
// Default save
let defaultID: string | null = localStorage.getItem("defaultSave")
if (defaultID == null) {
  defaultID = "0"
}
let selectedSave = ref({})
let savesList: {}[] = []
let finalSaves = ref(savesList)
let isDataLoaded = computed(
  () => {
    return finalSaves.value.length > 0 ? finalSaves.value : []
  }
)

const GoHome = () => {
  selectedSave.value = {}
  router.replace({path: "/", replace: true})
}

const testing = reactive(finalSaves.value)
let isLoaded = ref(false)

function beError(error: string): void {
  toast.add({
    severity: 'error',
    summary: 'Error!',
    detail: error,
    life: 3000,
  })
}

// Function to be called to retrieve the saves and add to the dropdown
// Occurs: first app open, new save added
function GetSaves(newID?: number): void {
  RetrieveSaves().then(async (response) => {
    if (response.Error != "") {
      beError(response.Error)
      return
    }
    let saveList = response.SaveList
    let savesMap = new Map<string, {}[]>()
    for (var save in saveList) {
      let gV: string = saveList[save].gameVersion
      let saveObj: { id: number, name: string, manager: string, image: string | undefined } = { id: saveList[save].id, name: saveList[save].saveName, manager: saveList[save].managerName, image: saveList[save].saveImage }
      
      if (saveObj.image != undefined) {
        GetImage(saveObj.image).then( (response) => {
          saveObj.image = response.b64Image
        })
      }
      
      if (!savesMap.has(gV)) {
        savesMap.set(gV, [saveObj])
      } else if (savesMap.has(gV)) {
        let versionSaveList: {}[] = savesMap.get(gV)!
        versionSaveList.push(saveObj)
        savesMap.set(gV, versionSaveList)
      }

      
      


      if (startup.value == true && defaultID != null && saveObj.id == +defaultID) {
        selectedSave.value = saveObj
      } else if (startup.value == false && saveObj.id == newID) {
        selectedSave.value = saveObj
      }
    }
    startup.value = false
    finalSaves.value = Array.from(new Map([...savesMap.entries()].sort()), ([gameVersion, saves]) => ({ gameVersion, saves }))
    localStorage.setItem("saves", finalSaves.value.length.toString())
    isLoaded.value = true
    // nextTick()
    if (newID != null) {
      if (newID != 0) {
        GoToSave(newID)
      } else {
        if (finalSaves.value.length > 0 && defaultID != null){
          GoToSave(+defaultID)
          return
        } else {
          router.replace("/")
          return
        }
      }
    } else if (defaultID != null && defaultID != "" && defaultID != "0") {
      GoToSave(+defaultID)
      return
    }
  })
}

// Change path
function GoToSave(id: number) {
  let newRoute: string = '/save/' + id + '/home'
  router.replace(newRoute)
}

// Function to be called when the dropdown changes selection
function SaveSelected(e: { originalEvent: Event; value: { id: number; name: string } }) {
  GoToSave(e.value.id)
}
GetSaves()
</script>

<style lang="scss">
@font-face {
  font-family: 'Nunito';
  font-style: normal;
  font-weight: 400;
  src: local(''), url('assets/fonts/nunito-v16-latin-regular.woff2') format('woff2');
}

html {
  background-color: rgba(33, 37, 43, 1);
  color: white;
  font-family: 'Nunito', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell',
    'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
}

body {
  margin: 0;
}

h3 {
  margin: 0;
}

#app {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.link {
  cursor: pointer;
  color: #df0000;

  &:hover {
    text-decoration: underline;
  }
}

.p-menu {
  height: 100% !important;
  padding: .75rem !important;
  background: rgba(240, 248, 255, 0) !important;
  border: rgba(240, 248, 255, 0) !important;
}

.p-menu-list {
  justify-content: space-around !important;
}

.col-fixed {
  padding: 0 !important;
}

.active-link-exact {
  color: fuchsia !important;
}

.p-menubar {
  background: rgba(240, 248, 255, 0) !important;
  border: rgba(240, 248, 255, 0) !important;
}

.p-menubar-start {
  width: 350px;
}
</style>
