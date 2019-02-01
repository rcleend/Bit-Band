import Tone from 'tone/Tone/index.js'

let synth = new Tone.Synth().toMaster();

document.getElementById("test").addEventListener("click", () => {
    synth.triggerAttackRelease('C4', '8n')
});
