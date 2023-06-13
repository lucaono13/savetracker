<template>
    <Sidebar class="col-fixed relative left-0" @beError="beError" style="width:205px!important"/>
    <div class="col-auto ">
            <router-view class="mt-3" @beError="beError"></router-view>
    </div>
</template>

<script setup lang="ts">
import Sidebar from '../Components/Sidebar.vue'
const emit = defineEmits(['beError'])

function beError(e: string) {
    emit('beError', e)
}

</script>

<style  lang="scss">

.fullWidthAndHeight {
    width: calc(100vw - 205px - 20px)!important;
}
</style>