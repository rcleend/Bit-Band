import Tone from 'tone'

// Configure websocket uri and websocket functions
let wsuri = "ws://127.0.0.1:8080";
let sock = new WebSocket(wsuri);

sock.onopen = () => { console.log(`connected to ${wsuri}`); }

sock.onclose = (e) => { console.log(`connection closed (${e.code})`); }

sock.onmessage = (e) => { console.log(`message received: ${e.data}`); }


// Tone.js Synth initialization
let synth = new Tone.Synth().toMaster();

// Play a C5 note with a duration of an 8th note when the test-button has been clicked
document.getElementById("test").addEventListener("click", () => {
    synth.triggerAttackRelease('C5', '8n')
    sock.send('hoi');
});
