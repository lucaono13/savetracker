<template>
    <div class="flex align-items-center align-content-center justify-content-center w-screen" style="height: calc(100vh - 79px)!important">
        <Card style="width:500px">
            <template #title>
                No Saves DB Found
            </template>
            <template #content>
                Either this is the first time loading/using the app or you deleted all your saves! (Hopefully not by accident).
                Click on the "Add Save" button below to add a save!
            </template>
            <template #footer>
                <Button label="Add Save" class="p-button" @click="openAddSave" />
            </template>
        </Card>
        <AddSaveDialog v-model:visible="addSaveModal" @saveAdded="$emit('saveAdded')" @closeDialog="addSaveModal=false"/>
    </div>
</template>

<script setup lang="ts">
    import { ref } from 'vue';
    import AddSaveDialog from "../Components/AddSaveDialog.vue";
    let addSaveModal = ref(false)
    
    function openAddSave() {
        addSaveModal.value = true
    }
</script>