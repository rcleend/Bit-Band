<template>
<td>
    <div class="sequence-note" v-bind:class="{ enabled: isEnabled }" @click="toggleButton()">O</div>
</td>
</template>

<script>
export default {
    props: [
        'noteID',
        'synth'
    ],
    data() {
        return {
            isEnabled: false
        }
    },
    methods: {
        toggleButton: function() {
            this.$data.isEnabled = !this.$data.isEnabled;
            console.log(`${this.$props.noteID} : ${this.$data.isEnabled}`);
            
            this.$props.synth.triggerAttackRelease("C1", "2n");

            let test = {
                "type": "update",
                "data": {
                   "instrument": "drum",
                   "note": this.$props.noteID
                }
             }

            if(this.$data.isEnabled){
                this.$sock.send(JSON.stringify(test));
            }
        }
    }
}
</script>