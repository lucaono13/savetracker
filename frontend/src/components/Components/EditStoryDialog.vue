<template>
    <Dialog modal :draggable="false" class="w-9" :close-on-escape="false" :closable="true" header="Edit your save story">
        <Editor v-model="story"/>
        <template #footer>
            <Button label="Cancel" text @click="$emit('closeDialog')"/>
            <Button label="Save" @click="updateStory"/>
        </template>
    </Dialog>
</template>

<script setup lang="ts">
import { ref, defineProps, onMounted, defineEmits } from 'vue';

const props = defineProps<{
    story: string,
}>()
const emit = defineEmits(['closeDialog','updateStory'])
const story = ref('')
onMounted( () => {
    story.value = props.story
})

function updateStory() {
    emit('updateStory', story.value)
    emit('closeDialog')
}

</script>