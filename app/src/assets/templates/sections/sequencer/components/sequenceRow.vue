<template>
<tr>
    <h2 class="sub-title">{{sound}}</h2>
    <sequence-note v-for="(index, note) in noteAmount" :noteID="`${rowID}n${index}`" :synth="synth"></sequence-note>
</tr>
</template>

<script>
import sequenceNote from './sequenceNote.vue';
export default {
    props: [
        'rowID',
        'sound',
        'noteAmount',
        'synth'
    ],
    components: {
        sequenceNote
    },
    data() {
        return {
            noteArray: null
        }
    },
    created() {
         console.log('hoi');
    },
   methods: {
      
   }
}
</script>