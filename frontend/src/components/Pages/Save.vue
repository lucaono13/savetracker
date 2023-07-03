<template>
    <Sidebar :saveSidebar="true" class="col-fixed relative left-0" @getSaves="getSaves" @beError="beError" style="width:205px!important"/>
    <div class="col-auto ">
            <router-view :key="$route.fullPath" class="mt-3" @beError="beError"></router-view>
    </div>
</template>

<script setup lang="ts">
import Sidebar from '../Components/Sidebar.vue'
const emit = defineEmits(['beError', 'getSaves'])

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