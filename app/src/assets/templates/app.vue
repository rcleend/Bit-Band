<template>
<div>
    <main>
        <component v-bind:is="currentView"></component>
    </main>
    <footer>
        <p class="text">
            &#169; 2019 
            <a href="https://github.com/rcleend">Rcleend</a> 
            & 
            <button id="test" @click="setCurrentView('sequencer')">Jelmuis</button>
            . All Rights Reserved.
        </p>
    </footer>
</div>
</template>

<script>
import titleScreen from './sections/titleScreen/titleScreen.vue';
import sequencer from './sections/sequencer/sequencer.vue';

export default {
    components: {
        titleScreen,
        sequencer
    },
    data() {
        return {
            currentView: "titleScreen",
        }
    },
    created() {
        // /** INIT TONE.JS **/
        // Tone.Transport.start();
        // const synth = new Tone.Synth().toMaster();

        // let sequenceSteps = [
        //     "C4", "E4", "F4", "G4"
        // ]
        // let sequence = new Tone.Sequence((time, note) => {
        //         synth.triggerAttackRelease(note, "8n", time);
        // }, sequenceSteps, "4n")




        this.$sock.onmessage = (e) => { 
            console.log("Message received")
        }
        // sequence.start();
    },
    methods: {
        setCurrentView: function(component) {
            this.$data.currentView = component; 
        }
    }
};
</script>

