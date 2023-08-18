<template>
    <Sidebar :saveSidebar="true" class="col-fixed relative left-0" @getSaves="getSaves" @beError="beError" @seasonAdded="seasonAdded" style="width:205px!important"/>
    <div class="col-auto ">
            <router-view :key="$route.fullPath" class="mt-3" @beError="beError" v-slot="{Component}">
                <component :is="Component" :key="updateComponent"></component>
            </router-view>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from '../Components/Sidebar.vue'
const emit = defineEmits(['beError', 'getSaves'])

let updateComponent = ref(0)

function seasonAdded() {
    updateComponent.value += 1
}

function beError(e: string) {
    emit('beError', e)
}

function getSaves() {
    emit('getSaves')
}

</script>

<style  lang="scss">

.fullWidthAndHeight {
    width: calc(100vw - 205px - 20px)!important;
}
</style>