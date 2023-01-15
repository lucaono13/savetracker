<template>
<div class="flex align-items-center justify-content-center h-full">
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

<script lang="ts">
    import { required } from "@vuelidate/validators";
    import { useVuelidate } from "@vuelidate/core";
    import { AddNewSave } from "../../../wailsjs/go/main/App";
    import AddSaveDialog from "../Components/AddSaveDialog.vue" ;
    export default {
        setup: () => ({ v$: useVuelidate()}),
        data() {
            return {
                addSaveModal: false,
                // saveName: '',
                // gameVersion: '',
                // managerName: '',
                // blockedDocument: false,
                // submitted: false,
                // versions: [
                //     {name: "2023"},
                //     {name: "2022"},
                //     {name: "2021"},
                //     {name: "2020"}
                // ],
                // addingToDB: false,
            }
        },
        methods: {
            openAddSave() {
                this.addSaveModal = true
            },
            
        },
        // eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
        
        components: {
            AddSaveDialog
        }
    }
</script>