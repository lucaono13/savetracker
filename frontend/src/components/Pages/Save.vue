<template>
    <Sidebar class="col-2" @beError="beError" style="width:205px!important"/>
    <div class="col-auto">
        
        <!-- <div class="col-auto"> -->
            <router-view class="mt-3" @beError="beError"></router-view>
        <!-- </div> -->
    </div>
</template>

<script setup lang="ts">
import Sidebar from '../Components/Sidebar.vue'
const emit = defineEmits(['beError'])

function beError(e: string) {
    emit('beError', e)
}

</script>