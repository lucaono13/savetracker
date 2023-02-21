<template>
    <div class="col-fixed h-full flex  justify-content-center flex-column relative" style="width:205px">
      <div class="flex justify-content-center">
        <Button label="changeImage" class="saveImageButton my-3" @click="NewImage">
          <img :src="saveImage" style="width:175px" v-if="saveImage != ''"/>
          <div class="p-2 justify-content-center align-items-center" v-if="saveImage == ''">
            <span style="height: 125px;color: var(--text-color)" >No Save Image</span>
            <p style=" color: var(--text-color)">Click here to add one!</p>
          </div>
          
        </Button>
      </div>
      <!-- <Button label="changeImage" style="flex-initial"><img :src="saveImage" style="width:100px" v-if="saveImage != ''"/></Button> -->
      
      <Menu :model="$router.getRoutes()" class=" align-content-evenly">
        <template #start>
          <div>
            <Checkbox @change="ChangeDefault" inputId="default" v-model="isDefault" :binary="true" />
            <label for="default" style="padding-left: 3px;">Default Save</label>
          </div>
          
        </template>
        <template #item="{ item }">
          <!-- <br> -->
          <span :class="{ 'hidden': item.meta.secondary}">
          <font-awesome-icon :icon="item.meta.icon" />
          <router-link :to="item.path.replace(':id', route.params.id)" icon custom v-slot="{  href, route, navigate, isActive, isExactActive }">            
            <a :href="href" @click="navigate" :class="{ 'active-link': isActive, 'active-link-exact': isExactActive }" style="color:darkturquoise">
              {{ route.name }}
            </a>
          </router-link>
          </span>
        </template>
      </Menu>
    </div>
</template>

<script setup lang="ts">
  import { ref } from 'vue'
  import { useRoute } from 'vue-router';
  import { SingleImage, GetImage, UploadSaveImage } from '../../../wailsjs/go/main/App'
  interface Props {
    id?: string | string[]
  }
  const route = useRoute()
  let isDefault = ref(false)
  const props = defineProps<Props>()
  let defaultSave: string | null = localStorage.getItem("defaultSave")
  console.log("default save ",defaultSave)
  console.log(route.fullPath)
  if (defaultSave) {
    if (+defaultSave ==  +route.params.id) {
      isDefault.value = true
    }
  }
  let saveImage = ref('')
  
  SingleImage(+route.params.id).then( (response) => {
    GetImage(response).then(async (b64) => {
      saveImage.value = b64
      
    })
  })

  function ChangeDefault() {
    if (isDefault.value == false) {
      localStorage.setItem("defaultSave", "")
    } else {
      localStorage.setItem("defaultSave", route.params.id.toString())
    }
  }

  function NewImage() {
    UploadSaveImage(+route.params.id).then( (response) => {
      if (response.length != 0) {
        GetImage(response).then(async (b64) => {
          saveImage.value = b64
        })
      }
    })
  }
  
</script>

<style lang="scss">

.saveImageButton {
  border: 1px solid white!important;
  padding: 0!important;
  border-radius: 0px!important;
  background-color: transparent!important;
}

</style>