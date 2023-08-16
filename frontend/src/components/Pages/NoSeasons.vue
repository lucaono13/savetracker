<template>
    <div class="flex align-items-center align-content-center justify-content-center w-screen" style="height: calc(100vh - 79px)!important">
        <Card style="width:500px">
            <template #title>
                No Seasons Found
            </template>
            <template #content>
                No seasons data found! Go into any of your saves by using the dropdown in the top left and add a new season to fill the database with some data!
            </template>
            <!-- <template #footer>
                <Button label="Add Save" class="p-button" @click="openAddSave" />
            </template> -->
        </Card>
        <!-- <AddSaveDialog v-model:visible="addSaveModal" @saveAdded="$emit('saveAdded')" @closeDialog="addSaveModal=false"/> -->
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue';
    // import AddSaveDialog from "../Components/AddSaveDialog.vue";
    // let addSaveModal = ref(false)
    
    // function openAddSave() {
    //     addSaveModal.value = true
    // }
</script>