<template>
    <Sidebar />
    <div class="col">
        <Card>
            <template #header>
                Welcome to {{ save.saveName }}'s home page
                ID: {{ save.saveID }}
            </template>
        </Card>
    </div>
</template>

<script lang="ts" setup>
    import Sidebar from '../Components/Sidebar.vue'
    import { useRoute } from 'vue-router';
    import { nextTick, ref } from 'vue';
    import { SingleSave } from '../../../wailsjs/go/main/App'
    import { backend } from '../../../wailsjs/go/models'
    // export interface Props {
    //     id?: number
    // }
    // const props = withDefaults(defineProps<Props>(), {
    //     id: 0
    // })
    
    // const props = defineProps({
    //     id: String
    // })
    // var saveID: number | null
    // if (props.id == null) {
    //     saveID = +localStorage.getItem("defaultSave")
    // }
    const route = useRoute()
    // console.log(route.params.id)
    let save = ref({saveID: 0, managerName: "", gameVersion: 0, saveName: ''})
    SingleSave(+route.params.id).then((response) => {
        // let sessionVal: string = response.id.toString() + '_save_'
        save.value.saveID = response.id
        save.value.managerName = response.managerName
        save.value.gameVersion = response.gameVersion
        save.value.saveName = response.saveName
        // console.log(save.value)
        nextTick()
    })
    // console.log(SingleSave(+route.params.id))
    // const save = SingleSave(+this.$route.params.id)
</script>