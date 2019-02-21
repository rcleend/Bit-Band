/** GLOBAL IMPORTS **/
import './assets/scss/index.scss';
import Tone from 'tone'
import Vue from 'vue'
import App from './assets/templates/app.vue'

// TODO: remove when in production
Vue.config.devtools = true;

/** INIT WEBSOCKET **/
let wsuri = "ws://127.0.0.1:8080"
let sock = new WebSocket(wsuri)

sock.onopen = () => { console.log(`connected to ${wsuri}`); }

sock.onclose = (e) => { console.log(`connection closed (${e.code})`); }

sock.onmessage = (e) => { 
    console.log("Message received")
}

/** INIT VUE **/
new Vue({
  el: '#app',
  render: h => h(App),
})

/** INIT TONE.JS **/
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
