<template>
    <div class=" flex almost-fh justify-content-center flex-column " v-if="saveSidebar">
      <div class="flex justify-content-center ">
        <Button label="changeImage" class="saveImageButton my-3" @click="NewImage" >
          <img :src="saveImage" style="width:175px" v-if="!imageError && saveImage.length > 0"/>
          <div class="p-2 justify-content-center align-items-center" v-if="saveImage == '' && !imageError">
            <span style="height: 125px;color: var(--text-color)" >No Save Image</span>
            <p style=" color: var(--text-color)">Click here to add one!</p>
          </div>
          <div class="p-4 justify-content-center align-items-center" v-if="imageError">
            <FontAwesomeIcon icon="circle-exclamation" size="3x" style="color: var(--text-color)" />
            <br/>
            <span style="height: 125px;color: var(--text-color)" >Error getting image</span>
          </div>
        </Button>
      </div>
      <Button class="p-button-help p-button-outlined mx-5 justify-content-center" @click="addSeasonModal=true">New Season</Button>
      <Menu :model="$router.getRoutes()" class=" align-content-evenly">
        <template #start>
          <div class="flex w-full justify-content-center my-2">
            <Checkbox @change="ChangeDefault" inputId="default" v-model="isDefault" :binary="true" />
            <label for="default" style="padding-left: 3px;">Default Save</label>
          </div>
          
        </template>
        <template #item="{ item }">
          <span :class="{ 'hidden': item.meta.hidden}">
          <font-awesome-icon :icon="item.meta.icon" />
          <router-link :to="item.path.replace(':id', route.params.id)" icon custom v-slot="{  href, route, navigate, isActive, isExactActive }">
            <Button @click="navigate" link :href="href" :class="{ 'active-link': isActive, 'active-link-exact': isExactActive }" class="w-full justify-content-center" style="color:darkturquoise">
              {{ route.name  }}
            </Button>
          </router-link>
          </span>
        </template>
        <template #end class="flex align-content-end align-self-end">
          
        </template>
      </Menu>
          <Button class="mx-4 justify-content-center align-self-end" @click="deleteSave"  severity="danger" text raised>Delete Save</Button>
        <AddSeasonDialog v-model:visible="addSeasonModal" @beError="beError" @closeDialog="addSeasonModal=false"/>
    </div>
    <div class=" flex almost-fh justify-content-center flex-column " v-if="!saveSidebar" style="width: 205px!important;">
      <Menu :model="$router.getRoutes()" class=" align-content-evenly">
        <template #start>
        </template>
        <template #item="{ item }">
          <span :class="{ 'hidden': !item.meta.totals}">
          
          <router-link :to="item.path.replace(':id', route.params.id)" icon custom v-slot="{  href, route, navigate, isActive, isExactActive }">
            <Button @click="navigate" link :href="href" :class="{ 'active-link': isActive, 'active-link-exact': isExactActive }" class="w-full justify-content-center" style="color:darkturquoise">
              {{ route.name  }}
            </Button>
          </router-link>
          </span>
        </template>
        <template #end class="flex align-content-end">
          
        </template>
      </Menu>
    </div>
    <ConfirmDialog group="saveDeletion"></ConfirmDialog>
</template>

<script setup lang="ts">
  import { onMounted, ref } from 'vue'
  import { useRoute, useRouter } from 'vue-router';
  import { useConfirm } from 'primevue/useconfirm'
  import { DeleteSave, SingleImage, GetImage, UploadSaveImage } from '../../../wailsjs/go/main/App'
  import AddSeasonDialog from './AddSeasonDialog.vue';
  import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { main } from '../../../wailsjs/go/models';
  interface Props {
    id?: string | string[],
    saveSidebar: boolean,
  }
  const confirm = useConfirm()
  const emit = defineEmits(['beError', 'getSaves'])
  const saveSidebar = ref(false)
  const route = useRoute()
  const router = useRouter()
  let isDefault = ref(false)
  const addSeasonModal = ref(false)
  const props = defineProps<Props>()
  let defaultSave: string | null = localStorage.getItem("defaultSave")
  if (defaultSave) {
    if (+defaultSave ==  +route.params.id) {
      isDefault.value = true
    }
  }
  let saveImage = ref('')
  const imageError = ref(false)
  
  const deleteSave = () => {
    confirm.require({
      message: "Are you sure you want to delete the save? This cannot be undone!",
      group: "saveDeletion",
      header: "Delete Confirmation",
      acceptClass: 'p-button-danger',
      accept: () => {
        DeleteSave(+route.params.id).then( (response: main.ErrorReturn) => {
          if (response.Error) {
            beError(response.Error)
            return
          }
          router.replace({name: "All Saves Home", replace: true})
        })
      },
      reject: () => {
        return
      }
    })
  }

  function beError(e: string) {
    emit('beError', e)
  }

  function ChangeDefault() {
    if (isDefault.value == false) {
      localStorage.setItem("defaultSave", "")
    } else {
      localStorage.setItem("defaultSave", route.params.id.toString())
    }
  }

  function NewImage() {
    UploadSaveImage(+route.params.id).then( (response) => {
      if (response.Error) {
        beError(response.Error)
        return
      }
      if (response.ImageFile.length != 0) {
        GetImage(response.ImageFile).then( (b64) => {
          saveImage.value = b64.b64Image
        })
      }
    })
  }
  
  onMounted(() => {
    saveSidebar.value = props.saveSidebar
    if (saveSidebar.value == true) {
      SingleImage(+route.params.id).then( (response) => {
        if (response.length == 0){
          return
        }
        GetImage(response).then( (b64) => {
          if (b64.Error) {
            console.log(b64)
            imageError.value = true
          }
          saveImage.value = b64.b64Image
        })
      })
    }
    
  })

</script>

<style lang="scss">

.saveImageButton {
  border: 1px solid white!important;
  padding: 0!important;
  border-radius: 0px!important;
  background-color: transparent!important;
}

.almost-fh {
  height: calc(100vh - 63px - 50px);
}

.router-btn {
  background: none!important;
  border: 0px solid white !important;
}

</style>