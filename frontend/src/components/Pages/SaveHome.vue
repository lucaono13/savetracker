<template>
    <Sidebar/>
    <div class="col w-full">
        <div class="grid">
            <!-- <img class="col-2" :src="save.image"  v-if="save.image" /> -->

            <!-- <Card class="col">
                <template #header>

                </template>
                <template #title>
                    {{ save.saveName }}
                </template>
            </Card> -->
            <div class="col-10 justify-content-center">
                <p>{{ save.saveName }}</p>
                <p>{{ save.managerName }}</p>
            </div>
        </div>

    </div>
</template>

<script lang="ts" setup>
import Sidebar from '../Components/Sidebar.vue'
import { useRoute } from 'vue-router';
import { nextTick, ref } from 'vue';
import { GetImage, SingleSave } from '../../../wailsjs/go/main/App'
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
const emit = defineEmits(['beError'])
// console.log(route.params.id)
let imgPlaceholder: string | undefined
let save = ref({ saveID: 0, managerName: "", gameVersion: 0, saveName: '', image: imgPlaceholder })
SingleSave(+route.params.id).then((response) => {
    if (response.Error != "") {
        emit('beError', response.Error)
    }
    let result = response.Save
    // let sessionVal: string = response.id.toString() + '_save_'
    save.value.saveID = result.id
    save.value.managerName = result.managerName
    save.value.gameVersion = result.gameVersion
    save.value.saveName = result.saveName
    if (result.saveImage) {
        GetImage(result.saveImage).then(async (result) => {
            save.value.image = result
        })
    }
    // save.value.image =  response.saveImage
 
    // console.log(save.value.image)

    nextTick()
})


    // console.log(SingleSave(+route.params.id))
    // const save = SingleSave(+this.$route.params.id)
</script>