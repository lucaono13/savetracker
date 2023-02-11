<template>
    <div class="col-fixed h-full" style="width:205px">
      <Menu :model="$router.getRoutes()" class="fixed align-content-evenly">
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

  function ChangeDefault() {
    if (isDefault.value == false) {
      localStorage.setItem("defaultSave", "")
    } else {
      localStorage.setItem("defaultSave", route.params.id.toString())
    }
  }
  
</script>