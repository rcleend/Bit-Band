<template>
<div>
    <main>
        <h1>instrument:<span id="instrument-name--test"></span></h1>
        <h1>band:<span id="band-name--test"></span></h1>
        <component v-bind:is="currentView"></component>
    </main>
    <footer>
        <p class="text">
            &#169; 2019 <a href="https://github.com/rcleend">Rcleend</a>. All Rights Reserved.
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
            currentView: "sequencer",
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

        this.handleSocketMessages();


          // sequence.start();
    },
    methods: {
        setCurrentView: function(component) {
            this.$data.currentView = component; 
        },
        handleSocketMessages: function() {
            this.$sock.onmessage = (e) => { 
                var message = JSON.parse(e.data)
                switch (message.type) {
                    case ("updateInstrument"):
                        this.initInstruments(message);
                    break;
                    default:
                    break;
                }
            }
        },
        initInstruments: function(message) {
            document.getElementById("instrument-name--test").innerHTML = message.data.instrument
            document.getElementById("band-name--test").innerHTML = message.Room
        }
    }
};
</script>

