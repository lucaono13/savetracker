

<template>
  <Menubar class="w-full" :model="items">
    <template #start>
      <!-- <OverlayPanel ref="op" :showCloseIcon="true" >
        
      </OverlayPanel> -->
      <CascadeSelect id="change" @change="saveSelect" v-model="selectedSave" :options="finalSaves" optionLabel="name" optionGroupLabel="gameVersion" :option-group-children="['saves']" placeholder="Select save">
        
      </CascadeSelect>
    </template>
    <template #end>
      <p>Save Tracker</p>
    </template>
  </Menubar>
  <div class="grid w-full h-full">
    <Sidebar v-if="sbVisible"/>
    <div class="col">
      <router-view @saveAdded="GetSaves" v-slot="{ Component, route }">
        <component :is="Component" :key="route.params.id"></component>
      </router-view>
    </div>
  </div>
</template>



<script lang="ts">
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

        // console.log(typeof(e))
      },
    },
    // watch: {
    //   router()

    //   }
    // }

  }

// This template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
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
  height: 100%!important;
  padding: .75rem!important;
  // width: 100%!important;
  // width:calc(max-content+25px)!important;
  background: rgba(240, 248, 255, 0)!important;
  border: rgba(240, 248, 255, 0)!important;
}

.p-menu-list {
  justify-content:space-around!important;
}

.col-fixed {
  padding: 0!important;
}

.active-link-exact {
  color: fuchsia!important;
}

.p-menubar{
  background: rgba(240, 248, 255, 0)!important;
  border: rgba(240, 248, 255, 0)!important;
  // padding-right: 5px!important;
}

</style>
