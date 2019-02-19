
import Vue from 'vue'
import App from './app.vue'

import Tone from 'tone'

new Vue({
  el: '#app',
  render: h => h(App),
})

// Configure websocket uri and websocket functions
let wsuri = "ws://127.0.0.1:8080"
let sock = new WebSocket(wsuri)

sock.onopen = () => { console.log(`connected to ${wsuri}`); }

sock.onclose = (e) => { console.log(`connection closed (${e.code})`); }

sock.onmessage = (e) => { 
    console.log("Message received")
}

// Tone.js Synth initialization

Tone.Transport.start();
const synth = new Tone.Synth().toMaster();

let sequenceSteps = [
    "C4", "E4", "F4", "G4"
]
let sequence = new Tone.Sequence((time, note) => {
        synth.triggerAttackRelease(note, "8n", time);
}, sequenceSteps, "4n")

document.getElementById("test").addEventListener("click", () => {
    sequence.start();
    console.log("hello world");
    sock.send({"text" : "hoi"})
});