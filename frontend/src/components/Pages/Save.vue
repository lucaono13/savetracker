<template>
    <Sidebar class="col-2" style="width:205px!important"/>
    <div class="col-auto">
        
        <!-- <div class="col-auto"> -->
            <router-view class="mt-3"></router-view>
        <!-- </div> -->
    </div>
</template>

<script setup lang="ts">
import Sidebar from '../Components/Sidebar.vue'
</script>