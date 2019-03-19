<template>
<div>
    <main>
        <component v-bind:is="currentView"></component>
    </main>
    <footer>
        <p class="text">
            &#169; 2019 
            <a href="https://github.com/rcleend">Rcleend</a>.
            All Rights Reserved.
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
            bandID: null,
            instrumentName: null,
            allInstruments: null
        }
    },
    created() {
        // /** INIT TONE.JS **/
        // Tone.Transport.start();
        // const synth = new Tone.Synth().toMaster();

        // let sequenceSteps = [
        //     "C4", "E4", "F4", "G4"j
        // ]

        // let sequence = new Tone.Sequence((time, note) => {
        //         synth.triggerAttackRelease(note, "8n", time);
        // }, sequenceSteps, "4n")


        


        this.$sock.onmessage = (e) => { 
            let message = JSON.parse(e.data);
            console.log(message);
            switch(message.type) {
                case "newInstrument":
                    if (!this.$data.instrumentName) this.$data.instrumentName = message.data.instrumentName
                    this.$data.allInstruments = message.data.allInstruments
                    this.$data.bandID = message.bandID;

                    break;
                case "removeInstrument":
                    console.log("removingg instrumenttt!!!!")
                    break;
                default:
                    console.log("something went wrong");
            }
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

