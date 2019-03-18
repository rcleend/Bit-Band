<template>
<table class="sequencer">
   <sequence-row v-for="(sound, index) of sounds" class="sequence-row" :rowID="`r${index}`" :sound="sound" :noteAmount="noteAmount" :synth="synth"></sequence-row>
</table>
</template>

<script>
import sequenceRow from './components/sequenceRow.vue';
export default {
   components : {
      sequenceRow
   },
   data() {
      return {
         Instrument: "Drum",
         sounds: ["Kick", "Snare", "Hi-Hat", "Tom A", "Tom B"],
         noteAmount: 16,
         rowArray: null,
         synth: new this.$tone.MembraneSynth().toMaster()
      }
   },
   created() {
         console.log('sequencer');
   },
   methods: {
      
   }
}
</script>

