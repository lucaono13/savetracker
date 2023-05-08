

<template>
  <Toast />
  <Menubar class="w-full">
    <template #start style="width:100px">
      <CascadeSelect style="width:75%" id="change" @change="SaveSelected" v-model="selectedSave" :options="isDataLoaded"
        optionLabel="name" optionGroupLabel="gameVersion" :option-group-children="['saves']" placeholder="Select save">
        <template #option="slotProps">
          <div class="flex align-items-center justify-content-center">
            <!-- <img :src="slotProps.option.image" style="height: 30px;" v-if="slotProps.option.name"> -->
            <FontAwesomeIcon icon="fa-solid fa-list" v-if="slotProps.option.gameVersion" />
            <!-- <i class="fa-solid fa-list"></i> -->
            <span style="padding-left: 5px;" v-if="slotProps.option.gameVersion">{{
              slotProps.option.gameVersion
            }}</span>
            <span style="padding-left: 5px;" v-if="slotProps.option.name">{{ slotProps.option.name }}</span>
          </div>
        </template>
      </CascadeSelect>

    </template>
    <template #end>
      <!-- <p>Save Tracker</p> -->
      <!-- <Button label="Toast" id="toastCheck" class="p-button-text p-button-outlined" @click="beError('Error adding transfers to DB.\nCheck log file for more details.')" >Toast Check</Button> -->
      <Button class="p-button-help p-button-outlined mr-3" @click="addSeasonModal=true">New Season</Button>
      
      <Button class="p-button-help p-button-outlined" @click="addSaveModal = true">New Save</Button>
      <AddSaveDialog v-model:visible="addSaveModal" @beError="beError" @saveAdded="GetSaves" @closeDialog="addSaveModal = false" />
      <AddSeasonDialog v-model:visible="addSeasonModal" @closeDialog="addSeasonModal=false"/>
    </template>
  </Menubar>

  <div class="grid w-full h-full">
    
    <!-- <div class="col"> -->
    <router-view @saveAdded="GetSaves" v-slot="{ Component, route }">
      <!-- <Sidebar v-if="sbVisible" :id="route.params.id" /> -->
      <!-- <div class="col"> -->
      <component :is="Component" :key="route.params.id"></component>
      <!-- </div> -->
    </router-view>
    <!-- </div> -->
  </div>
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

const toast = useToast()
const route = useRoute()
const router = useRouter()
let startup = ref(true)
let addSaveModal = ref(false)
let addSeasonModal = ref(false)

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
    // console.log(finalSaves.value.length)
    return finalSaves.value.length > 0 ? finalSaves.value : []
  }
)

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
    let savesMap = new Map<number, {}[]>()
    // let savesList: {}[] = []
    for (var save in saveList) {
      let gV: number = saveList[save].gameVersion
      // let image = await fetch(saveList[save].saveImage)
      let saveObj: { id: number, name: string, manager: string, image: string | undefined } = { id: saveList[save].id, name: saveList[save].saveName, manager: saveList[save].managerName, image: saveList[save].saveImage }
      
      if (saveObj.image != undefined) {
        GetImage(saveObj.image).then( (response) => {
          saveObj.image = response
        })
        // try {
        //   new URL(saveObj.image)
        //   console.log(saveObj.image)
        // } catch (e) {
        //   GetImage(saveObj.image).then(async (response) => {
        //   // console.log(response)
        //   // console.log('hi')
        //   saveObj.image = response
        // })
        // }
        
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
    nextTick()
    if (newID != null) {
      console.log(newID)
      if (newID != 0) {
        GoToSave(newID)
      } else {
        if (finalSaves.value.length > 0 && defaultID != null){
          GoToSave(+defaultID)
        } else {
          router.replace("/")
        }
        // TODO: show notification that error occurred when adding save
        // if 
      }
    } else if (defaultID != null) {
      GoToSave(+defaultID)
    }
  })
}

// Change path
function GoToSave(id: number) {
  let newRoute: string = '/save/' + id
  router.replace(newRoute)
}

// Function to be called when the dropdown changes selection
function SaveSelected(e: { originalEvent: Event; value: { id: number; name: string } }) {
  GoToSave(e.value.id)
}
GetSaves()
</script>

<!-- <script lang="ts">
  // Below is checking if there are any saves in database
  // TODO: when page is loaded, check
  import { nullLiteral } from '@babel/types'
  import { def } from '@vue/shared'
  import { nextTick, ref } from 'vue'
  import { RetrieveSaves } from '../wailsjs/go/main/App'
  import { backend } from '../wailsjs/go/models'
  import Sidebar from './components/Components/Sidebar.vue'
  import { useRoute } from 'vue-router'
  // import 

  const router = useRoute()

  // Setup for setting sidebar visibility (only if there are 1+ saves)
  var noOfSaves: null | string = localStorage.getItem("saves")
  if (noOfSaves == null) {
    noOfSaves = "0"
  }
  let sidebarV = false
  if (+noOfSaves > 0) {
    let sidebarV = true
  }
  let showSidebar = ref(sidebarV)

  // Setup for setting the default save
  let defaultID : string | null = localStorage.getItem("defaultSave")
  let defaultSaveObj : {id: number, name: string} | null
  let defaultSave = ref({})
  // console.log(localStorage.getItem("defaultSave"), localStorage.getItem("saves"))
  // Setup for getting all available saves
  let savesList: {}[] = []
  let finalSaves = ref(savesList)

  // Function that gets all the saves in DB and sets the selection to the selected default save
  function getSaves(): void {
    let savesMap = new Map<number, {}[]>()
    let savesList: {}[] = []
    // let finalSaves = ref(savesList)
    RetrieveSaves().then((response) => {
      for (var i in response) {
        let saveObj : {id: number, name: string, manager: string} = {id: response[i].id, name: response[i].saveName, manager: response[i].managerName}
        if (!savesMap.has(response[i].gameVersion)) {
          
          savesMap.set(response[i].gameVersion, [saveObj])
        } else if (savesMap.has(response[i].gameVersion)){
          let newSave: {}[] = savesMap.get(response[i].gameVersion)!
          
          newSave.push(saveObj)
          savesMap.set(response[i].gameVersion, newSave)
        }
        if (defaultID != null && saveObj.id == +defaultID) {
          defaultSave.value = saveObj
        }
      }    
      finalSaves.value = Array.from(new Map([...savesMap.entries()].sort()), ([gameVersion, saves]) => ({gameVersion, saves}))
      
      // Changes to the next tick in order to change the variables in the frontend
      nextTick()
      localStorage.setItem("saves", finalSaves.value.length.toString())
      
      
      showSidebar.value = true
    })
  }
  getSaves()
  
  
  export default {
    data() {
      return {
        finalSaves: finalSaves,
        
        selectedSave: defaultSave,
        sbVisible : showSidebar,
        items: [
          {
            label: "Test",
            icon: "pi pi-fw pi-file"
          }
        ]
      }
    },
    components: {
      Sidebar
    },
    methods: {
      GetSaves() {
        getSaves()
      },
      saveSelect(e: { originalEvent: Event; value : { id: number; name: string}}) {
        console.log(e.value.id, "changed")
        console.log(this.$router.currentRoute)
        this.$router.replace({name: 'save', params: { id: e.value.id } })
        console.log(this.$router.currentRoute)
        nextTick()
      },
    },
  }

// This template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
</script> -->

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
  // width: 100%!important;
  // width:calc(max-content+25px)!important;
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
  // padding-right: 5px!important;
}

.p-menubar-start {
  width: 250px;
}
</style>
